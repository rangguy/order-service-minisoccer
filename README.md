<h3>Order Service</h3>

<h3>Description</h3>

<p>This repository will be used to manage order</p>

<h3>Directory Structure</h3>

```
order-service
    L client                         → Contains the client for calling other services
    L cmd                            → Contains the main entry point or initial configuration of the application
    L common                         → Stores common functions used throughout the application
    L config                         → Contains application configurations such as environment variables and other settings
    L constants                      → Stores global constant values used across the application
    L controllers                    → Manages control logic for handling HTTP requests
    L domain                         → The application's domain module containing core domain elements
        L dto                        → Data Transfer Objects, used to define the structure of transferred data
        L models                     → Object models representing the application's or database's data structure
    L middlewares                    → Contains middleware for processing requests/responses before or after reaching the controller
    L repositories                   → Contains data access logic for interacting with the database
    L routes                         → Contains API route definitions
    L services                       → Stores the application's core business logic
```

## How to setup

```
- Clone this repository
- go mod tidy
- copy .env.example to .env (if you want to run with consul)
- copy .config.json.example to .config.json
```

## How to run

```bash
make watch-prepare (only for the first time or when you add new dependency)
make watch
```

## How to run with docker

```bash
docker-compose up -d --build --force-recreate
```

## How to build
```bash
make build
```

