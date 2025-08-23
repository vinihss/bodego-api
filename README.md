# Bodego API

A Go-based REST API for managing products and customers.

## Architecture

The project follows Clean Architecture principles with the following layers:

- Domain: Core business entities and interfaces
- Use Cases: Application business rules and use case implementations
- Infrastructure: External interfaces implementations (not shown in current codebase)

## Domain Entities

### Customer

- ID (uint)
- Name (string)
- Email (string)

### Product

- ID (uint)
- Name (string)
- Price (float64)
- Description (string)

## Use Cases

### Customer

- Create: Creates a new customer
- Delete: Removes a customer by ID
- Find: Retrieves a customer by ID
- Update: Updates customer information
- FindAll: Lists all customers with pagination

### Product

- Create: Creates a new product
- Delete: Removes a product by ID
- Find: Retrieves a product by ID
- Update: Updates product information
- FindAll: Lists all products with pagination
