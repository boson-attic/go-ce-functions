from golang:alpine
COPY . /src
WORKDIR /src/function
RUN go build -o function
ENTRYPOINT ["/src/function/function"]
