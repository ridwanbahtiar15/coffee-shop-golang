# Coffe Shop With Golang

<!-- ABOUT THE PROJECT -->

## About The Project

A web api project for ordering coffee and transactions online. There are 4 operations that can be performed, Get (fetching data), Post (insert data), Update (update partial data), delete (delete data)

## Built With

- [![Golang][Golang-logo]][Golang-url]
- [![PostgreSql][PostgreSql-logo]][PostgreSql-url]

### Package

- [![GinGonic][GinGonic-logo]][GinGonic-url]
- [![GoValidator][GoValidator-logo]][GoValidator-url]
- [![Cloudinary][Cloudinary-logo]][Cloudinary-url]
- [![Pq][Pq-logo]][Pq-url]
- [![DotEnv][DotEnv-logo]][DotEnv-url]

## Install And Run Locally

Clone project from github repository

    $ git clone https://github.com/ridwanbahtiar15/coffee-shop-golang.git

go to folder coffee-shop

    $ cd coffee-shop-golang

install dependencies

    $ go get -d ./...

Start the server

    $ go run ./cmd/main.go

## Configure app

Create file `.env` then edit it with your settings
according to your needs. You will need:

| Key               | Value                  |
| ----------------- | ---------------------- |
| DB_HOST           | Your Database Host     |
| DB_NAME           | Your Database Host     |
| DB_USER           | Your Database User     |
| DB_PASSWORD       | Your Database Password |
| DB_SSLMODE        | disable                |
| JWT_SECRET        | Your JWT Secret        |
| JWT_ISSUER        | Your Issuer            |
| CLOUDINERY_NAME   | Your Cloudinary Name   |
| CLOUDINERY_KEY    | Your Cloudinary Key    |
| CLOUDINERY_SECRET | Your Cloudinary Secret |

## Api Refrences

Auth
| Route | Method | Description |
| -------------- | ----------------------- | ------ |
| /auth/login | POST | Login user |
| /auth/register | POST | Register user |
| /auth/logout | DELETE | Logout user |

Product
| Route | Method | Method |
| -------------- | ----------------------- | ------ |
| /products | GET | Get all product |
| /products/:name | GET | Get product by name |
| /products/:category | GET | Get product by category |
| /products | POST | Create new product |
| /products/:id | PATCH | Update Product |
| /products/:id | DELETE | Delete Product |

Order
| Route | Method | Description |
| -------------- | ----------------------- | ------ |
| /orders | GET | Get all order |
| /orders | POST | Create new order |
| /orders/:id | GET | Get Order by id |
| /orders/:id | PATCH | Update status order |

User
| Route | Method | Description |
| -------------- | ----------------------- | ------ |
| /users | GET | Get all user |
| /users | POST | Create new user |
| /users/:id | GET | Get user by id |
| /users/:id | PATCH | Update user by id |
| /users/:id | DELETE | Delete user by id |
| /users/profile | GET | Get user profile |
| /users/profile/edit | PATCH | Update user profile |

## Documentation

[Postman Documentation](https://documenter.getpostman.com/view/28541505/2s9YsT4npX)

## Related Project

[Front End With React JS](https://github.com/ridwanbahtiar15/coffee-shop-react-vite.git)
[Back End With Express JS](https://github.com/ridwanbahtiar15/coffe-shop.git)

## Credit

[Ridwan Bahtiar](https://github.com/ridwanbahtiar15)

<!-- MARKDOWN LINKS & IMAGES -->

[Golang-url]: https://go.dev/
[Golang-logo]: https://img.shields.io/badge/Golang-blue
[Gingonic-url]: https://gin-gonic.com/
[Gingonic-logo]: https://img.shields.io/badge/Gin%20Gonic-lightskyblue
[PostgreSql-url]: https://www.postgresql.org/
[PostgreSql-logo]: https://img.shields.io/badge/Postgre%20SQL-blue
[GoValidator-url]: https://github.com/asaskevich/govalidator
[GoValidator-logo]: https://img.shields.io/badge/Go%20Validator-red
[Cloudinary-url]: https://github.com/cloudinary/cloudinary-go
[Cloudinary-logo]: https://img.shields.io/badge/Cloudinay-green
[Pq-url]: https://github.com/lib/pq
[Pq-logo]: https://img.shields.io/badge/pq-grey
[DotEnv-url]: https://github.com/joho/godotenv
[DotEnv-logo]: https://img.shields.io/badge/godotenv-black
