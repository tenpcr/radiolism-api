FROM golang:1.24.5 AS builder

WORKDIR /app

COPY . .

RUN go build -o radiolism_api


WORKDIR /app

COPY radiolism_api .


CMD ["go","run","radiolism_api"]