FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 5000
CMD [ "./server" ]