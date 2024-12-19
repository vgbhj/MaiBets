# MaiBets

# TODO
- [x] Добавить в миграции создание таблицы
- [x] При добавлении ивентов рандомно генерируются кефы
- [x] Добавление ставки
- [x] Добавление ивентов (рандом длительность)
- [x] Добавить документацию для API
- [ ] При получении награды учитывать множитель ставки
- [x] Гет запросы на получение всех ставок пользователя
- [x] Гет запросы на получение всех ивентов
- [ ] На стороне базы данных необходимо определить представления, триггеры, функции и хранимые процедуры,
    - [x] Триггер на закрытие ивента, когда настает его время
    - [ ] Триггер на закрытие ставки, когда ивент закрывается и рассет стаки
# Future todo
- [x] Добавить админа в миграции
- [x] Добавлять события можно только админ

---

## DB model
![dbmodel.png](docs/dbmodel.png)

## Api docs (swagger)

Для получение документации нужно запустить контейнер приложения
```bash
docker-compose up --build
```
И перейти по: http://localhost:8080/swagger/index.htm