FROM golang:1.17.2-buster AS app_builder
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -v -o main ./cmd/...

FROM alpine:3
WORKDIR /app

COPY --from=app_builder /build/main /app/main

ENTRYPOINT ["/app/main"]