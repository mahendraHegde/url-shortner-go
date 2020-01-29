FROM golang:1.13-alpine
ENV PATH="~/go/bin:${PATH}"
RUN apk add git
WORKDIR /go/src/api
COPY ./ ./
RUN go install
CMD go run main.go
 
