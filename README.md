# REST-API-in-Go

A **minimal REST API built in Go**, designed to demonstrate basic CRUD operations using mock data. This project provides a barebones example of how to set up a RESTful service in Go using native libraries.

## Features

1. `GET /books` — Retrieve all books  
2. `GET /books/{id}` — Retrieve a book by ID  
3. `POST /books` — Create a new book  
4. `PUT /books/{id}` — Update an existing book  
5. `DELETE /books/{id}` — Delete a book by ID  

> NOTE: This project contains hardcoded mock data and does not interact with data from either a JSON file or a database

## Instructions

1. Cloning Repo

```bash
git clone git@github.com:anushkaguptaaa/REST-API-in-Go.git
cd REST-API-in-Go
```

Optional Step: Updating the mod file

```bash
go mod init example.com/m
go mod tidy
```

2. Running the project in both Terminal and POSTMAN
```
go build
./m
```
The server will start running on http://localhost:8080

## Testing the API

You can use tools like Postman, curl, or any REST client to interact with the API.

If using Postman, import the provided collection (if available), or manually create requests to test each route.

## Example Payload [POST / PUT]
```json
{
    "isbn": "21800031781",
    "title": "Book Three",
    "author": {
        "firstname": "Diego",
        "lastname": "Xu"
    }
}
```

## Notes

- Intended for educational or quick prototyping purposes.
- TODO: Integrating a database