# JSON:
- Create имеет формат:
{
"name": "Leo",
"surname": "Ushakov",
"patronymic": "Vasilevich" 
}

- Delete имеет формат:
{
"user_id": "af9062ad-7303-409d-bf58-a57fca27c73a"
}

# Порядок запуска с миграцией:
1. docker-compose up
2. make migrate-create
  пишем схему таблицs в образовавшихся файлах в migrations
3. make migrate-up
Приложение готово к работе  

# Для удобного доступа в Postgres -- добавлен контейнер Adminer  