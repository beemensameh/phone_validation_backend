# Phone Validation

This app is a small example of how can validate phone number from sqlite3

## Tech Stack

**Server:** Go v1.16

**Database:** Sqlite3

## Installation

before you start steps you will need docker and docker compose for this. If you don't have, you need to install them first.

1. Just run these commands and the server will start automatically.

```bash
  $ docker-compose build
  $ docker-compose up
```

2. After the dockers start you can navigate to `127.0.0.1:1234` to see the app home page.
3. You will find a postman collection in the root directory in the project. Just import them.

## API Reference

### Get all countries code

```http
  GET /api/v1/countries-code
```

### Get all customers

```http
  GET /api/v1/customers?country_code=:code&state=:state
```

| Parameter | Type     | Description                                    |
| :-------- | :------- | :--------------------------------------------- |
| `code`    | `string` | **Optional**. The country code                 |
| `state`   | `string` | **Optional**. The state (`valid` or `invalid`) |

## Screenshots
![Screenshot from 2022-02-05 17-54-52](https://user-images.githubusercontent.com/41084077/152649235-b423c390-5e5b-402e-8c4b-d90fd0737f4f.png)
![Screenshot from 2022-02-05 17-55-01](https://user-images.githubusercontent.com/41084077/152649236-9942ff09-7e9e-4eef-b8a4-202c326da3af.png)
