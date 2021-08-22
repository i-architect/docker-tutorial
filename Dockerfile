FROM golang:1.16-alpine3.13 AS hoge

WORKDIR /

ADD main.go /main.go
ADD go.mod /go.mod
ADD go.sum /go.sum

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
RUN chmod +x /main

FROM scratch
COPY --from=hoge /main /main
CMD ["/main"]