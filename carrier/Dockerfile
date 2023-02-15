FROM golang
WORKDIR /carrier
COPY main.go .
RUN go build main.go
EXPOSE 8181
ENTRYPOINT ["./main"]