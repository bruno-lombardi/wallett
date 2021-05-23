# Wallett

Restful API that helps you take control of your money.

### Endpoints

Resource name is written in the plural form.

|Endpoint|Description|Example (User resource)|
|--------|-----------|-------|
|`GET /api/v1/<resources>`|List items|`GET /api/v1/users` List users|
|`POST /api/v1/<resources>`|Create new item|`POST /api/v1/users` Create new user|
|`GET /api/v1/<resources>/{id}`|Retrieve the item|`GET /api/v1/users/u_iWvqZTa2N9` Get the user which ID is u_iWvqZTa2N9|
|`PUT /api/v1/<resources>/{id}`|Update the item|`PUT /api/v1/users/u_iWvqZTa2N9` Update the product which ID is u_iWvqZTa2N9|
|`DELETE /api/v1/<resources>/{id}`|Delete the item|`DELETE /api/v1/users/u_iWvqZTa2N9` Delete the product which ID is u_iWvqZTa2N9|


### Install dependencies

```bash
➜ go mod download
```

### Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```