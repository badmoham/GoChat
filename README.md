# Go Chat App

## Description
This is a simple RESTful chat application written in Go as part of my journey to learn web development. The project uses Docker and Docker Compose for containerization, making it easy to set up and run. It serves as a platform to explore building APIs, working with databases, and implementing key web development concepts. Through this exploration, I aim to deepen my understanding of Go and its capabilities in creating robust web applications.

## Features
- RESTful API for chat functionality.
- Database integration (e.g., PostgreSQL).
- User authentication and message handling.
- Modular code structure for maintainability.
- Containerized deployment using Docker and Docker Compose.

## Technologies Used
- **Language**: Go (Golang)
- **Framework**: Gin (for building the API)
- **ORM**: GORM (for database interactions)
- **Database**: PostgreSQL (or another relational database)
- **Containerization**: Docker & Docker Compose

## Project Status
This project is currently in its early stages and is primarily a learning exercise. It may evolve over time as I continue to develop my skills in Go and Docker.

## How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/badmoham/GoChat
   ```
2. Navigate to the project directory:
3. Build and start the application using Docker Compose:
   ```
   docker-compose up --build
   ```
4. Access the API at `http://localhost:8080` (or the port specified in your `docker-compose.yml`).

### Notes:
- Ensure Docker and Docker Compose are installed on your system.
- The database will be automatically set up using the configuration in `docker-compose.yml`.

## Future Plans
- Implement real-time messaging using WebSockets or gRPC.
- Add more features like user profiles and group chats.
- Optimize performance and improve code quality.
- Explore advanced Docker configurations for scalability.

## License
This project is for personal learning purposes only and does not have an official license. Feel free to use it as a reference for your own learning journey.

---

Thank you for visiting this repository! If you have any questions or suggestions, feel free to reach out.