FROM golang
WORKDIR /tea
COPY main.go .
RUN go build ./main.go
EXPOSE 8282
ENTRYPOINT ["./main"]