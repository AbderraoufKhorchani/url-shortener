# URL Shortener

This Golang-based URL shortener service allows you to create shortened URLs and easily redirect users to the original links. It uses PostgreSQL as the backend database to store URL mappings. The service comes with Swagger annotations for a well-documented API, facilitating seamless integration into various applications.

## Features

-   URL shortening with unique shortcodes
-   Secure and reliable storage of original and shortened URLs
-   Well-documented API with Swagger annotations
-   Easy redirection to original URLs
## Getting Started

### Prerequisites

To run this authentication service, ensure you have the following prerequisites installed:

- Go programming language
- PostgreSQL database 

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/rf-krcn/url-shortener.git
   cd url-shortener
```
2. Install dependencies:
```go
   go install ./...
```
3. Configuration

Configure the service by setting up the PostgreSQL connection details by editing the DSN in the main.go file.




## API Documentation<a id="api-doc"></a>

The API is thoroughly documented using Swagger annotations. Access the Swagger UI at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) to explore and interact with the API.
