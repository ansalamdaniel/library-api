# library-api

Assignement for Go.

## Task

Write a simple library application that will have two components API and database. Your application should be able to accept a book detail and store that in database and similarly it should be able to list all the books that are there in the database. API is going to accept HTTP GET and POST method to get and store data successfully.
It will be nice to write unit test for above code.

<https://github.com/infracloudio/citadel-internal/blob/master/modules/go/README.md#task>

## Project Setup & Running

Setup installed packages.

```Shell
go mod download
```

Run the server locally

```Shell
go run cmd/web/!(*_test).go
```

Run test cases

```Shell
go test -v ./cmd/web/
```

Build this project

```Shell
go build ./cmd/web/
```

## Testing

Test cases written for this project tests the handlers funactionality. Techniques used ot test the handlers is by creating a temporary DB instance.

### TODO

* Write test cases using data mock pattern.
* Create API documentaion.

## References

Articles in the below blog helped to create a this project.

* <https://www.alexedwards.net/blog>

* <https://gorm.io/docs/>

* Stackoverflow
