FROM golang:latest as builder

WORKDIR /app
ADD . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix main.go -o hello


FROM alpine:latest 

WORKDIR /root/

COPY --from=builder /app/hello .

CMD ["./hello"]
