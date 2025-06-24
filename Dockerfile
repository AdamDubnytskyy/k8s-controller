    # syntax=docker/dockerfile:1.4
    FROM golang:1.24-alpine AS builder

    WORKDIR /app
    COPY . .

    ARG TARGETOS=linux
    ARG TARGETARCH=amd64
    ARG VERSION=dev

    RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -v -o k8s-controller -ldflags "-X=github.com/AdamDubnytskyy/k8s-controller/cmd.appVersion=$VERSION" main.go

    FROM gcr.io/distroless/static-debian12
    WORKDIR /
    COPY --from=builder /app/k8s-controller .
    EXPOSE 8080
    ENTRYPOINT ["/k8s-controller"]