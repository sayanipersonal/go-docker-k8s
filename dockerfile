# build stage
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# final stage
FROM alpine:latest AS final
WORKDIR /root/
COPY --from=build /app/server .
EXPOSE 8080
CMD [ "./server" ]

