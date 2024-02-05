## Лабораторная работа 1. Работа с http протоколом.

Задача

Передать http GET запрос напрямую из любого приложения (программного кода), но не из браузера. То есть передать «без мусора», только то что надо.

Сервер: http://109.167.241.225/http_example/ \
Порт: 8001 \
GET запрос: give_me_five

Параметры запроса: \
wday — день недели (вс = 1) \
student — номер в таблице 

Заголовки: \
REQUEST_AGENT = ITMO student \
COURSE = Protocols

В заголовках больше ничего не должно быть. \
Сервер присылает ответ — структура из числа и строки. \
Строка = «бинарное» представление ещё одного числа.

700/TCP,UDP EPP (Extensible Provisioning Protocol) — используется для управления регистрационной информацией DNS Официально

Пример запроса без параметров (2 возврата строки в конце, поскольку 2 параметра):
```GET /http_example/give_me_five?student=30&wday=5 HTTP/1.1```

Неправильный ответ на этот запрос:
```
HTTP/1.1 409 Conflict
Connection: Keep-Alive
Server: Embedthis-http
Content-Type: text/html
Cache-Control: no-cache
Date: Ср, 08 фев 2023 20:14:51 GMT
Content-Length: 0
Keep-Alive: timeout=60, max=199
```

Ответ на правильный запрос:
```
HTTP/1.1 200 OK
Connection: Keep-Alive
Server: Embedthis-http
Content-Type: application/json; charset=utf-8
Cache-Control: no-cache
Date: Ср, 08 фев 2023 20:28:20 GMT
Content-Length: 109
Keep-Alive: timeout=60, max=199

{"workResult":{"value":1008,"strMessage":"0001000000000000000100000000000000010000000000000001000000000000"}}
```

После получения правильного результата его необходимо отправить  для проверки в личном сообщении ТГ. После проверки, начинать готовить отчет.

Отчет выслать в течение 2-х недель на адрес akharitonov@itmo.ru. В теме письма: №группы ФИО (латинскими буквами) №работы (например: 5555 Fedor Sumkin 1).
