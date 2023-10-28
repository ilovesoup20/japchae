# Init project

## Init
```bash

$ go mod init japchae_api

$ go get github.com/gofiber/fiber/v2

```

## Using GORM

https://gorm.io/index.html

```
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/sqlite
```

## Go Fiber hot reloading

https://stackoverflow.com/questions/71643902/how-to-reload-go-fiber-in-the-terminal

```
npm install -g nodemon
nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

```

## Using Fiber

- https://docs.gofiber.io/
  - go get github.com/gofiber/fiber/v2

## Docker

https://hub.docker.com/_/golang

https://youtu.be/31ieHmcTUOk

- Add Dockerfile with basic instructions

## Ent (ORM)

https://github.com/ent/ent
https://entgo.io/docs/getting-started/

## Project Structure Standard

https://softchris.github.io/golang-book/03-projects/01-first-project/

https://www.wolfe.id.au/2020/03/10/starting-a-go-project/

https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/

https://github.com/golang-standards/project-layout