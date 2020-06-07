Simple REST api to test the [Amazon EC2](https://docs.aws.amazon.com/ec2/index.html) service.

## Running

The only requirement is [Golang](https://golang.org/dl) and/or [Docker](https://docs.docker.com/get-docker/).

### Golang

Run ```go run .```

### Docker

Run ```docker build -t microservice-test .```

Then ```docker run --rm -p 8080:8080 microservice-test```

## Endpoints

| Path          | Method    | Response         |
| ------------- | --------- | ---------------- |
| ```/```       | ```GET``` | ```text/plain``` |
| ```/thanks``` | ```GET``` | ```text/plain``` |