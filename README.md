# Project Simple API

## Overview

This document provides comprehensive instructions on setting up and configuring the Superindo API project. It includes steps for both Docker container deployment and local machine setup using a `.env` file.

## Prerequisites

Before you start, make sure you have the following:

- [ ] Docker installed on your machine.

## Instructions

### Running in Docker Container

To run the application using Docker, follow these steps:
1. Run Docker Compose:
    Execute the following command to launch the application along with its dependencies:
    ```cmd
    docker-compose up
    ```
    Note: Wait until the MySQL container, Redis, and the API application are running correctly.
    
2. Base URL for API app:
    Once the containers are up and running, access the API using the following base URL:
    ```cmd
    http://localhost:5580
    ```

### Running on Local Machine
To set up the application on your local machine, follow these steps:
1. **Copy the `.env.example` file:**

   If you don't have an `.env` file, start by copying the provided `.env.example` file.
   
   ```
   cp .env.example .env
   ```

2. Open the .env file:

    Use your preferred text editor to open the .env file.

    ```
    nano .env
    ```
3. Update the Database Configuration:

    Locate the database configuration section in the .env file. Update it to match your provided database information:
    
    ```
    DB_HOST=your_database_host
    DB_PORT=your_database_port
    DB_DATABASE=your_database_name
    DB_USERNAME=your_database_username
    DB_PASSWORD=your_database_password
    DB_DEBUG=false
    ```
    Replace the placeholder values (`your_database_name`, `your_database_user`, `your_database_password`) with your actual database information. Set `DB_DEBUG` to `true` if you want to enable query logging to the console.

4. Logger Configuration:

    Add the following lines for logger configuration:

    ```
    LOGGER_LOGS_WRITE=true
    LOGGER_FOLDER_PATH=./logs
    ```
    Adjust the LOGGER_FOLDER_PATH based on your preferred folder structure.

5. Ensure that the application is configured with the following environment variable:
    ```
    APP_PORT=5580
    ```

6. Save and Close the File:

    Save the changes and close the .env file.

7. Verify the Configuration:

    Make sure your application can connect to the database using the updated configuration. You can do this by running a database-related task or checking your application logs.

8. Run Unit Testing:

    Execute the following command to run unit tests and generate a coverage report:

    ```
    make test-coverage
    ```

## Additional

- Import Postman Collection:
    Use Postman to export the provided collection file [Superindo Postman](https://github.com/fadilahonespot/superindo/blob/master/resources/Superindo.postman_collection.json) to your postman.


# API Documentation

This documentation outlines the usage of the Simple API, providing details on various endpoints and their functionalities.

## Endpoints

### 1. Get Categories
- **Method:** GET
- **Endpoint:** `http://localhost:5580/category`
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": [
            {
                "id": 1,
                "name": "Sayuran"
            },
            {
                "id": 2,
                "name": "Protein"
            },
            {
                "id": 3,
                "name": "Buah dan Snack"
            }
        ]
    }
    ```

### 2. Add Product

- **Method:** POST
- **Endpoint:** `http://localhost:5580/product`
- **Request Body:**
    ```json
    {
        "name": "indomie rasa ayam bawang",
        "description": "mie dengan rasa yang nendang ayam bawang goreng",
        "weight": 300,
        "price": 2500,
        "image": "http://google.com/mie-indomie-ayam-bawang.jpg",
        "categoryId": 3
    }
    ```
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": {
            "id": "7aec2d81-0313-42bf-8e0e-b40730e31ddf"
        }
    }
    ```

### 3. Get List Product

- **Method:** GET
- **Endpoint:** `http://localhost:5580/product?page=1&limit=10`
- **Query Parameters:**
    - `search` (string): indomie
    - `categoryId` (int): 3
    - `minPrice` (int): 1000
    - `maxPrice` (int): 3000
    - `createdAt` (date format): 2024-03-18
    - `productName` (string): wortel
    - `page` (int): 1
    - `limit`(int): 10
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": [
            {
                "id": "1b61b02c-78bb-4b41-bfcf-c1c9c3b67158",
                "name": "Indomie Kuah Rasa Soto",
                "description": "Mie instan dengan kuah soto yang khas dan gurih.",
                "weight": 350,
                "price": 2800,
                "image": "http://example.com/mie-indomie-kuah-soto.jpg",
                "categoryId": 3,
                "createdAt": "2024-03-20T10:45:39.826Z"
            },
            {
                "id": "229af19a-1614-4ed4-9b4e-55f6e6d45e51",
                "name": "Roti Whole Wheat",
                "description": "Roti gandum utuh yang kaya serat dan nutrisi.",
                "weight": 350,
                "price": 2500,
                "image": "http://example.com/roti-whole-wheat.jpg",
                "categoryId": 3,
                "createdAt": "2024-03-18T10:44:39.826Z"
            },
        ],
        "pagination": {
            "page": 1,
            "limit": 10,
            "totalData": 2,
            "totalPage": 1
        }
    }
    ```

### 4. Get Product Detail

- **Method:** GET
- **Endpoint:** `http://localhost:5580/product/22c8e385-6d60-4ddb-87b2-3fb543d43177`
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": {
            "name": "indomie rasa ayam bawang cincang",
            "description": "mie dengan rasa yang nendang ayam bawang goreng",
            "weight": 300,
            "price": 2500,
            "categoryId": 2,
            "categoryName": "Protein",
            "image": "http://google.com/mie-indomie-ayam-bawang.jpg",
            "createdAt": "2024-03-20T11:43:39.826Z"
        }
    }
    ```

### 5. Update Product

- **Method:** PUT
- **Endpoint:** `http://localhost:5580/product/b34e8eac-ac43-4163-b9ad-49f15644b4fa`
- **Request Body:**
    ```json
    {
        "name": "indomie rasa ayam bawang cincang",
        "description": "mie dengan rasa yang nendang ayam bawang goreng",
        "weight": 300,
        "price": 2500,
        "image": "http://google.com/mie-indomie-ayam-bawang.jpg",
        "categoryId": 3
    }
    ```
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": null
    }
    ```

### 6. Delete Product

- **Method:** DELETE
- **Endpoint:** `http://localhost:5580/product/22c8e385-6d60-4ddb-87b2-3fb543d43177`
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": null
    }
    ```
