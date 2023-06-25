# Bookstore Management API

This is a RESTful API written in Go for managing a bookstore. The API uses the Gorilla mux router for handling HTTP requests, the gorm library for object-relational mapping with a MySQL database, and a custom utility package for common functions like parsing JSON from HTTP requests.

## Getting Started

To run this API locally, you'll need to have Go and MySQL installed on your machine.

### Installation

1. Clone the repository:
    ```
    git clone https://github.com/alptalhayazar/Bookstore-Management-API.git
    ```

2. Move into the project directory:
    ```
    cd Bookstore-Management-API
    ```

3. Build and run the API:
    ```
    go build
    ./Bookstore-Management-API
    ```

The API should now be running on `localhost:9010`.

## Usage

This API exposes the following endpoints:

- `GET /book/` - Retrieve all books
- `GET /book/{bookId}` - Retrieve a single book by ID
- `POST /book/` - Create a new book
- `PUT /book/{bookId}` - Update an existing book by ID
- `DELETE /book/{bookId}` - Delete an existing book by ID

### Example

Create a new book:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"My Book", "author":"John Doe", "publication":"Publishing House"}' http://localhost:9010/book/
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Please see the [LICENSE](./LICENSE) file for details.
