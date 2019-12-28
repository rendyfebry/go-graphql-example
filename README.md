# GraphQL with Golang Example Project

[![Build Status](https://github.com/rendyfebry/go-graphql-example/workflows/main-action/badge.svg)](https://github.com/rendyfebry/go-graphql-example/actions)

## Info

GraphQL implementation with Golang

## Prerequisites

- Go 1.12
- Dep

## Install Depedencies

```
dep ensure -v
```

## Run

```
go run main.go
```

```
curl --location --request POST 'http://localhost:8080/graphql' \
  --header 'Content-Type: application/graphql' \
  --data-raw '{"query":"query{user(id: 1){id,name,age}}"}'
```

# Test

```
go test -v ./...
```
