# Microservice Server Cake

This is a microservice-based server application for managing cakes and users, built with Go.

## Features

- **User Management:**
    - Register new users
    - User login
    - Get user information by ID
- **Cake Management:**
    - Create new cakes
    - Get cake information by ID
    - Search for cakes
    - Update cake information

## Technologies Used

- **Language:** Go
- **Frameworks/Libraries:**
    - [Twirp](https://github.com/twitchtv/twirp): A framework for simple RPC from Twitch.
    - [Protocol Buffers](https://developers.google.com/protocol-buffers): A language-neutral, platform-neutral, extensible mechanism for serializing structured data.
    - [GORM](https://gorm.io/): The fantastic ORM library for Go.
    - [PostgreSQL](https://www.postgresql.org/): A powerful, open-source object-relational database system.
    - [RabbitMQ](https://www.rabbitmq.com/): A popular open-source message broker.
    - [Wire](https://github.com/google/wire): A code generation tool for dependency injection.
    - [Docker](https://www.docker.com/): A platform for developing, shipping, and running applications in containers.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.24.3 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [make](https://www.gnu.org/software/make/)

### Installation & Running

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/Anhbman/microservice-server-cake.git
    cd microservice-server-cake
    ```

2.  **Create a `.env` file:**

    Create a `.env` file in the root of the project and add the following environment variables:

    ```
    HTTP_PORT=8081
    DB_USER=developer
    DB_PASSWORD=developer
    DB_NAME=cake_dev
    DB_HOST=localhost
    DB_PORT=5432
    ```

3.  **Run the application with Docker Compose:**

    This will start the Go application and a PostgreSQL database.

    ```bash
    docker-compose up -d
    ```

    The server will be running at `http://localhost:8081`.

4.  **Run the application locally (without Docker):**

    Make sure you have a running PostgreSQL instance.

    ```bash
    dotenv -e .env -- go run ./cmd/server
    ```

## Project Structure

```
.
├── cmd/server           # Main application entry point
├── internal             # Internal application logic
│   ├── config           # Configuration management
│   ├── controller       # Controllers (request handlers)
│   ├── service          # Business logic
│   └── storage          # Database interactions
├── pkg/rabbitmq         # Reusable RabbitMQ package
├── rpc/service          # RPC service definitions (.proto files)
└── docker               # Docker-related files
```

## Dependency Injection

This project uses [Wire](https://github.com/google/wire) for dependency injection. To regenerate the `wire_gen.go` file, run the following command:

```bash
go run github.com/google/wire/cmd/wire
```