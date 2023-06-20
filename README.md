# http-server-mock
Tool to run a local server with mocked data from a config file

### Configuration File

Very handy to **quickly** set up a mock server and define everything in a single file.

```yaml
version: 1.0
paths:
  /health:
    GET:
      200:
        body: |-
          {"status":"OK"}
  /top-movies:
    GET:
      200:
        filepath: "payload/top-movies.json"
  /movies/{id}:
    GET:
      200:
        filepath: "payload/movies/{id}.json"
  /movies/{moviesId}/actors/{actorsId}:
    GET:
      200:
        filepath: "payload/movies/{moviesId}/actors/{actorsId}.json"
  /movies/suggest:
    GET:
      200:
        body: |-
          {
            "title": "The Godfather",
            "description": "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
            "duration": 175,
            "actors": [
              "Marlon Brando",
              "Al Pacino",
              "James Caan"
            ]
          }
  /top-actors:
    GET:
      200:
        filepath: "payload/top-actors.json"
  /genres:
    GET:
      200:
        body: |-
          ["Action","Adventure","Animation","Comedy","Crime","Drama","Family","Fantasy","Music","Science Fiction","Thriller"]
    POST:
      200:
        body: |-
          {"message":"new genre added"}
    DELETE:
      200:
        body: |-
          {"message":"genre deleted"}
```

You can either set data as : 
- a filepath
- an inline string

### Usage

```go
//go:embed routes.yml
var RoutesYAML string

func main() {
    router, err := http_server_mock.GenerateRouter(RoutesYAML)
    if err != nil {
    // handle error
    }
    http.ListenAndServe(":8080", router)	
}
```

Real usage here : https://github.com/anthonysyk/http-backtest/blob/main/example/httpbacktest_test.go

### Use Cases

- To run a mock http server for developing, testing or debugging with mocked data
- To mock an API you don't have access to
- To mock response payloads and match a specific environment
- To implement integration tests

### Troubleshoot

Warnings can be raised if :
- Path parameters definitions must be unique per URL
  - Good : `/movies/{moviesId}/actors/{actorsId}`
  - Bad : `/movies/{id}/actors/{id}`
- Routes must be defined only once (no duplicate)
