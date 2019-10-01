# Simple CRUD API in Golang

This project implements a simple CRUD API in Golang that is capable of saying hello in multiple languages, and even learning (and forgetting) new ones !

## Dependencies

This project is using ```gorilla/mux```.
```bash
# Install gorilla/mux
go get -u github.com/gorilla/mux
```

## Installation & Run
```bash
# Download this project
go get github.com/AlexandreLadriere/Go-Simple_CRUD_API
```

```bash
# Build
cd go/src/github.com/AlexandreLadriere/Go-Simple_CRUD_API
go build
```

```bash
# Run
./Go-Simple_CRUD_API
# API Endpoint : http://127.0.0.1:10001
```

## Structure

```
├── pkg
│    ├── config
│    ├── controllers 
│    │    └── hello-controllers.go
│    ├── models 
│    │    ├── hello.go
│    │    ├── hello_test.go
│    │    ├── hello_cover.txt
│    │    └── hello_cover.html
│    ├── routes
│    │    └── hello-routes.go
│    └── utils 
│         └── utils.go
└── main.go
```

## API

#### /
* `GET` : Get default Hello

#### /hellos
* `GET` : Get all ```Hello``` objects

#### /hello
With a ```lang=value``` query (e.g. ```/hello?lang=fr```)
* `GET` : Get the ```Hello``` in the specified ```value``` language (e.g. http://127.0.0.1:10001/hello?lang=de) or an empty string if it does not exist

* `DELETE` : Delete the ```Hello``` in the specified ```value``` language.

With a json body (e.g. ```/hello {"language": "de", "hello": "Hallo"}```)

* `POST` : Create a new ```Hello``` with the corresponding language and the corresponding message if it does not already exist

#### 

## Developed with
  - go version go1.13.1 linux/amd64