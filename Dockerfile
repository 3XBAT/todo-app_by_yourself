FROM golang:1.23-alpine AS Builder

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /bin/application cmd/main.go

FROM alpine:latest AS Runner

COPY --from=builder /bin/application ./

COPY configs/config.yaml /config.yaml

CMD ["/application"]
