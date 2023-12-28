# CarJod (Cari Jodoh) API

## Introduction

Welcome to the Dating App Backend project! This repository contains the backend system for a simple Dating App, similar to Tinder or Bumble. The goal of this project is to evaluate technical skills and the ability to design a backend system that meets the requirements of a dating application, as part of the technical assessment for the company.

## Problem Statement

In this technical test, you are required to design and implement the backend system for the Dating App, considering the following basic functionalities:

1. **Sign up & Login to the App:**
   - Implement user authentication and authorization.
   - Users should be able to sign up with a new account and log in with existing credentials.

2. **User Profile Interaction:**
   - Users can view other dating profiles.
   - Users can swipe left (pass) or swipe right (like) on up to 10 different dating profiles in total (pass + like) within a single day.
   - The same profiles should not appear twice in the same day.

3. **Premium Packages:**
   - Implement a premium package system.
   - Users can purchase premium packages that unlock one premium feature of your choosing.
   - Examples of premium features include:
     - No swipe quota for the user.
     - A verified label for the user.

## Technologies Used

- [Golang](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)

## Clean Code Practices
![Clean Code Structure](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

This project follows the Clean Architecture pattern, where the code structure is well-organized into clear layers, including Entities, Use Cases and Interface Adapters.Complete references can be found at [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch).


## Getting Started

Follow these steps to get a copy of the project up and running on your local machine for development and testing purposes.

### How To Run This Project
> Database will auto migrate when services run
variable environent can be edit at .env file

> docker and docker-compose must be installed

### **Clone the repository:**

```bash
git clone git@github.com:akmalulginan/carjod-be.git
```

#### Run using docker-compose

```bash
$ docker-compose build
```
```bash
$ docker-compose up -d
```

#### Run the Testing

```bash
$ go test ./...
```

## API Documentation
### [Link to Postman Collection](https://www.postman.com/glyndrwn/workspace/carjod/collection/3857959-72b4aaca-8ffe-4f2b-a850-e45e7dae286a)
Explore and test the Dating App Backend API using [Postman](https://www.postman.com/glyndrwn/workspace/carjod/collection/3857959-72b4aaca-8ffe-4f2b-a850-e45e7dae286a). The provided Postman collection offers a detailed overview of endpoints, enabling you to interact with and understand the backend system.

Feel free to experiment with different requests, parameters, and payloads for a hands-on experience. Check the Postman collection's documentation for guidance on API functionalities.

For any questions or assistance, refer to the API documentation in Postman or reach out to the project contributors.

Happy exploring!