FROM golang:1.18-rc-alpine
COPY . .
RUN go get .
RUN go build
CMD["go", "run"]