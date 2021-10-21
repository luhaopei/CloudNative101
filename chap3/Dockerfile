FROM golang:1.17

WORKDIR /src/

COPY main.go go.* /src/

RUN go build -o /bin/httpserver

CMD ["/bin/httpserver"]

