# Go Products API

A simple and efficient RESTful API developed in GoLang 1.21 with PostgreSQL integration. The API provides endpoints for managing 'products' and 'user' entities, supporting CRUD operations for products and user creation, along with token generation using JWT authentication.

## Features

- **Products Endpoint:**
  - CRUD operations for managing products.
  - Authentication using JWT for secure access.

- **User Endpoint:**
  - User creation.
  - Token generation for authentication.

## Technologies Used

- GoLang 1.21
- PostgreSQL 13+
- JWT Authentication
- Swagger for API documentation

## Installation

1. Ensure you have GoLang 1.21 installed.
2. Set up PostgreSQL and update the database configuration in `go-products-api/cmd/server/.env`.
3. Run the following commands:

    ```bash
    go mod tidy
    ```

4. Run the application:

    ```bash
    go run cmd/server/main.go
    ```
5. By default the server will be available at [http://localhost:8080](http://localhost:8080)

## API Endpoints

- **Products Endpoint:**
  - `GET /products`: Retrieve all products.
  - `GET /products/{id}`: Retrieve a specific product by ID.
  - `POST /products`: Create a new product.
  - `PUT /products/{id}`: Update a product by ID.
  - `DELETE /products/{id}`: Delete a product by ID.

- **User Endpoint:**
  - `POST /users`: Create a new user.
  - `POST /generate_token`: Generate JWT token for authentication.

## JWT Authentication

To access protected endpoints, include the JWT token in the `Authorization` header of your requests.

## Documentation

Swagger documentation is available at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html).
