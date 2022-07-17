# Word of Wisdom

## Требования

Test task for Server Engineer

Design and implement “Word of Wisdom” tcp server.  
• TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work),
the challenge-response protocol should be used.  
• The choice of the POW algorithm should be explained.  
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other
collection of the quotes.  
• Docker file should be provided both for the server and for the client that solves the POW challenge

## Пример запуска

Данные о настройках можно передать как через параметры командной строки,
так и через переменные окружения (далее будут из названия указаны в скобках рядом с параметром).

Запуск сервера:

```
go run ./cmd/server/ --addr=:12898 --complexity=6
```

- `addr` (`ADDR`) адрес порта на котором запустится TCP сервер
- `complexity` (`COMPLEXITY`) сложность алгоритма для Proof of Work

Запуск клиента:

```
go run ./cmd/client --host=0.0.0.0:12345
```

- `host` (`HOST`) - адрес сервера

В докере:

```
docker-compose -f ./build/local/docker-compose.yml up  
```