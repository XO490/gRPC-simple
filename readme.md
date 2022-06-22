gRPC-servers
================

### Задача
В общем виде сервис должен предоставлять внешние gRPC-методы (эндпоинты) для выполнения следующих операций
Операции для управления “продуктами” - добавление, изменение, удаление, получение по уникальному идентификатору, получение всего списка (с сортировкой по названию в алфавитном порядке, с сортировкой по времени создания (в т.ч. в обратном порядке)

---

### Структура одной сущности
- ID - первичный ключ
- Name
- Timestamp - таймстамп создания сущности (time.Time{})
- Price
- Description


- [x] получение по id
- [x] получение по name
- [x] получение по created_at `!есть вопросы по timestamp`
- [ ] добавление
- [ ] изменение
- [ ] удаление
- [x] Получение всего списка продуктов
  - [x] по уникальному идентификатору
  - [x] по уникальному идентификатору  | реверс
  - [x] по названию в алфавитном порядке
  - [x] по названию в алфавитном порядке  | реверс
  - [x] с сортировкой по времени создания
  - [x] с сортировкой по времени создания | реверс
---
```
grpcurl -plaintext localhost:9092 describe

grpc.reflection.v1alpha.ServerReflection is a service:
service ServerReflection {
  rpc ServerReflectionInfo ( stream .grpc.reflection.v1alpha.ServerReflectionRequest ) returns ( stream .grpc.reflection.v1alpha.ServerReflectionResponse );
}
proto.Product is a service:
service Product {
  rpc GetAllByCreated ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetAllByCreatedDesc ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetAllById ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetAllByIdDesc ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetAllByName ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetAllByNameDesc ( .proto.GetByRequest ) returns ( .proto.GetAllByIdResponse );
  rpc GetBy ( .proto.GetByRequest ) returns ( .proto.GetByResponse );
}
```
---
```
grpcurl -plaintext -d '{"id": "4"}' localhost:9092 proto.Product.GetBy

    {
        "id": "4",
        "name": "NFT-4",
        "createdAt": "1654837200",
        "price": "0.94",
        "description": "bla bla 4"
    }
```
---
```
grpcurl -plaintext -d '{"name": "NFT-2"}' localhost:9092 proto.Product.GetBy

    {
        "id": "2",
        "name": "NFT-2",
        "createdAt": "1654661408",
        "price": "547",
        "description": "bla bla 2"
    }
```
---
```
grpcurl -plaintext localhost:9092 proto.Product.GetAllById

{
  "result": [
    {
      "id": "1",
      "name": "nft-1",
      "createdAt": "1654355386",
      "price": "4",
      "description": "bla bla 1"
    },
    {
      "id": "2",
      "name": "NFT-2",
      "createdAt": "1654661408",
      "price": "547",
      "description": "bla bla 2"
    },
    {
      "id": "3",
      "name": "NFT-3",
      "createdAt": "1655310000",
      "price": "76",
      "description": "bla bla 3"
    },
    {
      "id": "4",
      "name": "NFT-4",
      "createdAt": "1654837200",
      "price": "0.94",
      "description": "bla bla 4"
    }
  ]
}
```
---
```
grpcurl -plaintext localhost:9092 proto.Product.GetAllByIdDesc

{
  "result": [
    {
      "id": "4",
      "name": "NFT-4",
      "createdAt": "1654837200",
      "price": "0.94",
      "description": "bla bla 4"
    },
    {
      "id": "3",
      "name": "NFT-3",
      "createdAt": "1655310000",
      "price": "76",
      "description": "bla bla 3"
    },
    {
      "id": "2",
      "name": "NFT-2",
      "createdAt": "1654661408",
      "price": "547",
      "description": "bla bla 2"
    },
    {
      "id": "1",
      "name": "nft-1",
      "createdAt": "1654355386",
      "price": "4",
      "description": "bla bla 1"
    }
  ]
}
```
---
```
grpcurl -plaintext localhost:9092 proto.Product.GetAllByNameDesc

{
  "result": [
    {
      "id": "4",
      "name": "NFT-4",
      "createdAt": "1654837200",
      "price": "0.94",
      "description": "bla bla 4"
    },
    {
      "id": "3",
      "name": "NFT-3",
      "createdAt": "1655310000",
      "price": "76",
      "description": "bla bla 3"
    },
    {
      "id": "2",
      "name": "NFT-2",
      "createdAt": "1654661408",
      "price": "547",
      "description": "bla bla 2"
    },
    {
      "id": "1",
      "name": "nft-1",
      "createdAt": "1654355386",
      "price": "4",
      "description": "bla bla 1"
    }
  ]
}
```
---
```
grpcurl -plaintext localhost:9092 proto.Product.GetAllByCreated

{
  "result": [
    {
      "id": "1",
      "name": "nft-1",
      "createdAt": "1654355386",
      "price": "4",
      "description": "bla bla 1"
    },
    {
      "id": "2",
      "name": "NFT-2",
      "createdAt": "1654661408",
      "price": "547",
      "description": "bla bla 2"
    },
    {
      "id": "4",
      "name": "NFT-4",
      "createdAt": "1654837200",
      "price": "0.94",
      "description": "bla bla 4"
    },
    {
      "id": "3",
      "name": "NFT-3",
      "createdAt": "1655310000",
      "price": "76",
      "description": "bla bla 3"
    }
  ]
}
```
---

- [ ] Установка и растройка инфраструктуры для тестового задания
  - [x] СУБД PostgreSQL
  - [ ] migrate-tool (для создания и отката миграций, избегать автомиграций)
  - [ ] Bun ORM (построение и отправка SQL-запросов )
  - [x] Protobuf
  - [ ] know-types
  - [x] Кодогенерация: protoc
  - [x] gRPCURL (для вз-ия с gRPC-сервером)

- [ ] Реализация задания
  - [x] Реализовать .proto структуру для сериализации и десериализации данных, кодогенерации (gRPC, go-types, gateway, proto-descriptors)
  - [x] Создание и запуск gRPC-сервера (tcp-listener).  
  - [x] Реализовать структуру данных и методы управления сущностями.
  - [ ] Реализовать для каждого базового метода управления сущностями внешний gRPC-метод для удаленного вызова и обращения к gRPC-серверу.