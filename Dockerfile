FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#RUN go build -o main .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 8080
CMD ["./app"]
#ENTRYPOINT ["/app/go_client"]