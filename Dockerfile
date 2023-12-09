FROM golang:1.21.4-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
COPY main.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-app .

FROM alpine:3.18.4
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/simple-app ./
CMD ["/root/simple-app"]