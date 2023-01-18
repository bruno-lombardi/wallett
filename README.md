# Wallett

Simple HTTP REST API that helps you take control of your money, manage multiple
resources, like: `users`, `wallets` and `transactions`.

Currently supported database is SQLite. Any SQL database should work.

## Endpoints

All endpoints are available below.

<details>
  <summary>POST /api/v1/users</summary>
  
  Creates an user with email and password. Example:
```sh
curl --request POST \
  --url http://localhost:3333/api/v1/users \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "bruno.vbl@hotmail.com",
	"name": "Bruno Lombardi",
	"password": "123456",
	"password_confirmation": "123456"
}'
```
201 CREATED:
```json
{
	"id": "u_uZpAwRNITM",
	"email": "bruno.vbl@hotmail.com",
	"name": "Bruno Lombardi"
}
```
</details>

<details>
  <summary>GET /api/v1/users</summary>
  
  
List all users. Example:

```sh
curl --request GET \
  --url 'http://localhost:3333/api/v1/users?Page=1&Limit=10'
```

200 OK:

```json
{
  "total_pages": 1,
  "count": 1,
  "per_page": 10,
  "page": 1,
  "data": [
    {
      "id": "u_uZpAwRNITM",
      "email": "bruno.vbl@hotmail.com",
      "name": "Bruno Lombardi"
    }
  ]
}
```

</details>

<details>
  <summary>GET /api/v1/users/{user_id}</summary>

Fetch user by id. Example:

```sh
curl --request GET \
  --url http://localhost:3333/api/v1/users/u_uZpAwRNITM
```

200 OK:

```json
{
  "id": "u_uZpAwRNITM",
  "email": "bruno.vbl@hotmail.com",
  "name": "Bruno Lombardi"
}
```

</details>

## Install dependencies

```bash
➜ go mod download
```

## Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```
