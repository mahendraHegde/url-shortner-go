FROM golang:1.13-alpine
ENV PATH="~/go/bin:${PATH}"
RUN apk add git
RUN go get -u github.com/swaggo/swag/cmd/swag
EXPOSE 3000
WORKDIR /go/src/api
CMD  swag init && go run main.go
 
