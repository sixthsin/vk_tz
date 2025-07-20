Микросервис для размещения и управления объявлениями с аутентификацией пользователей.

## Запуск проекта

### Требования
- Docker
- Docker-compose

### 1. Настройка окружения

Создайте файл `.env` в корне проекта:

Настройте файлы окружения, скопировав их из .env.example
```bash
cp .env.example .env
```
Например
```bash
AUTH_SERVICE_DSN="host=db user=postgres password=my_pass dbname=marketplace_api_db port=5432 sslmode=disable"
AUTH_SERVICE_REST_PORT="8081"
AUTH_DB_NAME="go_api_auth"
AUTH_DB_USER="postgres"
AUTH_DB_PASSWORD="my_pass"
JWT_SECRET="secret"
```

Запустите Docker-compose файл
```bash
docker-compose up -d
```

1. Регистрация
   ```bash
   POST /api/v1/auth/register
   {
    "username":"username",
    "password":"PassWord123"
   }
   ```

   Ответ
   ```bash
   {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTMxMjA4OTgsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.HdapgMMjFfNr_0kXIaPusxVSdelLJxV0gCBSzkfs9DI",
    "username": "username"
   }
   ```

2. Логин
   ```bash
   POST /api/v1/auth/login
   {
    "username":"username",
    "password":"PassWord123"
   }
   ```

   Ответ
   ```bash
   {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTMxMjA5MDIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.Hjar78kP7xB4O5hhEdNreUUxDxZd-5fl3hxDyDLfe20"
   }
   ```

3. Создание объявления
   ```bash
   POST /api/v1/adverts/
   {
    "title": "title",
    "description": "description",
    "image_url": "https://photo.com/photo.jpeg",
    "price": 5
   }
   ```

   Ответ
   ```bash
   {
    "advert": {
        "title": "title",
        "description": "description",
        "image_url": "https://photo.com/photo.jpeg",
        "price": 5,
        "author": "username"
    },
    "status": "ok"
    }
   ```
   
6. Получение списка объявлений
   ```bash
   GET /api/v1/adverts/?limit=0&offset=0&sort_by=created_at&sort_order=desc&min_price=0&max_price=0
   ```

   Ответ
   ```bash
   {
    "data": [
        {
            "id": "13",
            "created_at": "2025-07-20T18:06:00.716858Z",
            "is_mine": false,
            "title": "MacBook Pro 16\"",
            "description": "M1 Max, 32GB RAM, 1TB SSD",
            "image_url": "https://example.com/macbook.jpg",
            "price": 250000,
            "author": "tech_seller"
        },
        {
            "id": "12",
            "created_at": "2025-07-20T17:44:48.666433Z",
            "is_mine": true,
            "title": "Sony PlayStation 5",
            "description": "Новая, с 2 играми",
            "image_url": "https://example.com/ps5.jpg",
            "price": 55000,
            "author": "username"
        },
        {
            "id": "11",
            "created_at": "2025-07-20T17:44:46.950753Z",
            "is_mine": false,
            "title": "Квартира в центре",
            "description": "2-комнатная, 60 кв.м.",
            "image_url": "https://example.com/flat.jpg",
            "price": 12000000,
            "author": "realtor"
        },
        {
            "id": "10",
            "created_at": "2025-07-20T17:44:45.138677Z",
            "is_mine": true,
            "title": "Велосипед горный",
            "description": "Trek Marlin 5, 2024 г.",
            "image_url": "https://example.com/bike.jpg",
            "price": 45000,
            "author": "username"
        },
        {
            "id": "9",
            "created_at": "2025-07-20T17:44:42.523066Z",
            "is_mine": false,
            "title": "Apple Watch Series 9",
            "description": "45mm, GPS+Cellular",
            "image_url": "https://example.com/watch.jpg",
            "price": 45000,
            "author": "gadget_lover"
        },
        {
            "id": "8",
            "created_at": "2025-07-20T17:42:32.660216Z",
            "is_mine": true,
            "title": "Фотоаппарат Canon EOS R5",
            "description": "Комплект с объективом",
            "image_url": "https://example.com/camera.jpg",
            "price": 320000,
            "author": "username"
        },
        {
            "id": "7",
            "created_at": "2025-07-20T17:42:32.122571Z",
            "is_mine": false,
            "title": "Диван угловой",
            "description": "Новый, нераспакованный",
            "image_url": "https://example.com/sofa.jpg",
            "price": 75000,
            "author": "furniture_seller"
        },
        {
            "id": "6",
            "created_at": "2025-07-20T17:42:31.541028Z",
            "is_mine": true,
            "title": "iPhone 15 Pro",
            "description": "256GB, Space Black",
            "image_url": "https://example.com/iphone.jpg",
            "price": 120000,
            "author": "username"
        },
        {
            "id": "5",
            "created_at": "2025-07-20T17:42:30.701882Z",
            "is_mine": false,
            "title": "Кольцо обручальное",
            "description": "Золото 585 пробы",
            "image_url": "https://example.com/ring.jpg",
            "price": 45000,
            "author": "jeweler"
        },
        {
            "id": "4",
            "created_at": "2025-07-20T17:42:30.189659Z",
            "is_mine": true,
            "title": "Ноутбук ASUS ROG",
            "description": "RTX 4080, 32GB RAM",
            "image_url": "https://example.com/asus.jpg",
            "price": 210000,
            "author": "username"
        }
    ],
    "meta": {
        "limit": 10,
        "offset": 0,
        "total": 10
    },
    "status": "ok"
   }
   ```
