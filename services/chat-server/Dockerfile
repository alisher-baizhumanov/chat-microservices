# Run from root project directory
# Step 1: Prepare
FROM golang:1.22 as builder

WORKDIR /app

COPY vendor ./vendor
COPY . .

WORKDIR /app/services/chat-server
RUN go env -w CGO_ENABLED=0
RUN go build -o /bin/chat-server -mod=vendor -v ./main.go

# Step 2: Final
FROM alpine:3.18

COPY --from=builder /bin/chat-server /app

CMD ["/app"]