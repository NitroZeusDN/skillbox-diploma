# Skillbox diploma
Итоговая работа по курсу Skillbox Golang.

## Требования:
- Установленный `Docker` на вашем компьютере
- Установленный `Make` на вашем компьютере

## Конфигурация

Если ваши порты заняты, то можете изменить файл конфигурации.
Файл конфигурации [config.yaml](.infrastructure/config.yml) в директории [.infrastructure](.infrastructure)

## Использование

> **_Примечание:_**  Если вы хотите изменения данных при каждом запуске, 
>  то требуется изменить правила записи `volumes` в [docker-compose.yml](.infrastructure/docker-compose.yml).
>  По стандарту используется правило `READ-ONLY`.

Запуск линтера GolangCI (правила расположены [здесь](.golangci.yml):
```bash
make lint
```

Запуск симулятора и обработчика данных:
```bash
make up
```

Отправьте `GET` запрос на адрес `http://localhost:8484/`, где в результате будет
JSON ответ как в [примере](output/1.json).

Остановка симулятора и обработчика данных:
```bash
make down 
``` 

## Сторонние пакеты:

- Go Chi [[link]](https://github.com/go-chi/chi)
- Gorilla Mux [[link]](https://github.com/gorilla/mux)
- Clean Env [[link]](https://github.com/ilyakaznacheev/cleanenv)