#
# Build
#
FROM golang:1.17.8-alpine as build
WORKDIR /go/src/
COPY . .
RUN apk add make && make build

#
# Deploy
#
FROM alpine
WORKDIR /go/src/
COPY --from=build /go/src /go/src
EXPOSE 8080
CMD ./service
