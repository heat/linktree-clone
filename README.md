
## Pre requisites
- GO 1.12 or above https://golang.org/dl/
- Mongodb service

## Installation

clone repository
`git clone https://github.com/heat/linktree-clone.git`

### Linux or OSX

build application
`go build -o main cmd/linktree/main.go` 

running in 8080 port and a mongo simple database
`./main -http.addr=":8080" -database.url "mongodb://..."`

# Endpoints

## POST [localhost:8080/links]

data :
```json
{
	"name": "Onezino Moreira",
	"slug": "onezino",
	"email": "zinogabriel@gmail.com",
	"links": [{
		"kind": "instagram",
		"name": "My public instagram Profile",
		"url": "https://www.instagram.com/ladygaga/?hl=pt-br"
	}]
}
```


## PUT [localhost:8080/links/{slug}]

Replace the whole document. data:
```json
{
	"name": "Onezino Gabriel",
	"slug": "onezino",
	"email": "zinogabriel@gmail.com",
	"links": [{
		"kind": "instagram",
		"name": "My public instagram Profile",
		"url": "https://www.instagram.com/ladygaga/?hl=pt-br"
	}]
}
```

## GET [localhost:8080/links/{slug}]

Retrive a user information

## DELETE [localhost:8080/links/{slug}]

Delete a user information
