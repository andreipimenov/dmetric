# Тестовое задание

Реализация приложения для мониторинга состояния устройств. Устройства посылают метрики. Требуется отслеживать их состояние вхождения в заданные диапазоны.
Информация о метриках и уведомлениях о выходе из диапазона сохраняется в БД Postgres.
Последнее уведомление о выходе из диапазона для каждого устройства сохранятеся в Redis.
Каждое уведомление о выходе из диапазона также отправляется письмом по электронной почте.

### API

Спецификация описана в формате OpenAPI, ссылка на файл спецификации <https://github.com/andreipimenov/dmetric/blob/master/doc/swagger.yaml>

Запуск локального SwaggerUI. Доступ через браузер на 127.0.0.1:3000
```
docker run -p 3000:8080 -e "API_URL=swagger.yaml" -v $(pwd)/doc/swagger.yaml:/usr/share/nginx/html/swagger.yaml swaggerapi/swagger-ui
```

#### Пример запроса к API
Отправка метрик от девайса 15
```
curl -X POST -d '{"metric_1":5, "metric_2":6, "metric_3":7, "metric_4":8, "metric_5":9, "local_time": "2018-04-06T13:00:00Z"}' 127.0.0.1:8080/api/v1/device/15/metrics

{"message":"OK"}
```

### Запуск

Для этого требуется установленный Docker, Docker-compose

```
docker-compose up
```

### Конфигурация

Приложение конфигурируется с помощью файла <https://github.com/andreipimenov/dmetric/blob/master/data/config.json>

При запуске внутри в качестве докер-контейнера изменение конфигураций возможно с помощью флагов, определяющих переменные окружения внутри контейнера. Обработкой переменных окружения и формированием обновленного конфигурационного файла занимается <https://github.com/andreipimenov/dmetric/blob/master/entrypoint.sh>