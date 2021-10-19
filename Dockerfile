FROM golang:1.15-alpine as builder
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN GOOS=linux GOARCH=amd64 go build -o cssh-acs main.go

FROM alpine
WORKDIR /cssh-acs
COPY --from=builder /cssh-acs .
EXPOSE 80
ENTRYPOINT ["/cssh-acs/cssh-acs"]
CMD ["-dbDir=/cssh-acs/db"]
