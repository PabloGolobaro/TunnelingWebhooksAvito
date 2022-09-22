FROM golang:1.18 as build

WORKDIR /src

COPY . /src

RUN CGO_ENABLED=0 GOOS=linux go build -o bot

FROM alpine

COPY --from=build /src/bot .

COPY --from=build /src/bot_config.yaml .

EXPOSE 8080

CMD ["/bot"]