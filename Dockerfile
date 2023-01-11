FROM golang:1.17-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git curl && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b /go/bin v1.42.1

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN golangci-lint run
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]