from golang:alpine

COPY ./function /project/function
COPY ./main.go /project/
COPY ./go.mod /project/
WORKDIR /project

RUN go build -o f
ENTRYPOINT ["/project/f"]
