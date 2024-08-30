# Bookstore RESTful API

## Project Overview

This project is a RESTful API for managing a bookstore, allowing users to perform various operations related to books in the store. The API is built using the Gin framework for Go, and it interacts with a MySQL database to store and retrieve book data. The primary features of the API include fetching all books, adding a new book, fetching a book by its ID, and running tests to ensure the API endpoints function as expected.

## Features

- **Fetch All Books**: Retrieve a list of all books available in the bookstore.
- **Add a New Book**: Add a new book to the bookstore with details such as title, author, and price.
- **Fetch a Book by ID**: Retrieve detailed information about a specific book using its unique ID.
- **API Endpoint Tests**: Unit tests are provided to verify the functionality of the API endpoints.

## API Endpoints

### 1. Fetch All Books
- **Endpoint**: `/books`
- **Method**: `GET`
- **Description**: Returns a list of all books available in the bookstore.
- **Response**:
  ```json
  [
      {
          "id": 1,
          "title": "The Great Adventure",
          "author": "John Smith",
          "price": 99.99
      },
      {
          "id": 2,
          "title": "Mystery of the Lost Temple",
          "author": "Jane Doe",
          "price": 79.99
      }
      ...
  ]
  
### 2. Add a New Book
- **Endpoint**: `/books`
- **Method**: `POST`
- **Description**: Adds a new book to the bookstore.
- **Request Body**:
  ```json
  {
      "title": "The New Book",
      "author": "New Author",
      "price": 49.99
  }
### 3. Fetch a Book by ID
- **Endpoint**: `/books/:id`
- **Method**: `GET`
- **Description**: Fetches detailed information about a specific book using its unique ID.
- **Response**:
  ```json
  {
      "id": 1,
      "title": "The Great Adventure",
      "author": "John Smith",
      "price": 99.99
  }

## Testing the API

The project includes a suite of tests to ensure the API endpoints function correctly. These tests cover scenarios such as fetching all books, adding a new book, and fetching a book by its ID.

### Running the Tests

To run the tests, execute the following command in your terminal:

```bash
go test ./handlers -v
```
## Project Setup

### Prerequisites

- **Go**: Ensure you have Go installed on your system.
- **MySQL**: The API uses a MySQL database to store book data. Ensure MySQL is installed and running.

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/bookstore-api.git
   cd bookstore-api
    ```
2. **Install dependencies**:
   ```bash
   go mod tidy
   
3. Set up the database:

Create a MySQL database named bookstore.
Import the database schema provided in the schema.sql file to set up the books table.

4.Configure the database connection:

Update the database/database.go file with your MySQL connection details (username, password, database name).

5. Run the application:
  ```bash
     go build
