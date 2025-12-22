# syntax=docker/dockerfile:1
FROM golang:1.25-alpine as build
RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 go build -o /bin/app ./main.go

FROM alpine:latest
COPY --from=build /bin/app /bin/app
EXPOSE 8080
CMD ["/bin/app"]