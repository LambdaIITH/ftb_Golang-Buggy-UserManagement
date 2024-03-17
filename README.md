<!-- only include .env, go.main, go.mod, swagger.yml -->

# Problem Statement - User Management System
Max Verstappen was a developer at [Lambda](https://iith.dev) who created new user management system for [Elan](https://elan.org.in), but the system is plagued with lot of bugs in it. As the fest comes nearer you are tasked with fixing the bugs associated with the user management system. The system is crucial for smooth conduction of the fest.

Checkout the [OpenAPI specification](swagger.yml) file and the associated [go code](main.go) to fix the issues. The system should comply with the API specification.

## Installation and Running

To install and run the Go code, follow these steps:

1. Clone the repository and Navigate to the project directory.
2. Make sure you have Go installed on your system.
3. Run `go run go.main` to start the server.
4. Access the API at `http://localhost:8000`.

## Notes
The provided code is a simplified version aimed at illustrating a user management system, it does not include a database for simplicity. You need not setup databases as it is not part of problem statement. Expected bugs include syntax errors, HTTP handling issues and so on. 
