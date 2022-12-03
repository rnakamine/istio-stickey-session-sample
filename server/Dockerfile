FROM golang:1.9
WORKDIR /tmp/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /tmp/server ./
CMD [ "./server" ]
