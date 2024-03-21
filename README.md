# Go Simple Native REST API

This is a simple Go project that interacts with a database to perform CRUD operations on a `products` table.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/benebobaa/simple-native-crud.git
```

Navigate to the project directory:

```bash
cd simple-native-crud
```

Install the required dependencies:

```bash
go mod download
```

## Running the application

You can run the application using the following command:

```bash
go run main.go
```

## Code Overview

The main parts of the project are:

- `main.go`: This is the entry point of our application.
- `database.go`: This file contains the `NewDB` function which is used to establish a connection with the database.
- `service.go`: This file contains the `Service` struct and its methods. The `GetProducts` method retrieves all products from the database.

## Built With

- [Go](https://golang.org/) - The programming language used.
- [pq](https://github.com/lib/pq) - The PostgreSQL driver used.

## Authors

- **benebobaa** - *Initial work* - [benebobaa](https://github.com/benebobaa)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.