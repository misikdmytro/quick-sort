FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app main.go

FROM alpine

WORKDIR /out
COPY --from=builder /out/app /out/app

CMD ["./app"]