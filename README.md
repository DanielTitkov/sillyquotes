# Silly Quotes

## Overview

This project is a TCP server that provides a silly quote (joke) to clients after they solve a Proof of Work challenge. This approach helps to protect the server against DDoS attacks. The server and client both run in Docker containers for easy deployment.

## Requirements

- Go 1.20 or higher (for native)
- Docker (for containerized)

## How to Run

### Native

1. To start the server, run:  
    ```bash
    make server
    ```

2. To start the client, run:  
    ```bash
    make client
    ```

### Docker

1. To start the server container, run:  
    ```bash
    make docker-server
    ```

2. To start the client container, run:  
    ```bash
    make docker-client
    ```

## Proof of Work Algorithm

The server uses a custom challenge-response Proof of Work algorithm. Normally, one might employ a more traditional self-imposed approach like Hashcash for the client to prove its computational efforts. However, the choice to use a challenge-response mechanism was guided by the technical specifications of the project.

### Implementation Details

The implemented Proof of Work algorithm is fairly straightforward:

1. The server generates a random challenge string and sends it to the client.
2. The client needs to find a string that, when concatenated with the challenge and hashed using SHA-256, produces a hash with a predetermined number of leading zeros.
3. The client sends the string back to the server.
4. The server verifies the string against the hash and either grants or denies access.

#### Why This Implementation?

1. **Simplicity**: This approach keeps both the client and server implementations simple, enabling easier testing, debugging, and scaling.
2. **Rate Limiting**: The algorithm is effective for rate-limiting requests, which protects the server against DDoS attacks.
3. **Customizability**: By simply altering the length of the challenge string or the number of required leading zeros, the difficulty of the problem can be adjusted.
4. **Сomputational asymmetry**: Solving the challenge requires a substantial amount of computational power and time, as the client needs to perform numerous hash computations to find a solution that meets the criteria. This deters DDoS attacks by making it computationally expensive for attackers to spam the server with requests. On the other hand, verifying the solution is computationally trivial for the server. Once the client sends back the string, the server can quickly hash it along with the challenge and verify whether it meets the criteria. This allows the server to maintain high performance while verifying incoming requests, thereby ensuring that legitimate clients are not adversely impacted.
5. **Adherence to Specifications**: This implementation was chosen to strictly meet the challenge-response requirement set forth in the project's technical specifications.

## Makefile Commands

- `make server`: Starts the native Go server
- `make client`: Starts the native Go client
- `make create-network`: Creates a Docker network for communication between containers
- `make docker-server`: Builds and starts the server Docker container
- `make docker-client`: Builds and starts the client Docker container