FROM golang:alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

EXPOSE 3000
ENTRYPOINT ["/backend"]