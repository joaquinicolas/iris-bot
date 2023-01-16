# FROM golang image
FROM golang:1.18-bullseye as builder
WORKDIR /app/github.com/joaquinicolas/iris-bot


COPY go.mod go.sum ./
COPY src/ ./src
COPY main.go ./
COPY .golangci.yml ./
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
RUN golangci-lint run -v
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

FROM scratch
# Set environment variables
ENV GSHEET_TOKEN ""
ENV TELEGRAM_TOKEN ""
ENV SHEET_ID ""
ENV SHEET_RANGE ""

WORKDIR /
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]
