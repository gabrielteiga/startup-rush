FROM golang:1.23.5

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api-server/main.go

EXPOSE 8080
CMD ["/app/main"]