# Goa Application README

Welcome to the README for your Goa application! This guide assumes you have successfully set up your development environment by following these steps:

1. Install Goa:
    ```
    go install goa.design/goa/v3/cmd/goa@v3
    ```

2. Download the protoc binary from releases and ensure it's in your path.

3. Install the protoc plugin for Go:
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

## Getting Started

### Running the Application

To run your Goa application, follow these steps:

1. Clone this repository:

    ```
    git clone https://github.com/redigongo/Book.git
    cd bookapp
    ```

2. Install dependencies:

    ```
    go mod tidy
    ```

3. Generate Goa files:

    ```
    goa gen bookapp/design
    ```

4. Run the application:

    ```
    go run bookapp/cmd/books
    ```

Your Goa application should now be running locally on `http://localhost:8080`

## Postman Collection

Explore our Postman Collection to easily test your Goa application's API endpoints: [Postman Collection](https://www.postman.com/gold-star-692236/workspace/bookapi/collection/20581054-fe817ed7-15d1-4616-bc55-c6bb0d484d39?action=share&creator=20581054)


# Database Package

The `database` package in this project manages the connection to the MySQL database using the GORM library.

## Configuration

Before using the `database` package, configure the database connection parameters in the `database.go` file:

```go
// Database configurations
// Change based on your MySQL configurations
const DB_USERNAME = "root"
const DB_PASSWORD = "1234"
const DB_NAME = "bookapp"
const DB_HOST = "localhost"
const DB_PORT = "3306"
```
PS: Don't forget to create a schema on MySQL, called `bookapp`.


# Books Service API Design

**API Declaration:**
- The API is named "books" and serves as an HTTP service for managing books.
- It is hosted on http://localhost:8080.

**Service Declaration:**
- The "books" service includes methods for listing, showing, creating, updating, and deleting books.

**Data Types:**
- File: Represents the structure for handling image uploads.
- BookDTO: Defines the payload structure for book data, including ID, title, author, book cover details, and the published date.

**HTTP Endpoints:**
- Various HTTP methods (GET, POST, PUT, DELETE) are defined for CRUD operations on books.
- Multipart requests are specified for endpoints that involve file uploads.

**Response Types:**
- The API uses a custom media type (application/vnd.book+json) for representing book data.



# **BooksCreateDecoderFunc**

- **Purpose:** This function is designed to decode multipart requests specifically for the "create" endpoint of the Books Service API.

  - **Decoding Logic:**
      1. **Iterating through Parts:** Utilizing the `multipart.Reader`, the function iterates through each part of the multipart request.

    ```go
    for {
        part, err := mr.NextPart()
        if err == io.EOF {
            break
        }
        // ... (rest of the loop)
    }
    ```

    2. **Content-Disposition Parsing: The Content-Disposition header is parsed to extract parameters such as name, filename, etc.
    ```
    _, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
    ```

    3. **Handling "book_cover" Part: For the part named "book_cover," it checks if it's an image (based on Content-Type) and reads its bytes.
    ```
    if params["name"] == "book_cover" {
      if err != nil || strings.HasPrefix(disposition, "image/") {
        bytes, err := io.ReadAll(part)
        // ... (rest of the logic)
        }
      }
    ```
    
    4. **Populating Payload: The function populates a books.CreatePayload object with relevant data extracted from different parts.

    ```
      res := books.CreatePayload{
          // ... (populating other fields)
          BookCover: &imageUpload,
      }
    ```
- **Common Logic with BooksUpdateDecoderFunc:**
  - The same logic is applied in the `BooksUpdateDecoderFunc` for decoding multipart requests in the "update" endpoint, where similar steps are followed for handling different parts and populating the respective payload.


