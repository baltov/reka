# Stage 1: Build the Go Backend
FROM golang:1.23-alpine AS builder-backend
# Set the Current Working Directory inside the container
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev
WORKDIR /app

# Copy go.mod and go.sum files from the backend directory
COPY ./rekaGo/go.mod ./rekaGo/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY ./rekaGo/. .
# Build the Go application
RUN go build -ldflags='-s -w -extldflags "-static"' -o /app/rekaBo ./main.go
#COPY ./mydb.sqlite /app/mydb.sqlite

# Stage 2: Build the Frontend
FROM node:18-alpine AS builder-frontend

# Set the Current Working Directory inside the container
WORKDIR /frontend

# Copy package.json and package-lock.json from the frontend directory
COPY ./reka_fe/package.json ./reka_fe/package-lock.json ./

# Install frontend dependencies
RUN npm install

# Copy frontend source code
COPY ./reka_fe/. .

# Build the frontend application
RUN npm run build

# Stage 3: Create the Final Image with Nginx
FROM nginx:alpine

# Set environment variables (modify as needed)
ENV PORT=8080

# Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy the built frontend from the builder-frontend stage
COPY --from=builder-frontend /frontend/dist /usr/share/nginx/html

# Copy the built backend from the builder-backend stage
COPY --from=builder-backend /app/rekaBo /usr/local/bin/rekaBo
#COPY --from=builder-backend /app/mydb.sqlite /mydb.sqlite

# Copy custom nginx configuration
#COPY nginx.conf /etc/nginx/nginx.conf

# Expose the port on which the app will run
EXPOSE 8080
RUN export GIN_MODE=release
# Start both Nginx and the Go backend
CMD ["sh", "-c", "nginx && /usr/local/bin/rekaBo"]