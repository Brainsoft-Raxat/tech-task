FROM golang:1.21-alpine AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -mod vendor -ldflags '-w -s' -o /app/prog -v ./cmd/app/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/prog .

ENTRYPOINT ["/app/prog"]
