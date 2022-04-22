# FROM golang as build
FROM golang:1.17.1-buster AS build

WORKDIR /app

ADD . .
RUN go mod download

# RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o apigateway
# RUN CGO_ENABLED=0 GOOS=linux go build -o apigateway
RUN go build -o application

# FROM scratch as production
# FROM alpine as production
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/application .

# RUN apk add tzdata
ENV TZ=Asia/Bangkok
USER nonroot:nonroot

EXPOSE 8080


ENTRYPOINT ["/application"]