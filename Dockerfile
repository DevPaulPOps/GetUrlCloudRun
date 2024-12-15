FROM golang:1.22.2
LABEL authors="Paul"
WORKDIR /go/src/app
COPY . /go/src/app
RUN go mod download
RUN go build -o /go/bin/app
EXPOSE 3000
CMD ["/go/bin/app"]
