FROM golang:1.21.1-alpine AS builder

COPY . /app/
WORKDIR /app

RUN go mod download
RUN go build -o /bin/uzum_deliviry_service ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/
copy --from=builder /bin/uzum_deliviry_service .

CMD ["./uzum_deliviry_service"]