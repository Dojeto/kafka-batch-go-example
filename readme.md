# Kafka Batch Processing Go Example

Welcome to the **Kafka Batch Processing Go Example** project! This project is designed to demonstrate a basic setup for a Kafka producer and consumer using Go. The aim is to provide a straightforward example of how to Batch Process messages from Kafka, leveraging the `Sarama` library in Go. This project could serve as a starting point for anyone looking to integrate Kafka into their Go applications.

## Features

- **Producer**: A simple HTTP server using Gin that reads data from a CSV file and sends it to a Kafka topic.
- **Consumer**: A Kafka consumer that processes messages in batches and writes them to a database.
- **Debouncing**: Implements a debouncing technique to efficiently process unbatched messages.

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+ recommended)
- [Kafka](https://kafka.apache.org/quickstart) (running locally or in Docker)
- Docker (facing issue to containerized the application)


### Contributing
We welcome contributions to enhance this project! Currently I am facing Issue of Dockerized this App.If Some who know how to dockerized it, then help would be appriciated :\)


### Note

Currently Dockerized Version Of This App Not Working.