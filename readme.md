# Coffee Shop Golang

A project coffe shop, build with using Golang, Gingonic and PostgreSQL

---

## Requirements

For development, you will only Golang, PostgreSQL and global package like Gingonic, Json Web Token, argon2, dotenv installed in your environement.

### Golang

- #### Golang installation on Windows

  Just go on [official Golang website](https://go.dev/doc/install) and download the installer.

If the installation was successful, you should be able to run the following command.

    go version

###

## Install

    git clone https://github.com/ridwanbahtiar15/coffe-shop
    cd PROJECT_TITLE
    go mod init

## Configure app

Create file `.env` then edit it with your settings. You will need:

- DB_HOST = "localhost"
- DB_NAME = "coffeeshop"
- DB_USER = "Your PostgreSQL Usrename"
- DB_PASS = "Your PostgreSQL Password"

- JWT_KEY = "SECRET"
- ISSUER = "Your Issuer, Up to you"

## Running the project

    $ go run ./cmd/main.go
