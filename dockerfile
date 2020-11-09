FROM golang:1.14

RUN mkdir -p /dockerFiles/get

ADD . /dockerFiles/get

WORKDIR  /dockerFiles/get

RUN go get -u github.com/go-sql-driver/mysql

RUN go build -o get .
 
CMD ["/dockerFiles/get/get"]

