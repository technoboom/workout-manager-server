# Backend for Workout Manager App

## Getting started
Clone the repository:
`git clone https://github.com/technoboom/workout-manager-server`

### Commands:
Start services: `docker-compose up -d`.

Stop services: `docker-compose stop`.

To rebuild the images you must use `docker-compose build
` or `docker-compose up --build`

## Stack:
1. Golang
2. PostgreSQL
3. Angular 4
4. Docker and Docker Compose
5. Redis

## Services:
1. Users - stores users information, provides RESTFul API for users collection
2. Web - interface for managing personal information (for customer)
