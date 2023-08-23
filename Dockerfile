# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build
RUN apk --no-cache add curl

# Run stage
FROM golang:1.20-alpine
WORKDIR /app
COPY --from=builder /app/service-name .

EXPOSE 8001
CMD [ "/app/service-name" ]

