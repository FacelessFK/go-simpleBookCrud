# Golang CRUD Application

This project is a simple CRUD application built using Golang with a PostgreSQL database. Docker is used for containerizing and running the services.

## Prerequisites

-   [Docker](https://www.docker.com/)
-   [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### 1. Clone the repository

First, clone the project from your repository:

```bash
git clone https://github.com/FacelessFK/go-simpleBookCrud.git
cd go-simpleBookCrud
```

Use Docker Compose to build and start the services:

```bash
docker-compose up --build
```

then create test database

```bash
docker-compose exec db psql -U postgres -d test
```

if you cant find db container use :

```bash
docker compose ps
```

than copy the name of container db and put here:

```bash
docker-compose exec <YOUR_DB_CONTAINER_NAME> psql -U postgres -d test
```

Then, run the following SQL command to create the book table:

```bash
CREATE TABLE public.book(
    id serial4 NOT NULL,
    "name" varchar NULL,
    CONSTRAINT book_pk PRIMARY KEY (id)
);
```
