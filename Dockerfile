FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o radiolism_api .

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/radiolism_api .

EXPOSE 8080

ENTRYPOINT ["/app/radiolism_api"]