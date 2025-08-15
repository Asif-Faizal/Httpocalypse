# Httpocalypse - TCP to http

An educational HTTP/1.1 parser and server built from scratch in Go, focusing on core protocol mechanics by handling TCP connections, parsing requests, and crafting responses without high-level libraries. A hands-on way to learn web transport fundamentals while creating a functional server

## Prerequisites

- Understanding of the Go programming language. If you're fuzzy on that, have we got a course for you.
- What binary is, which is covered in our Learn to Code in Python course.
- Understanding of how to use HTTP, because in this course we want to build it. We have an HTTP Clients course and an HTTP Servers course that can help with that.

## Learning Goals

Understand the big ideas of HTTP, and implement the HTTP/1.1 protocol. (Look, the RFC is long, but the main points are fairly straightforward). You'll understand from a high level what happens when you fetch("google.com").

## File Parsing Fundamentals

The project includes a foundational example that demonstrates core concepts needed for HTTP server development. Here's the complete code:

```go
package main

import (
 "fmt"
 "log"
 "os"
)

func main() {
 f, err := os.Open("message.txt")
 if err != nil {
  log.Fatal(err)
 }
 for {
  buf := make([]byte, 8)
  n, err := f.Read(buf)
  if err != nil {
   log.Fatal(err)
  }
  if n == 0 {
   break
  }
  fmt.Printf("read: %s\n", string(buf[:n]))
 }
}
```

This code demonstrates:

### Buffer-Based Reading

```go
buf := make([]byte, 8)
n, err := f.Read(buf)
```

- Reads data in fixed-size chunks (8 bytes in this example)
- Similar to how HTTP servers read from TCP connections
- Prevents memory issues with large requests

### Stream Processing

```go
for {
    // Read chunk
    if n == 0 {
        break  // End of stream
    }
    // Process chunk
}
```

- Processes data incrementally as it arrives
- Continues until end-of-stream is detected
- Essential for handling HTTP requests of unknown size

### Connection to HTTP Development

This file reading pattern directly translates to HTTP server development:

- **File → TCP Connection**: Replace file reading with socket reading
- **Buffer Management**: Handle HTTP headers and body in chunks
- **Stream Processing**: Parse requests as they arrive over the network
- **Error Handling**: Manage network timeouts, disconnections, and malformed requests

The `message.txt` file contains test data that demonstrates how the server will handle different types of request content.
