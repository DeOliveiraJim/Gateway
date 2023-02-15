FROM golang
WORKDIR /planner
COPY main.go .
RUN go build ./main.go
EXPOSE 8585
ENTRYPOINT ["./main"]