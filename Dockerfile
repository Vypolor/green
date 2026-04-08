FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/server /server
COPY --from=builder /app/static /static

ENV SERVER_HOST=0.0.0.0
ENV SERVER_PORT=8080

EXPOSE 8080

ENTRYPOINT ["/server"]
