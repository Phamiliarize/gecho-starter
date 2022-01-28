# gecho-starter
gecho-starter is a boilerplate for writing services in Golang quicker out the gates.

## Foreword
This is my **first attempt** to write non-trivial code in this architecture style in Golang, so feedback is both greatly welcomed and desired.

The architecture used for this boilerplate is more or less derived from [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and similar patterns such as Onion Architecture, and Hexagonal Architecture but is not strictly following any one pattern.

It attempts to strike a balance between being *just wet enough* to bring seperation but *DRY enough* to not make you catch a cold.

## Architecture & Structure
It mostly adheres Clean Architecture:

1. External Interfaces (/http via Echo, /repository)
2. Controllers (/http/handler)
3. Use Cases (/services)
4. Entities (/entity)

If you get Clean Architecture, this pattern should be pretty similar and easy to follow. Things like like inward dependencies only (layers can reach can reach in but never out) all hold true.

## Running It
If you would like to  run it make sure to populate a `.env` file with your database settings, reference `.template_env` for the naming.

To run the HTTP interface (Echo):
```shell
go run http/main.go
```

To run the unit tests with coverage:
```shell
go test -v ./... -race -covermode=atomic -coverprofile=coverage.out
```

You can generate a coverage report in HTML with:
```shell
go tool cover -html=coverage.out -o=coverage.html
```

## Special Thanks
Special thanks to the "dev" community (ya'll know who you are) for helping me understand many of these concepts and providing feedback.