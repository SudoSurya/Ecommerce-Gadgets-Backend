# E-Commerce Application with Go, Gin, and MongoDB
# Backend Setup

The backend of the e-commerce application is built using Go, Gin, and MongoDB. Go is a powerful programming language known for its simplicity and efficiency, making it an excellent choice for building robust web applications. Gin is a lightweight web framework for Go that provides a simple and intuitive API for creating RESTful APIs. MongoDB, a popular NoSQL database, is used for storing and managing the application's data. With MongoDB, you can benefit from its flexibility and scalability, making it a suitable choice for an e-commerce application that may handle a large volume of data. The backend project is organized into different components, such as routes, controllers, and models, following best practices for structuring a Go web application. By leveraging the power of Go, Gin, and MongoDB, the backend provides a reliable and efficient foundation for the e-commerce application.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Postman](https://www.postman.com/downloads/)

## Getting Started

To get started, clone the repository and navigate to the project directory.

```bash 
git clone 
cd backend
go mod download
go run main.go
```

## Project Structure

The project is organized into different components, such as routes, controllers, and models, following best practices for structuring a Go web application.

``` 
├── controllers
│   ├── products.go
│   └── users.go
├── database
│   └── database.go
├── middlewares
│   └── auth.go (JWT authentication)
├── models
│   ├── product.go
│   └── user.go
├── routes
│   ├── products.go
│   └── users.go
├── utils
│   ├── error.go
│   └── response.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```
<style>
    * {
        font-family: 'Cascadia Code';
    }
</style>
## Technologies

- [Go](https://golang.org/)
- [Gin]
- [MongoDB](https://www.mongodb.com/)
- [JWT](https://jwt.io/)
- [Postman](https://www.postman.com/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [MongoDB Compass](https://www.mongodb.com/products/compass)

## API Endpoints
https://documenter.getpostman.com/view/21427214/2s93sc4sWv

The backend provides the following API endpoints for managing the application's data.

### Products 

- `GET /products` - Get all products
- `GET /products/:id` - Get a single product
- `GET /products/page?page?=1&pageSize=2` - Get products with pagination
### Users

- `POST /users/register` - Register a new user
- `POST /users/login` - Login an existing user
- `GET /users/profile` - Get user profile


