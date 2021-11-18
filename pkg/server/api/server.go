package api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kreexzoi/grpc-mock/pkg/stub"
	pbutils "github.com/kreexzoi/grpc-mock/pkg/utils/pb"
	mockpb "github.com/kreexzoi/grpc-mock/proto/mock"
)

type StubManager interface {
	FindStubs(service, method string, in map[string]interface{}) []*stub.Stub
	AddStub(stub *stub.Stub) error
	DeleteStub(service, method string) error
}

type Server struct {
	stubMgr StubManager
}

func NewServer(mgr StubManager) *Server {
	return &Server{stubMgr: mgr}
}

func (s *Server) AddStubs(ctx context.Context, req *mockpb.AddStubsRequest) (*mockpb.AddStubsResponse, error) {
	if len(req.Stubs) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing stubs")
	}

	stubs, err := PBStubsToStubs(req.Stubs)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid stub: %v", err)
	}

	for _, stub := range stubs {
		if err := s.stubMgr.AddStub(stub); err != nil {
			return nil, err
		}

		log.Printf("api stub added: %s", spew.Sdump(stub))
	}
	return &mockpb.AddStubsResponse{}, nil
}

func (s *Server) FindStubs(ctx context.Context, req *mockpb.FindStubsRequest) (*mockpb.FindStubsResponse, error) {
	if req.Service == "" {
		return nil, status.Error(codes.InvalidArgument, "missing service name")
	}

	if req.Method == "" {
		return nil, status.Error(codes.InvalidArgument, "missing method")
	}

	stubs := s.stubMgr.FindStubs(req.Service, req.Method, pbutils.ToMap(req.In))

	rsp := &mockpb.FindStubsResponse{
		Stubs: StubsToPBStubs(stubs),
	}

	return rsp, nil
}

func (s *Server) DeleteStubs(ctx context.Context, req *mockpb.DeleteStubsRequest) (*mockpb.DeleteStubsResponse, error) {
	if req.Service == "" {
		return nil, status.Error(codes.InvalidArgument, "missing parameter: service_name")
	}
	if err := s.stubMgr.DeleteStub(req.Service, req.Method); err != nil {
		return nil, err
	}
	return &mockpb.DeleteStubsResponse{}, nil
}
