FROM golang

WORKDIR /src
EXPOSE 8080
ADD . .
RUN go mod download
RUN go build
ENTRYPOINT ["./server"]