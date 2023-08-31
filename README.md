# avito-user-segment-service
"Сервис управления сегментами пользователей: создание, добавление и удаление сегментов, а также просмотр активных сегментов пользователя"

## Сервис динамического сегментирования пользователей

Этот сервис предоставляет API для управления сегментами пользователей.

### Метод создания сегмента:

**URL:** `/createSegment`

**Метод:** POST

**Параметры запроса:**
- `slug` (string, обязательный) - Название нового сегмента.

**Пример запроса:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/53bc8718-499c-4484-b83e-9c4ddb5deec6)

**Результат ответа:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/0fbfa1a3-bf61-4c81-86cc-152dbbfc8794)



### Метод удаления сегмента:

**URL:** `/deleteSegment`

**Метод:** POST

**Пример запроса:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/563d9267-4d28-4000-90fc-91a186ff7200)

**Результат ответа:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/12df76f8-9314-480e-9e29-929cd8e4753a)


### Метод получения активных сегментов пользователя:

**URL:** `/getUserSegments`

**Метод:** GET

**Пример запроса:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/aaced72d-4bcf-44a0-bcfd-a50e0bdc5b94)

**Результат ответа:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/2bf55ad3-8eb2-4d41-966b-2af113258dff)

### Метод добавления пользователя в сегмент:

**URL:** `/getUserSegments`

**Метод:** POST

**Пример запроса:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/7fa86dff-7b42-4cdf-b9e5-a8b1ec653e31)

**Результат ответа:**

![image](https://github.com/rinatziatdinov/avito-user-segment-service/assets/128255454/ea9cafd8-6a79-4bb2-b26b-9812d0b882be)
