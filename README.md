# gecho-starter
gecho-starter is a boilerplate for writing services in Golang quicker out the gates.

## Foreword
This is my **first attempt** to write non-trivial code in this architecture style in Golang, so feedback is both greatly welcomed and desired.

The architecture used for this boilerplate is more or less derived from [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and similar patterns such as Onion Architecture, and Hexagonal Architecture but is not strictly following any one pattern.

It attempts to strike a balance between being *just wet enough* to bring seperation but *DRY enough* to not make you catch a cold.

## Architecture & Structure
In keeping with Clean Architecture, there are four layers

1. External Interfaces (/http via Echo, /repository)
2. Controllers (/http/controller)
3. Use Cases (/interactor)
4. Entities (/entity)

If you get Clean Architecture, this pattern should be pretty similar and easy to follow. Things like like inward dependencies only (layers can reach can reach in but never out) all hold true.

Various notes on implementation:
- Entities, or the domain, are just structs. They describe our data in it's pure form and how we interact with it in Golang.
  - Interactors should be using entities, keeping it interface agnostic.
  - Repositories should be using entities too for actual data. Control entities needed such as limit or cursot are just implemented as parameters on the function.
- Interactors are responsible for business logic.
- Controllers are thin. In our case they just convert data to/from the interface.
- Echo is our "http" interface in this example, which has a `main.go` file for running/starting that interface.
- Error handling is centralized at the interface level so errors can be customized to the specific interface. This means everything should return both value and error and bubble up on error.
- I split up the controllers and interactors into their resource package, then by the specific interactor/controller. I dislike really long files and you're free to not do this .

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