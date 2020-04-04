FROM golang:1.13 as builder

WORKDIR /myapp
COPY . .
RUN go build


FROM ubuntu:18.04

WORKDIR /myapp
COPY --from=builder /myapp/test-server .

CMD ["./test-server"]
