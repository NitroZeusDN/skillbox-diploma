FROM golang:1.20 AS base

WORKDIR /build

# инфраструктура: конфиги и csv файлы.
ADD ./.infrastructure ./.infrastructure

# main сервиса
ADD ./cmd/simulator ./cmd/simulator
ADD ./internal/config ./internal/config

# go.mod && go.sum
ADD ./go.* ./

# Собираем бинарник
RUN CGO_ENABLED=0 go build -o simulator ./cmd/simulator/

FROM scratch AS runtime

WORKDIR /app

COPY --from=base /build/simulator ./
COPY --from=base /build/.infrastructure/config.yml ./
COPY --from=base /build/.infrastructure/data ./data

CMD ["./simulator"]