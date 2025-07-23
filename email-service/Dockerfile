FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

#Build the binary
RUN go build -o email-service ./cmd/service/main.go

FROM scratch

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY ./configs /configs

COPY --from=builder /build/email-service .

ENTRYPOINT ["/email-service", "configs/local.yaml"]