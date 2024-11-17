# GO GIN POSTGRES API Template [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/gunavault/Go-Gin-Postgres-template/blob/main/LICENSE)

This is a GO GIN POSTGRES API Template built with Go and the Gin framework. It allows user registration, login, and access to user information with JWT-based authentication. you can custom this API as you need. _I'll update this template little by little when I have a chance_.


## Features

- User registration with password hashing
- User login with JWT token generation
- JWT-based authentication middleware
- Testing API with get User

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/my-portfolio-api.git
   cd my-portfolio-api
2. Install Dependency:
   ```bash
   go mod tidy
3. Create a PostgreSQL database and Ensure the required uuid-ossp extension is enabled for UUID generation.
4. run the query from `query.sql`
5. run the code :
   ```bash
   go run main.go

## Endpoints
- POST`/login`
- POST`/register`
- GET`/users`
- GET`/user/username`
  
You can find more the endpoint at thunderclient collection `thunder-collection_Go Gin postgres API template.json`

## License
[MIT licensed](./LICENSE).

