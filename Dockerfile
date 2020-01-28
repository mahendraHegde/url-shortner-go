FROM golang:1.13-alpine3.10
ENV PATH="~/go/bin:${PATH}"
RUN apk add git
RUN go get github.com/codegangsta/gin
EXPOSE 3000
WORKDIR /go/src/api
CMD  ["go","run","main.go"]
 
