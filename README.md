# Evermos Rakamin

This repository contains the backend project developed as part of Rakamin Academy activities related to Evermos.

## Technologies Used

- **Backend:** Golang with Fiber
- **Database:** MySQL
- **Authentication:** JWT
- **Integration:** [API Wilayah Indonesia](https://github.com/emsifa/api-wilayah-indonesia)

## Main Features

- User authentication with JWT
- CRUD operations for products and categories
- User management
- Integration with external services, including API Wilayah Indonesia for regional data

## Installation

### Prerequisites
- Golang installed
- PostgreSQL configured

### Steps
1. Clone this repository:
   ```sh
   git clone https://github.com/fauzan264/evermos-rakamin.git
   cd evermos-rakamin
   ```
2. Install backend dependencies:
   ```sh
   go mod tidy
   ```
3. Configure the environment in `.env`:
   ```env
   # APP CONFIG
   SECRET_KEY=
   APP_HOST=localhost
   APP_PORT=

   # DB CONFIG
   # The DB_HOST & DB_PORT can be left empty when using Docker.
   DB_HOST=
   DB_PORT=
   DB_USER=
   DB_PASSWORD=
   DB_NAME=

   # Integration
   API_LOCATION=
   ```
4. Run the backend application:
   ```sh
   go run main.go
   ```

## API Documentation & Database Schema

- [API Specification](https://github.com/fauzan264/evermos-rakamin/tree/master/docs/api-spec)
- [Database Diagram](https://github.com/fauzan264/evermos-rakamin/tree/master/docs)