FROM golang:alpine as builder

RUN apk --update add upx

WORKDIR /app
ADD . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -installsuffix cgo -o hello main.go &&\
    upx /app/hello


FROM scratch 

COPY --from=builder /app/hello /bin/hello

CMD ["/bin/hello"]
