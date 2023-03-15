# Yet Another ChatGPT Web

[中文](README_CN.md)

## Init

### Install Pre Commit Check 

#### Install `golangci-lint`

[Install tutorial](https://golangci-lint.run/usage/install/#local-installation)

#### Install `pre-commit`

```bash
pip install pre-commit
```

#### Run `pre-commit`

```bash
pre-commit install
```

### Install Go Dependencies

#### Using `make`

```bash
make init
```

#### Using Command

```bash
go mod download
```

## Run

### Using `make`

```bash
make run
```

### Using Command

```bash
go run main.go
```

## Project Structure

```
├── main.go # Entry point of the program
├── .env # Environment variables file
├── controllers # Controllers directory
│   ├── user.go # User controller
│   └── product.go # Product controller
├── models # Models directory
│   ├── user.go # User model
│   └── product.go # Product model
├── routes # Routes directory
│   └── routes.go # Routes definition file
├── services # Services directory
│   ├── user.go # User service
│   └── product.go # Product service
├── tests # Test directory
└── utils # Utility classes directory
    ├── logger.go # Logging utility class file
    └── db.go # Database connection utility class file
```