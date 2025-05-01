# Terraform Provider: Rubber Duck

This Terraform provider is a simple example designed to manage resources from the [Rubber Duck API](https://github.com/kaerimichi/rubber-duck-api). It demonstrates how to interact with a RESTful API using Terraform, providing basic CRUD operations for managing rubber duck resources.

## Features

This provider allows you to:
- **Create** a new rubber duck resource.
- **Read** details of an existing rubber duck by its ID.
- **Update** an existing rubber duck's attributes.
- **Delete** a rubber duck resource.

## Prerequisites

- A running instance of the [Rubber Duck API](https://github.com/kaerimichi/rubber-duck-api).
- Terraform installed on your machine.

## Installation

1. Build the provider binary:
   ```bash
   make build
