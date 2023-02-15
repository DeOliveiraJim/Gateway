FROM golang
WORKDIR /delivery
COPY main.go .
RUN go build ./main.go
EXPOSE 8484
ENTRYPOINT ["./main"]