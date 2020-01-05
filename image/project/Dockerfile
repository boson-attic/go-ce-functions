from golang:alpine as builder
RUN apk add make git gcc g++
COPY . /src
WORKDIR /src/user-app
run go build -o user-app

from alpine:latest
run apk add --no-cache ca-certificates
COPY --from=builder /src/user-app /bin/
ENTRYPOINT ["user-app"]
