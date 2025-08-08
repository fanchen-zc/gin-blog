# Gin-Test: A Gin-based RESTful API

This project is a RESTful API built with the Gin framework for Go. It serves as a backend for a blog application, with features for managing articles and tags. It also includes a UI analysis endpoint that can analyze a website's UI from a screenshot.

## Features

- **JWT Authentication**: Secure your endpoints with JSON Web Tokens.
- **Tag Management**: CRUD operations for tags.
- **Article Management**: CRUD operations for articles.
- **Image Upload**: Upload images with size and extension validation.
- **UI Analysis**: Analyze a website's UI from a screenshot and get feedback (mocked response).
- **Swagger Documentation**: API documentation with Swagger.
- **Configuration Management**: Easy configuration with an `.ini` file.
- **Logging**: Log application events to files.
- **Redis Cache**: Caching layer with Redis.

## Getting Started

### Prerequisites

- Go 1.17 or higher
- MySQL
- Redis

### Installation

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd gin-test
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Configure the application:**
    - Rename `conf/app.ini.example` to `conf/app.ini` (if an example file were provided).
    - Update `conf/app.ini` with your database and Redis connection details.

    ```ini
    [database]
    Type = mysql
    User = your-db-user
    Password = your-db-password
    Host = 127.0.0.1:3306
    Name = blog
    TablePrefix = blog_

    [redis]
    Host = 127.0.0.1:6379
    Password = your-redis-password
    ```

4.  **Run the application:**
    ```bash
    go run main.go
    ```
    The application will be running at `http://127.0.0.1:8000`.

## API Endpoints

### Authentication

-   `GET /auth`: Get an authentication token.

### Tags

-   `GET /api/v1/tags`: Get a list of tags.
-   `POST /api/v1/tags`: Create a new tag.
-   `PUT /api/v1/tags/:id`: Update an existing tag.
-   `DELETE /api/v1/tags/:id`: Delete a tag.
-   `POST /tags/export`: Export tags to an Excel file.
-   `POST /tags/import`: Import tags from an Excel file.

### Articles

-   `GET /api/v1/articles`: Get a list of articles.
-   `GET /api/v1/articles/:id`: Get a specific article.
-   `POST /api/v1/articles`: Create a new article.
-   `PUT /api/v1/articles/:id`: Update an existing article.
-   `DELETE /api/v1/articles/:id`: Delete an article.
-   `POST /api/v1/articles/poster/generate`: Generate a poster for an article.

### Image Upload

-   `POST /upload`: Upload an image.

### UI Analysis

-   `POST /api/v1/analyse`: Upload a website screenshot for UI analysis.

### Swagger Documentation

-   `GET /swagger/index.html`: View the Swagger API documentation.
