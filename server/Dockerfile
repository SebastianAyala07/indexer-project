FROM golang:alpine AS build

WORKDIR /go/src/server
COPY . .

RUN go build -o /go/bin/server server/cmd/main/

FROM scratch
COPY --from=build /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]
