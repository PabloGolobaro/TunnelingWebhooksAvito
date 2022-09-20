FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o tg_bot .

FROM scratch AS production

COPY --from=build /src/tg_bot .

#COPY --from=build /src/bot_config.yaml .

COPY --from=build /src/clients.db .

EXPOSE 8080

CMD ["/tg_bot"]