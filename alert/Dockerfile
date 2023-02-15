FROM golang
WORKDIR /alert
COPY main.go .
RUN go build ./main.go
EXPOSE 8383
ENTRYPOINT ["./main"]