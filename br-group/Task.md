# Тестовое задание для Golang разработчика 

## Описание

Требуется реализовать подключение к WebSocket биржи [ascendex](https://ascendex.github.io/ascendex-pro-api/#websocket).<br>
Интерфейс подключения содержится в файле [apiclient.go](/br-group/apiclient.go)

+ **Connection** - функция должна реализовать подключение к WebSocket биржи.<br>
Если происходит проблемы с подключением, то должна вернуть ошибку
+ **Disconnect** - функция должна реализовывать отключение от WebSocket биржи.
+ **SubscribeToChannel** - функция должна реализовывать прослушивание канала WebSocket по получению BBO.<br>
Если происходит проблемы с прослушиванием, то должна вернуть ошибку и корректно отключиться от WebSocket.
+ **ReadMessagesFromChannel** - функция должна реализовывать чтение канала WebSocket о BBO, правильно преобразовывать данные и записывать их с chan.
+ **WriteMessagesToChannel** - функция должна реализовывать записи в канал WebSocket,<br>
чтобы поддерживать подключение открытым.

## Требования
+ Код выложен на [github](https://github.com/)
+ При решении использовать https://github.com/gorilla/websocket для работы с WebSocket
+ Написать тесты
