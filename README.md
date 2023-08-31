# Тестовое задание для стажёра Backend

## Установка

Используйте git для того, чтобы скачать проект из репозитория
```bash
git clone https://github.com/UhaBach/avitotest.git
```

## Запуск

Откройте терминал, перейдите в корневую папку проекта и введите следующие команды
```bash
docker-compose up --build -d # выполняем сборку
docker-compose start db # запускаем контейнер с бд
docker compose start server # запускаем контейнер с web-api
```
Теперь проект готов к работе.
Примечание: сервер запускается отдельно потому, что при запуске после билда попытка подключения к бд осуществляется во время ее инициализации и 
естественно получается connection refused.

## Запросы

В проекте определен следующий набор запросов:
![Image alt](https://github.com/UhaBach/avitotest/blob/pictures/image.png)

Перед началом работы, пожалуйста, выполните запрос "/hello", чтобы поздороваться с сприложением.

Запрос "/users/all" не принимает параметров и возвращает всех пользователей, без указания подключенных сегментов.
Запрос "/users/create/{name}" принимает имя юзера, который будет добавлен в бд, возвращает объект пользователя.
Запрос "/users/delete/{id:[0-9]+}" принимает в пользователя, которого надо удалить, ничего не возвращает.
Запрос "/users/{id:[0-9]+}" принимает id пользователя и возвращает информацию о нем, в ключая подключенные сегменты.
Запрос "/users/{id:[0-9]+}/change" принимает id пользователя, а также в теле запроса json соответствующий следующей структуре:
```bush
type ReqBody struct {
		AddSegment    []string
		RemoveSegment []string
}
```
возвращает объект пользователя с проведенными изменениями.
Для запросов "/segments/..." все аналогично, но вмесо id везде принимается name сегмента, а также запрос "/segments/{name}/change"
принимает в теле запроса json соответствующий следующей структуре:
```bush
type ReqBody struct {
		AddUser    []int
		RemoveUser []int
}
```
