# product-service

This is an example of simple REST API implementation with clean architecture written in Go with Dependency Injection along with Mocking (generated using "github.com/vektra/mockery") example 

Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has 4 Domain layer :

- Domain Layer
- Repository Layer
- Service Layer
- Controller Layer

#### The diagram:

![](clean-arch.png)
The pattern itself is designed to create decoupled systems where the implementation of lower-level domains is not a concern for the implementer. This allows for the implementation to be replaced without worrying about breaking the functionality of the implementer.

To ensure this decoupling, every implementation should be done using interfaces. Direct access from the implementer to the implementation should be avoided. This allows for the injection of dependencies and the replacement of implementation with mock objects during unit tests.

For example, the Interfaces folder should contain all the structs under the interfaces namespace. These interfaces act as bridges between different domains, allowing them to interact with each other. In our case, this should be the only way for them to interact.

In Go, interfaces are implemented implicitly, which differs from other languages like Java or C# where interfaces are implemented explicitly. To implement an interface in Go, all the methods that the interface has must be implemented. Once this is done, you're good to "Go".

### Dependency injection:
Dependency injection is a design pattern that allows for the separation of concerns and the decoupling of components in software development.

When a component depends on another component, it creates a tight coupling between them. This makes the code harder to maintain, test, and extend. Dependency injection solves this problem by decoupling the components, allowing for them to be developed, tested, and maintained independently.

In dependency injection, instead of a component creating its own dependencies, the dependencies are injected into the component from an external source. This external source could be a framework, a configuration file, or another component. This allows for the component to be easily replaced with a different implementation of the dependency, or a mock object during testing.

By using dependency injection, you can improve the modularity, testability, and maintainability of your code. It also promotes the separation of concerns, making your code easier to understand and reason about.

## Framework

- Web : Echo
- Configuration : Env File
- Database : Postgres

## Architecture
Handler -> Service -> Repository

## Database Diagram
![](product-service-database-diagram.jpeg?raw=true)


### How To Run This Project

> Make Sure you have run the init.sql in your postgres if not use docker for database
> 
> Fill .env file with your postgres configuration (change POSTGRES_HOST if not using docker as database)

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone git@github.com:mhaqiw/product-service.git

#move to project
$ cd product-service

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps


# Stop
$ make stop
```

# Request example

### Add Product

```bash
curl --location 'http://localhost:9090/products' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Xiaome Note 12",
    "price": 20.00,
    "description": "Test"
}'
```

### Get All Product
You can use this option for sorting:
- newest: Newest product
- cheapest: Product with the lowest price
- expensive: Product with the highest price
- az: Sort by product name (A-Z)
- az: Sort by product name (Z-A)
```bash
# without sorting
curl --location 'http://localhost:9090/products'
```

```bash
# with sorting
curl --location 'http://localhost:9090/products?sorting=az'
```


