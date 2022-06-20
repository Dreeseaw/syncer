FROM golang:1.18-bullseye
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o /usr/local/bin/server .
ENTRYPOINT ["server"]
