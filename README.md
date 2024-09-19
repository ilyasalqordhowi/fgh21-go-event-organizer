# SnagTick BACKEND

Let me introduce my awesome online-based ticket booking, SnagTick. This repository utilized PGX for package and Gin-gonic as framework API for development.

Tickhub is an easy way to order an event from a distance place. Without order on the spot, we can choose the event as we want with the simple process and various payment. Also, we can save the listed event as the wishlist.

# Tech Stack

- Backend's Programming Language: Go
- Framework : Gin-Gonic
- Package Manager : Golang PGX
- Data Migration : Golang Migrate
- RDBMS : PostgreSQL
- API Testing : ThunderClient
- Containerization : Docker

# Config / installation process

## 1. Clone this repository

```sh
  git clone https://github.com/ilyasalqordhowi/fgh21-go-event-organizer.git
  cd <project-name>
```

## 2. Open in VSCode

```sh
  code .
```

## 3. Install all the dependencies

```sh
  go mod tidy
```

## 4. Run the program

```sh
  go run main.go
```

# API References

## Login

```http
  POST auth/login
```

## Register

```http
  POST auth/register
```

| Parameter               | Type     | Description                                            |
| :---------------------- | :------- | :----------------------------------------------------- |
| `users`                 | `GET`    | `Get a list of users data`                             |
| `users/:id`             | `GET`    | `Select the user data according to registered id`      |
| `users`                 | `POST`   | `Create new user data`                                 |
| `users/update`          | `PATCH`  | `Edit the selected user data`                          |
| `users/:id`             | `DELETE` | `Remove the selected user data`                        |
| `users/password`        | `PATCH`  | `Change the user's password`                           |
| `events`                | `GET`    | `Get a list of events data`                            |
| `events/:id`            | `GET`    | `Select the event data according to registered id`     |
| `events/`               | `POST`   | `Create new event`                                     |
| `events/:id`            | `PATCH`  | `Edit the selected event data`                         |
| `events/:id`            | `DELETE` | `Remove the selected event data`                       |
| `events/payment_method` | `GET`    | `Get a list of payment methods data`                   |
| `events/section/:id`    | `GET`    | `Get a list of event sections data`                    |
| `events/data`           | `GET`    | `Get a list of detail new event  data`                 |
| `transactions`          | `POST`   | `Create new transactions`                              |
| `transactions`          | `GET`    | `Get a list of transactions by registered user`        |
| `profile`               | `GET`    | `Select the profile data according to registered user` |
| `profile/update`        | `PATCH`  | `Change the profile data from registered user`         |
| `profile/img`           | `PATCH`  | `Change the profile's image  from registered user`     |
| `profile/national`      | `GET`    | `Get a list of nationalities data`                     |
| `profile/national/:id`  | `GET`    | `GSelect the national data according to registered id` |
| `categroies`            | `GET`    | `Get a list of categories data`                        |
| `categroies/:id`        | `GET`    | `Select the category data according to registered id`  |
| `categroies`            | `POST`   | `Create new category data`                             |
| `categroies/:id`        | `PATCH`  | `Edit the selected category data`                      |
| `categroies/:id`        | `DELETE` | `Remove the selected category data`                    |
| `locations`             | `GET`    | `Get a list of locations data`                         |
| `partners`              | `GET`    | `Get a list of partners data`                          |
| `whislist`              | `GET`    | `Get a list of wishlist data`                          |
| `whislist/:id`          | `GET`    | `Select the wishlist data according to registered id`  |
| `whislist`              | `POST`   | `Create new wishlist data`                             |
| `whislist/:id`          | `DELETE` | `Remove the selected wishlist data`                    |

## Contributing

Feel free to contribute the repo for better code!

## Authors

- Me

## Feedback

If you have any feedback, please reach out to us at ilyasalqordhowi@gmail.com
