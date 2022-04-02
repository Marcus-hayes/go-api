# GoLang Sandbox API

This is a Sandbox GoLang API in which I will experiment with GoLang's functionalities, interact with other API's, and make notes on interesting findings.

[JSON-> Struct (Or vice-versa) Converter](https://transform.tools/json-to-go)

## Pre-Requisites:

- [Install GoLang](https://go.dev/doc/install)
- [Install Docker](https://www.docker.com/get-started/)
- [Windows Only - Install & Enable WSL](https://docs.microsoft.com/en-us/windows/wsl/install)

## Development
The following commands will build an app container for the API using Docker Compose.
```
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```

## Testing
In Development

## API Reference
/
- Example Endpoint Call:
    - curl http://localhost:8080/

/joke
- Example Endpoint Call:
    - curl http://localhost:8080/joke

/movie/{movie+title}
- Example Endpoint Call:
    - curl http://localhost:8080/movie/dances+with+wolves