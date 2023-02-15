FROM golang
WORKDIR /errand
COPY main.go .
RUN go build ./main.go
EXPOSE 8282
ENTRYPOINT ["./main"]