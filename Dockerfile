FROM golang:alpine

ENV GOPATH=/

WORKDIR /app

COPY . .

RUN go get -d ./...
RUN go build cmd/apiserver/main.go

CMD [ "./main" ]