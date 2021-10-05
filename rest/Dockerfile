FROM golang:1.16 as build

COPY . /go/src/go-rest

WORKDIR /go/src/go-rest

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o go-rest

FROM scratch as Image

COPY --from=build /go/src/go-rest/go-rest .

EXPOSE 9999

CMD ["/go-rest"]