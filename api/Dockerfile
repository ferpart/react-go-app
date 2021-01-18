FROM golang:1.15.6
WORKDIR /api
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go get github.com/gorilla/mux
RUN go get github.com/sirupsen/logrus
EXPOSE 5000
COPY main.go .
RUN go build main.go
CMD ["./main"]