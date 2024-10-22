# To-Do List Application

## Overview

The **To-Do List Application** is a full-stack web project designed to manage tasks effectively. It allows users to create, read, update, and delete tasks. The project is built using modern technologies, aiming for scalability, maintainability, and ease of development.

## Features

- Add, edit, and delete tasks.
- Mark tasks as completed or pending.
- View a list of all tasks with status indicators.
- Responsive user interface for both desktop and mobile.

## Technologies Used

### Backend (Go with Gin):
- **Go**: A statically typed programming language known for its simplicity and performance.
- **Gin**: A web framework for Go designed to build high-performance REST APIs.
- **PostgreSQL**: A powerful, open-source object-relational database system.
- **Bun ORM**: A lightweight SQL-first ORM for Go, used for querying PostgreSQL.
- **Docker**: Used for containerizing the application to ensure consistency across different environments.

### Frontend (Vue.js with Nuxt.js and TypeScript):
- **Vue.js**: A progressive JavaScript framework used for building the user interface.
- **Nuxt.js**: A framework built on top of Vue.js to provide server-side rendering, static site generation, and powerful routing features.
- **TypeScript**: A superset of JavaScript that adds static typing, improving the maintainability and scalability of the codebase.
- **Tailwind CSS**: A utility-first CSS framework for designing a responsive and modern UI.

### API Design:
- **RESTful API**: The backend API is designed following REST principles to handle CRUD operations (Create, Read, Update, Delete) for tasks.
- **JSON Format**: All data is exchanged in JSON format between the frontend and backend.

## Project Structure

- **Backend**: Written in Go, using Gin for routing and Bun ORM for database interaction.
- **Frontend**: Written in Vue.js, with Nuxt.js for SSR and TypeScript for static typing.
- **Database**: PostgreSQL to store task data.

## Setup and Installation

### Backend Setup
1. Clone the repository and navigate to the backend directory:
    ```bash
    git clone https://github.com/yourusername/todo-list-go-vuejs.git
    cd todo-list-go-vuejs/backend
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Configure your PostgreSQL database in `main.go` or a `.env` file:
    ```go
    dsn := "postgres://user:password@localhost:5432/todolist?sslmode=disable"
    ```
4. Run the backend:
    ```bash
    go run main.go
    ```

### Frontend Setup
1. Navigate to the frontend directory:
    ```bash
    cd ../frontend
    ```
2. Install dependencies:
    ```bash
    npm install
    ```
3. Run the frontend:
    ```bash
    npm run dev
    ```

### Running with Docker
To set up the entire project with Docker, use the provided `docker-compose.yml` file.

1. Start the application:
    ```bash
    docker-compose up --build
    ```

2. The backend will run on `http://localhost:8080`, and the frontend on `http://localhost:3000`.

## Contributions

Feel free to fork this repository, create a new branch, and make improvements. Pull requests are welcome!

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
