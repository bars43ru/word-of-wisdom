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

## Фитбэк
Смущает то, как реализована конкурентность в сервере. Сильно упрощенный протокол (соединение закрывается с случае ошибки) - нет возможности послать solution в другом соединении. Использовал 3rd-party библиотеку для работы с флагами, но в тестовом хотелось бы видеть максимум примитивов.

## Пример запуска

Данные о настройках можно передать как через параметры командной строки,
так и через переменные окружения (далее будут из названия указаны в скобках рядом с параметром).

Запуск сервера:

```
go run ./cmd/server/ --addr=:12345 --complexity=6
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

## Алгоритм

Для реализации Proof of Work был взят за основу алгоритм hashcash. 
Для каждого открытого соединения создается псевдослучайный набор байтов, 
которые отправляется клиенту с параметрами поиска хэша. Клиент перебором 
вычисляет хэш с заданными параметрами и результат перебора отправляет на сервер, 
где при небольших вычислительных ресурсах можно проверить правильность решения задачи.

## Протокол

### Формат
Протокол обмена сообщений был реализован только для этапа проведения proof of work, 
после окончания процедуры, строка с цитатой просто записывается в сокет.

Формат сообщений для этапа proof of work:
|`uint64`|`[]byte`|`...`|`uint64`|`[]byte`|


Т.е. значением записан в бинарном виде размер сообщения, дальше за ним идет json объект 

### Workflow

1. Сервер отправляет набор данных и параметры для поиска хэш
2. Клиент подбирает по заданным параметрам хэш
3. Клиент отправляет подобранные добавочные данные на севре
4. Сервер проверяет полученные данные от клиента
5. В случае успешной проверки, уведомляет об этом клиента. При не успешном результате, разрывает соединение.
