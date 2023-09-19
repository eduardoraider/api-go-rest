# Go REST API with Postgres and GORM

This project offers a comprehensive learning experience in building a REST API using native Go packages and the powerful ORM tool, GORM. With this repository, you'll gain insights into crafting a robust REST API, including seamless integration with a Postgres database deployed within a Docker container. Additionally, you'll explore techniques for crafting middleware to enhance codebase efficiency and avoid redundancy. Furthermore, this project will provide expertise in integrating your Go API with a React front-end.

## Frontend Application

Check out the frontend of this REST API in the following repository, where a React application consumes data from this Go API: [React Frontend for Go REST API](https://github.com/eduardoraider/frontend-react-api-go-rest)

This application is designed to record the names and histories of celebrities honored with street names.

## Learning Topics

By studying this application, you can learn the following topics related to Go programming and REST API development:

1. **Go: Developing a REST API**: Gain practical experience in developing a RESTful API using the Go programming language.

2. **Integrate Your Go API with a Database Running on Docker**: Learn how to integrate your Go API with a Postgres database deployed within a Docker container.

3. **Learn How to Use GORM, Go's Most Famous ORM**: Explore GORM, a popular Object-Relational Mapping (ORM) tool for Go, and leverage it to interact with your database.

4. **Creating, Deleting, and Editing with Gorm**: Master the CRUD (Create, Read, Update, Delete) operations in your Go API using GORM.

5. **Learn How to Create Middleware and Avoid Duplicate Code**: Discover the importance of middleware in enhancing codebase efficiency and avoiding redundancy.

6. **Middleware and Integrate Your Go API with a React Frontend**: Extend your knowledge of middleware while seamlessly integrating your Go API with a React frontend.

7. **JSON, Request, Response, and Go**: Understand how to work with JSON data, handle HTTP requests, and manage responses in a Go REST API.

8. **Router, Resources by ID, and Docker**: Implement routing, access resources by ID, and leverage Docker for database deployment.

9. **Learn Basic Layout for Go Application Projects**: Explore the basic project layout and organization for Go applications.

## Running the Application

To run this Go REST API application, follow these steps:

1. Ensure you have Go installed on your system. If not, you can download and install it from the official [Go website](https://golang.org/dl/).

2. Make sure you have Docker installed and running on your system.

3. Open a terminal and navigate to the project directory.

4. Run the following Docker Compose command to set up the Postgres database and pgAdmin:

   ```bash
   docker-compose up
   ```

   This command will create a Postgres database running on port 5432, import a database from `docker-database-initial.sql`, and install pgAdmin4 on port 54321. You can access pgAdmin at [http://localhost:54321](http://localhost:54321).

5. Open another terminal and run the following command to start the Go API:

   ```bash
   go run main.go
   ```

   This will execute the `main.go` script, which contains the Go API application. The application will start a server on port 8000.

6. To stop both the application and the server, use the `Ctrl + C` keyboard shortcut in both terminals.

7. Check the backend of the Go REST API in your web browser by navigating to:

   [http://localhost:8000/](http://localhost:8000/)


## Using Postman

You can perform all CRUD operations using the provided endpoints on Postman. Here are some cURL commands to get you started:

### Get All Personalities
```bash
curl --location 'http://localhost:8000/api/personalities'
```

### Get Personality by ID
```bash
curl --location 'http://localhost:8000/api/personalities/1'
```

### Create Personality
```bash
curl --location 'http://localhost:8000/api/personalities' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Nikola Tesla",
    "history": "Nikola Tesla was a Serbian-American inventor, electrical engineer, \
    mechanical engineer, and futurist best known for his contributions to the design \
    of the modern alternating current electricity supply system."
}'
```

### Edit Personality by ID
```bash
curl --location --request PUT 'http://localhost:8000/api/personalities/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ada Lovelace",
    "history": "Augusta Ada King, Countess of Lovelace was an English mathematician \
    and writer, chiefly known for her work on Charles Babbage'\''s proposed mechanical \
    general-purpose computer, the Analytical Engine. She was the first to recognise \
    that the machine had applications beyond pure calculation."
}'
```

### Delete Personality by ID
```bash
curl --location --request DELETE 'http://localhost:8000/api/personalities/1'
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.

---

#### by Eduardo O Raider
🛠 🥋 **Software Engineer**