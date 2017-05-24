FROM golang
WORKDIR /go/src/start
COPY . .
WORKDIR /go/src/start/app
RUN go get
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
WORKDIR /go/src/start

# Open up the port where the app is running.
EXPOSE 9000

CMD ["revel", "run"]