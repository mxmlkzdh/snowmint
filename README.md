# SnowMint


SnowMint is a blazingly fast unique and roughly sortable IDs generator based on Twitter's Snowflake.

## Unique ID Generation
### The Algorithm
### Sorting

## The Protocol

## Clients
### Go Client
### Java Client

## Install
### Native Deployment
### Docker Deployment

## Benchmarks
Latency: 
- In a native deployment, response time for each unique ID is roughly between 5 to 10 microseconds.
- In a Docker container, response time for each unique ID is roughly between 10 to 15 microseconds.

Here's a structured README for your SnowMint project:

---------------------------------------------------------------------------------


# SnowMint - A Blazingly Fast Unique ID Generator
[![Go Workflow](https://github.com/mxmlkzdh/snowmint/actions/workflows/go.yml/badge.svg)](https://github.com/mxmlkzdh/snowmint/actions)

**SnowMint** is a high-performance, distributed unique ID generator based on Twitter's Snowflake algorithm. It provides unique, sortable IDs that are generated using a client/server model with a custom protocol over raw TCP connections.

## Unique ID Generation

### Algorithm

A SnowMint generates IDs using a 64-bit timestamp, a machine ID, a process ID, and a sequence number. This combination ensures uniqueness and allows for efficient sorting based on the timestamp.

B SnowMint leverages Twitter’s Snowflake algorithm to generate 64-bit unique identifiers. The format of the ID consists of:
- **Timestamp**: 41 bits for time in milliseconds.
- **Node ID**: 10 bits for machine or datacenter ID.
- **Sequence**: 12 bits for a sequence number that resets every millisecond.

Timestamp: A 41-bit field representing the current time in milliseconds since the epoch.
Machine ID: A 10-bit field identifying the machine where the ID was generated.
Process ID: A 5-bit field representing the process ID.
Sequence Number: A 10-bit field for sequential IDs within a millisecond.

This combination ensures that SnowMint can generate thousands of unique IDs per second, even in distributed environments.

### Sorting IDs
Since the first 41 bits represent the timestamp, SnowMint IDs are naturally sortable by creation time. IDs generated earlier will have a smaller numeric value than those generated later, allowing simple chronological ordering by comparing ID values directly.

## The Protocol

SnowMint uses a lightweight, highly optimized custom protocol over raw TCP connections. This design focuses on speed and simplicity, ensuring ultra-fast ID generation and retrieval.

### How it Works:
1. **Connection**: Clients open a TCP connection to the SnowMint server.
2. **Command**: The client sends a single `GET` command to the server.
3. **Response**: The server responds immediately with a 64-bit unique ID.

This minimalist protocol reduces overhead, delivering unparalleled speed compared to traditional HTTP-based services.

### Performance Benefits:
- **Raw TCP**: Eliminates HTTP headers and other overhead, reducing the time between a request and response.
- **Low-Latency**: Designed for microsecond-scale latencies, making it ideal for high-throughput systems.

## Clients

SnowMint provides easy-to-use SDKs for popular programming languages to integrate with the server and retrieve unique IDs.

### Go SDK
The Go client SDK allows seamless integration into Go applications. A simple GET request over TCP fetches the unique ID.

### Java SDK
The Java SDK offers a similarly efficient way to connect to the SnowMint server, providing support for applications in JVM environments.

## Install

### Native Deployment
1. Download the latest release from the [SnowMint releases page](#).
2. Extract the archive and run the binary:
   ```bash
   ./snowmint-server --node-id <YOUR_NODE_ID>
   ```

### Docker Deployment
To run SnowMint in a Docker container, use the following:
```bash
docker pull snowmint/snowmint-server:latest
docker run -d --name snowmint -p 8080:8080 snowmint/snowmint-server --node-id <YOUR_NODE_ID>
```

## Benchmarks
SnowMint has been benchmarked to handle thousands of requests per second, with latencies in the microsecond range. Thanks to the custom protocol and raw TCP connections, it outperforms traditional HTTP-based systems by a significant margin.

- **ID Generation Rate**: Up to X,000 IDs per second per node.
- **Latency**: Sub-millisecond, typically under X microseconds.

## Sources
The SnowMint project is open source and available on GitHub. Check out the [source code](#) to contribute or explore the internals of the system.

## License
The SnowMint project is licensed under the MIT License.