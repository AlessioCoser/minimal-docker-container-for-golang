FROM golang:1.9-alpine
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /
COPY --from=0 /go/src/app /
CMD ["/app"]
