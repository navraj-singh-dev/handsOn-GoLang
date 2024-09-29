# GoLang REST API with SQL Database & Authentication

Welcome to the **GoLang REST API with SQL Database & Authentication** project! This folder showcases a comprehensive REST API built using GoLang, integrated with a SQL database, and with authentication mechanisms. The project is a culmination of all the handsOn i did with golang.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)

## Introduction

This project is a culmination of my deep dive into GoLang, where I explored various aspects of backend development and golang. The API is designed to handle CRUD operations, manage user authentication, with database of SQL.

## Features

- **RESTful API**: Adheres to REST principles.
- **SQL Database Integration**: Utilizes SQLite for data management.
- **Authentication**: Secure user authentication using [JWT tokens](https://jwt.io/).
- **Error Handling**: Error handling for a smooth user experience.
- **Logging**: Detailed logging for easy debugging and monitoring.
- **Modular Codebase**: Clean and modular code structure for maintainability.

## Technologies Used

- **GoLang**: The core programming language used for building the API.
- **Gin**: A high-performance HTTP web framework for Go.
- **SQLite**: A lightweight SQL database for data storage.
- **JWT**: JSON Web Tokens for secure authentication.

## Installation

To get started with this project, follow these steps:

1. **Clone the repository**:
   ```sh
   git clone https://github.com/navraj-singh-dev/handsOn-GoLang.git
   cd handsOn-GoLang
   ```

2. **Install dependencies**:
   ```sh
   go mod tidy
   ```

3. **Run the server**:
   ```sh
   go run main.go
   ```

## Usage

Once the server is up and running, you can interact with the API using tools like `curl` or Postman. Below are some example commands:

- **Get Events**:
  ```sh
  curl -X GET http://localhost:8080/events
  ```

- **Create Event**:
  ```sh
  curl -X POST http://localhost:8080/events -d '{"name":"Event Name","description":"Event Description","location":"Event Location","dateTime":"2024-09-29T15:08:06.383952Z","user_id":1}'
  ```

## API Endpoints

- **GET /events**: Retrieve all events.
- **POST /events**: Create a new event.
- **PUT /events/:id**: Update an existing event.
- **DELETE /events/:id**: Delete an event.
- and many more, see the /routes/routes.go file..

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request. Let's make this project even better together.
