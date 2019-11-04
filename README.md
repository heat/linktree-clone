
## POST [localhost:8080/links]

data :
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


## PUT [localhost:8080/links/{slug}]

replaces the whole document. data:
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
