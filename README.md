
# **cajita: a simple key-value store*

## **Overview**
cajita is a basic key-value store implemented in Go. It exposes a simple HTTP API to store, retrieve, and delete key-value pairs. I built this project as an introductory project to build a go-based service.

## **Features**

- **Set Key-Value Pairs:** Store a key-value pair in the in-memory store.
- **Get Values:** Retrieve the value associated with a given key.
- **Delete Key-Value Pairs:** Remove a key-value pair from the store.
- **Concurrent Access:** The store is thread-safe and handles concurrent requests.
- **Command-Line Interface (CLI):** Interact with the store through a simple HTTP API.

## **Getting Started**

### **Installation**

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/guillego/cajita.git
   cd cajita
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

### **Running the Application**

To start the cajita kv-store service:

```bash
go run cmd/cajita/main.go
```

By default, the service runs on `http://localhost:8080`.

### **API Endpoints**

- **Set a Key-Value Pair:**

  ```http
  GET /set?key=<key>&value=<value>
  ```

  Example:

  ```bash
  curl "http://localhost:8080/set?key=mykey&value=myvalue"
  ```

- **Get a Value by Key:**

  ```http
  GET /get?key=<key>
  ```

  Example:

  ```bash
  curl "http://localhost:8080/get?key=mykey"
  ```

- **Delete a Key-Value Pair:**

  ```http
  GET /delete?key=<key>
  ```

  Example:

  ```bash
  curl "http://localhost:8080/delete?key=mykey"
  ```
