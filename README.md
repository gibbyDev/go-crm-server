# Project go-crm-server

This is a set of go microservices. Each of which are dockerized in their own containers with their own database and they all communicate with one another. The first service is an auth service for sign up, sign in, and sign out with email verification. The second is an api that handles an IT ticketing system. The third is a video & text chat app using WebRTC 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
