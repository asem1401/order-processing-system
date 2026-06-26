# Event-Driven Order Processing System

A microservices-based order processing system developed in Go. The project demonstrates communication between services, retry mechanisms, and Dead Letter Queue (DLQ) handling using Docker.

## Architecture

The system consists of three independent services:

- Order Service
- Payment Service
- Notification Service

## Features

- Microservice architecture
- Event-driven communication
- Retry mechanism
- Dead Letter Queue (DLQ)
- Docker Compose deployment
- Independent services

## Technologies

- Go
- Docker
- Docker Compose

## Project Structure

```
order-service/
payment-service/
notification-service/
docker-compose.yml
```

## Learning Objectives

- Build distributed systems
- Implement reliable message processing
- Practice microservice architecture
- Improve fault tolerance using retry and DLQ
