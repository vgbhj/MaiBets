# MaiBets

# TODO
- [x] Добавить в миграции создание таблицы
- [x] При добавлении ивентов рандомно генерируются кефы
- [x] Добавление ставки
- [ ] Гет запросы на получение всех ставок пользователя
- [ ] Гет запросы на получение всех ивентов
- [ ] На стороне базы данных необходимо определить представления, триггеры, функции и хранимые процедуры,
    - [ ] Триггер на закрытие ивента, когда настает его время
    - [ ] Триггер на закрытие ставки, когда ивент закрывается и рассет стаки
    - [ ] Триггер на удаление кефов, когда ивент закрывается
- [ ] Добавление ивентов (рандом длительность)
- [ ] Добавить документацию для API

# Future todo
- [x] Добавить админа в миграции
- [ ] Добавить проверку на админа в мидлваре?
- [ ] Добавлять события можно только админ

---

## DB model
![dbmodel.png](docs/dbmodel.png)

## Curl запросы

* добавление ивента
```bash
curl -X POST http://localhost:8080/api/event \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-d '{
    "name": "Название события",
    "date": "2023-10-01T15:00:00Z",
    "location": "Место проведения",
    "description": "Описание события"
}'
```

* логин за админа
```bash
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{"username": "admin", "password": "admin"}'
```

* добавление ставки
```bash
curl -X POST http://localhost:8080/api/bet \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ1NDYzODksImlkIjo1fQ.AQOBJfxLfMmLsnyd6Wagkamz_gVWi2bjLhaUIp61f5s" \
-d '{
    "name": "Название события",
    "bet_amount": 10,
    "bet_date": "2023-10-01T15:00:00Z"
}'
```