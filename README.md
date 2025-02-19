# Листовки за курса на МА "Водач на малък кораб по Дунав"

This project provides a web and backend solution for processing leaflets issued by the Maritime Administration for examinations.

**Note:** Листовките **НЕ СА** за курса Водач на кораб до 40БТ.

## How to Run

### Clone the Repository

```bash
git clone https://your-repository-url.git
```

### Using Dockerfile

1. **Build the Docker Image**

    ```dockerfile
    docker build -t reka .
    ```

2. **Run the Docker Container**

    ```dockerfile
    docker run -d -p 1111:8080 --name reka reka:latest
    ```

### Using Docker Compose

```dockerfile
docker-compose up -d
```

### Access the Application

Open your browser and navigate to:
[http://localhost:1111](http://localhost:1111)

*Replace `localhost` with your actual IP address if the application is installed on a remote server.*

---