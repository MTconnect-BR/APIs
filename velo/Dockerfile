FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY velo/go.mod velo/go.sum ./
RUN go mod download

COPY velo/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o velo ./cmd/velo

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/velo .
COPY --from=builder /app/configs/velo.production.yaml ./configs/velo.yaml

EXPOSE 8080

CMD ["./velo"]
