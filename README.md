# Описание api
## Авторизация - для api предусмотрен тип авторизации basic+e.
## Для пользования сервисом необходимо поместить access token доступа gitlab в конфигурационный файл gitlab_api_client.json

# Регистрация пользователя

## url - POST /api/gitlab/add_user?employee_id=1029871

### Входные данные

### Тело запроса:
```json
{
  "user_name": "mylonskiy.vladimir23",
  "first_name": "Мылонский",
  "last_name": "Владимир",
  "email": "mylonskiy.vladimir23@wildberries.work"
}
```
### Username пользователя должен быть указан латинскими буквами в нижнем регистре (цифры в конце допускаются).
### Почта пользователя должна оканчиваться на @wb.ru или @wildberries.work. Также должна быть указана латинскими буквами в нижнем регистре (цифры в почте допускаются).
### Пароль также генерится автоматически и возвращается в ответе

### Выходные данные в случае успеха:
```
{
  "data": {
    "main_info": {
      "name": "Мылонский Владимир",
      "username": "mylonskiy.vladimir23",
      "email": "mylonskiy.vladimir23@wildberries.work",
      "password": "tYHiBRsfzahlxMTA",
      "can_create_group": false,
      "skip_confirmation": false
    },
    "other_info": {
      "id": 2,
      "username": "mylonskiy.vladimir23",
      "email": "mylonskiy.vladimir23@wildberries.work",
      "name": "Мылонский Владимир",
      "state": "active",
      "locked": false,
      "avatar_url": "https://www.gravatar.com/avatar/2bea694be6d60e0a907701f5cb1d9edc?s=80&d=identicon",
      "web_url": "http://d6ecd115f788/mylonskiy.vladimir23",
      "created_at": "2024-02-28T12:01:20.503Z",
      "is_admin": false,
      "bio": "",
      "location": "",
      "skype": "",
      "linkedin": "",
      "twitter": "",
      "discord": "",
      "website_url": "",
      "organization": "",
      "job_title": "",
      "last_sign_in_at": null,
      "confirmed_at": null,
      "theme_id": 3,
      "last_activity_on": "",
      "color_scheme_id": 1,
      "projects_limit": 100000,
      "current_sign_in_at": null,
      "note": "",
      "identities": [],
      "can_create_group": false,
      "can_create_project": true,
      "two_factor_enabled": false,
      "external": false,
      "private_profile": false,
      "current_sign_in_ip": null,
      "last_sign_in_ip": null,
      "namespace_id": 2,
      "bot": false,
      "public_email": "",
      "shared_runners_minutes_limit": 0,
      "extra_shared_runners_minutes_limit": 0,
      "using_license_seat": false
    }
  }
}
```
### Возможные ошибки:
```
{
    "errors": [
        {
            "error": "validation.error",
            "message": "api validation error",
            "detail": "Имя должно быть написано кириллицей"
        }
    ]
}
```
```
{
  "errors": [
    {
      "error": "validation.error",
      "message": "api validation error",
      "detail": "Ошибка валидации: неверный формат для email - 'mylonskiy.vladimir23@mail.ru'"
    }
  ]
}
```
```
{
    "errors": [
        {
            "error": "internal.service_error",
            "message": "internal service error",
            "detail": "create user error"
        }
    ]
}
```

# Блокировка пользователя

## url - POST /api/gitlab/block_user?employee_id=1029871

### Входные данные

### Тело запроса:
```json
{
  "id": 1
}
```

### Возможные ошибки:
```
{
    "errors": [
        {
            "error": "internal.service_error",
            "message": "internal service error",
            "detail": "block user error"
        }
    ]
}
```
### В случае успеха ничего не возращается. Статус - 204.

# Получить всех пользователей с возможностью фильтрации
## url - GET /api/gitlab/get_users?page=1&employee_id=1029871

## Возможна сортировка по параметрам:
### 1) order_by - Возвращает пользователей, упорядоченных по полям id, name, username, created_at. По умолчанию id.
### 2) sort - Возвращает пользователей, отсортированных в порядке asc или desc. По умолчанию desc.
### 3) username - Получить одного пользователя с определенным username.
### 4) search - Поиск по имени или фамилии пользователя. Достаточно указать 1 из 2-ух параметров.
### 5) blocked - Фильтрует только заблокированных пользователей. Значение по умолчанию - false. Доступно только администраторам.
### 6) admins - Возвращает только администраторов. По умолчанию false. Доступно только администраторам.

## ! Обязательным параметром является - page (номер страницы). Возвращает до 20 пользователей.

## Пример: url - GET api/gitlab/get_users?order_by=name&sort=asc&page=1&employee_id=1029871

### Выходные данные

```json
 "data": [
{
"id": 157,
"username": "hamitov.a",
"email": "",
"name": "Александр Хамитов",
"state": "active",
"locked": false,
"avatar_url": "https://secure.gravatar.com/avatar/88709803dcba87afd1ee1fb135340edb?s=80&d=identicon",
"web_url": "https://wh-hard01.kol.wb.ru/hamitov.a",
"created_at": null,
"is_admin": false,
 ...
},
{
"id": 336,
"username": "alehin.aleksandr13",
"email": "",
"name": "Александр Алехин",
"state": "active",
"locked": false,
"avatar_url": "https://wh-hard01.kol.wb.ru/uploads/-/system/user/avatar/336/avatar.png",
"web_url": "https://wh-hard01.kol.wb.ru/alehin.aleksandr13",
"created_at": null,
"is_admin": false,
 ...
}]
```


# Получить пользователя по ID

## url - GET /api/gitlab/get_user?id=421&employee_id=1029871

### Выходные данные

```json
{
  "data": {
    "id": 421,
    "username": "baroyan.giorgiy",
    "email": "",
    "name": "Георгий Бароян",
    "state": "active",
    "locked": false,
    "avatar_url": "https://secure.gravatar.com/avatar/d222302f9515097358bc29409ad8ae90?s=80&d=identicon",
    "web_url": "https://wh-hard01.kol.wb.ru/baroyan.giorgiy",
    "created_at": "2023-10-31T10:56:17.457Z",
    "is_admin": false,
    "bio": "",
    "location": "",
    "skype": "",
    "linkedin": "",
    "twitter": "",
    "discord": "",
    "website_url": "",
    "organization": "",
    "job_title": "",
    "last_sign_in_at": null,
    "confirmed_at": null,
    "theme_id": 0,
    "last_activity_on": "",
    "color_scheme_id": 0,
    "projects_limit": 0,
    "current_sign_in_at": null,
    "note": "",
    "identities": [
      {"provider": "github", "extern_uid": "2435223452345"},
      {"provider": "bitbucket", "extern_uid": "john.smith"}
    ],
    "can_create_group": false,
    "can_create_project": false,
    "two_factor_enabled": false,
    "external": false,
    "private_profile": false,
    "current_sign_in_ip": null,
    "last_sign_in_ip": null,
    "namespace_id": 0,
    "bot": false,
    "public_email": "",
    "shared_runners_minutes_limit": 0,
    "extra_shared_runners_minutes_limit": 0,
    "using_license_seat": false
  }
}
```

### Возможные ошибки:
```
{
    "errors": [
        {
            "error": "validation.error",
            "message": "api validation error",
            "detail": "cannot bind query for request, err: Key: 'ID' Error:Field validation for 'ID' failed on the 'required' tag"
        }
    ]
}
```
```
{
    "errors": [
        {
            "error": "internal.service_error",
            "message": "internal service error",
            "detail": "get user error"
        }
    ]
}
```


# Проверка валидности access token
## url - GET /api/gitlab/check_valid_token
### В случае успеха ничего не возращается. Статус - 204.