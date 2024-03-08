# Registration API


### General Info
- * This project provides an API for user registration, login, and modification of user data.

## Features

- **Registration**: Users can register by providing their email and password.
- **Login**: Registered users can log in using their email and password.
- **Modify User Data**: Authenticated users can modify their profile information, such as name, email, mobile number, and date of birth.

## API Endpoints

### Registration
- **URL**: /register
- **Method**: POST
- **Request Body**:
  {
      "email": "user@example.com",
      "password": "strongpassword",
      "repeated_password": "strongpassword"
  }
- **Response**: Returns user details upon successful registration.

### Login
- **URL**: /login
- **Method**: POST
- **Request Body**:
  {
      "email": "user@example.com",
      "password": "strongpassword"
  }
- **Response**: Returns a session cookie upon successful login.

### Modify User Data
- **URL**: /modify
- **Method**: POST
- **Request Body**:
  {
      "name": "John Doe",
      "email": "user@example.com",
      "mobile_number": "+1234567890",
      "date_of_birth": "1990-01-01",
      "password": "newpassword"
  }
- **Headers**: Include session cookie in the Authorization header for authentication.
- **Response**: Returns updated user details upon successful modification.

## Setup

1. Adjust the environment variables in the .env file:

   DATABASE_URL=user=Yournickname password=Yourpassword dbname=mobydev sslmode=disable
   MIGRATION_PATH=./migrations/init.sql

2. Run the server:

   go run ./cmd/

