# http-server-mock
Tool to run a local server with mocked data from a config file

## Route Definitions
You can define routes in 2 different ways :

### 1 - Configuration File

Very handy to **quickly** set up a mock server and define everything in a single file.

```yaml
version: 1.0
paths:
  /health:
    GET:
      responses:
        200:
          body: |
            {"status":"OK"}
  /recommended-movies:
    GET:
      responses:
        200:
          filename: "payload/recommended-movies.json"
  /movies/{id}:
    GET:
      responses:
        200:
          filename: "payload/movies/{id}.json"

```

### 2 - File System Structure

Very handy when you have a **lot of data to mock**, and you want to keep it humanly readable.

### Use Cases

- To run a mock http server for developing, testing or debugging with mocked data
- To mock an API you don't have access to
- To mock response payloads and match a specific environment
- To implement integration tests