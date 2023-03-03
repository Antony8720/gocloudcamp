# Тестовое задание для поступления в GoCloudCamp
Техническое задание: https://github.com/gocloudcamp/test-assignment  

## Ответы на вопросы   

1. Опишите самую интересную задачу в программировании, которую вам приходилось решать?  
    - Разработка сервиса с кулинарными рецептами, с возможностью публиковать рецепты, подписываться  
    на авторов, скачивать список покупок в файл.
2. Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?  
    - Не успел доделать тестовое задание, потому что неправильно рассчитал время на выполнение.  
    Теперь закладываю на задание в 2 раза больше времени на выполнение.
3. Каковы ваши ожидания от участия в буткемпе?  
    - Получить опыт разработки на реальных проектах, прокачать свои навыки программирования на Go, попасть в штат компании по окончании стажировки.  

## Стек
`Golang`, `gRPC`, `PostgreSQL`, `Docker`.

## Описание  
В сервисе реализовано:
1. Модуль для работы с плейлистом (Методы Start, Pause, Next, Prev, Stop, AddSong).
2. CRUD для песен (Методы AddSong, ShowSong, UpdateSong, DeleteSong).
3. gRPC сервер для управления плейлистом.
4. Клиент для тестирования сервера.
5. База данных PostgreSQl для хранения плейлиста.
6. Сервис и БД развернуты в docker контейнерах и могут быть запущены через docker-compose.

## Запуск проекта

1. Клонировать репозиторий.

```
git clone git@github.com:Antony8720/gocloudcamp.git
```

2. Запустить контейнеры.
```
docker compose up
```