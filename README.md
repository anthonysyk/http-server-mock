# http-server-mock
Tool to run a local server with mocked data from a config file

### Configuration File

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
          filepath: "payload/top-movies.json"
  /movies/{id}:
    GET:
      responses:
        200:
          filepath: "payload/movies/{id}.json"
  /movies/{id}/actors/{id}:
    GET:
      responses:
        200:
          filepath: "payload/movies/{id}/actors/{id}.json"
  /popular-actors:
    GET:
      responses:
        200:
          filepath: "payload/top-actors.json"
  /genres:
    GET:
      responses:
        200:
          body: |
            ["Action","Adventure","Animation","Comedy","Crime","Drama","Family","Fantasy","Music","Science Fiction","Thriller"]
```

You can either set data as : 
- a filepath
- an inline string

### Use Cases

- To run a mock http server for developing, testing or debugging with mocked data
- To mock an API you don't have access to
- To mock response payloads and match a specific environment
- To implement integration tests