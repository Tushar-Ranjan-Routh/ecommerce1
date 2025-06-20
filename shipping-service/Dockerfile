FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o shipping-service .
EXPOSE 8090
ADD https://github.com/eficode/wait-for/releases/download/v2.2.3/wait-for /wait-for
RUN chmod +x /wait-for
CMD ["/wait-for", "kafka:9092", "--", "./shipping-service"]
