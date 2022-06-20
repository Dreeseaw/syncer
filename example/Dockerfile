FROM golang:1.18-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o server .
CMD ["./server"]
