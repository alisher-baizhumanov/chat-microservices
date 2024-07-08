# Run from root project directory
# Step 1: Prepare
FROM golang:1.22 as builder

WORKDIR /app

COPY vendor ./vendor
COPY . .

WORKDIR /app/services/auth
RUN go env -w CGO_ENABLED=0
RUN go build -o /bin/auth -mod=vendor -v ./main.go

# Step 2: Final
FROM alpine:3.18

COPY --from=builder /bin/auth /app

CMD ["/app"]