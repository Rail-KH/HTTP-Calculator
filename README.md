# Веб-сервис для вычисления арифметических выражений

Этот проект реализует веб-сервис, принимающий выражение через HTTP запрос и возвращающий результат вычислений 

# Структура проекта

cmd/ — точка входа приложения.

internal/ — внутренняя логика и модули приложения.

pkg/ — вспомогательные пакеты.

# Инструкция по запуску:

Убедитесь, что у вас установлен Go (желательно версия 1.23.1 или выше).

Скопируйте репозиторий:

git clone https://github.com/Rail-KH/HTTP-Calculator

cd HTTP-Calculator


Запустите сервер:

go run ./cmd/main.go

Сервер будет доступен по адресу http://localhost:8080/api/v1/calculate


# Примеры использования(запрос через командную строку):

1) Успешный запрос:

   Запрос: curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}""

   Ответ: {"result":"2"}

3) Ошибка 422 (невалидное выражение):
   
   Запрос: curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1*\"}""
   
   Ответ: {"error":"Expression is not valid"}

4) Ошибка 500 (неизвестная ошибка сервера):
   
   Ответ: {"error":"Internal server error"}
