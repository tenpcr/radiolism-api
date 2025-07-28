# Stage 1: Build statically linked binary
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o radiolism_api

FROM scratch

WORKDIR /app

COPY --from=builder /app/radiolism_api .

CMD ["/app/radiolism_api"]