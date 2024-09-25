FROM golang:1.23-alpine AS Builder

WORKDIR /app

COPY ./ ./

#install psql
RUN apk update
RUN apk add --no-cache postgresql-client bash

RUN go mod download

RUN chmod +x wait-for-postgres.sh

RUN go build -o /bin/application cmd/main.go

FROM alpine:latest AS Runner

COPY --from=builder /bin/application ./

COPY configs/config.yaml /config.yaml

CMD ["/application"]
