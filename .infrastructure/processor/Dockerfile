FROM golang:1.20 AS base

WORKDIR /build

# инфраструктура: конфиги и csv файлы.
ADD ./.infrastructure ./.infrastructure

# main сервиса и бизнес-логика
ADD ./cmd/processor ./cmd/processor
ADD ./internal ./internal

# go.mod && go.sum
ADD ./go.* ./

# Собираем бинарник
RUN CGO_ENABLED=0 go build -o processor ./cmd/processor/

FROM scratch AS runtime

WORKDIR /app

COPY --from=base /build/processor ./
COPY --from=base /build/.infrastructure/config.yml ./
COPY --from=base /build/.infrastructure/data ./data

CMD ["./processor"]