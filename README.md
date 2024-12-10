# Change Data Capture (CDC) in Nats JetStream

Пример использования Debezium с Postgres и Nats JetStream **без использования фреймворка Apache Kafka Connect**.

В данном примере используется Debezium для отслеживания изменений в базе данных Postgres и отправки их в Nats JetStream.

## Описание

Проект включает:
• Debezium для отслеживания изменений в Postgres.
• Nats JetStream для обработки и доставки событий.
• Настройка Debezium-коннектора для Postgres с базовой трансформацией данных. Конфигурация находится в
файле [application.properties](debezium/application.properties).
• Скрипт и API для создания тестовых данных.

## Запуск проекта

Для запуска используйте Docker Compose. Выполните команду:

```bash
docker-compose up -d --build
```

Для остановки и удаления контейнеров с их данными выполните:

```bash
docker-compose down -v
```

## Создание тестовых данных

Для создания начальной схемы и данных используется SQL-скрипт [init.sql](postgres/init.sql). Скрипт автоматически
выполняется при первом запуске контейнера Postgres.

## Добавление данных через API

Для создания дополнительных тестовых данных в таблице **public.products** можно воспользоваться REST-методом. Выполните
запрос:

```bash
curl -X POST http://localhost:8081/products
```

## References

1) [Debezium documentation](https://debezium.io/documentation/reference/stable/index.html)
2) [Debezium connector for Postgres](https://debezium.io/documentation/reference/stable/connectors/postgresql.html)
3) [Debezium Sink configuration for Nats](https://debezium.io/documentation/reference/stable/operations/debezium-server.html#_nats_jetstream)
4) [Nats JetStream documentation](https://docs.nats.io/nats-concepts/jetstream/streams)
5) [Nats CLI](https://docs.nats.io/using-nats/nats-tools/nats_cli)