Simple REST api to test the [AWS Elastic Beanstalk](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/Welcome.html) service.

## Running

The only requirement is [Golang](https://golang.org/dl) and/or [Docker](https://docs.docker.com/get-docker/).

### Golang

Run ```go run .```

### Docker

Run ```docker build -t microservice-test .```

Then ```docker run --rm -p 8080:8080 microservice-test```

<!-- TODO: Running on Docker -->

## Endpoints

| Path          | Method    | Response         |
| ------------- | --------- | ---------------- |
| ```/```       | ```GET``` | ```text/plain``` |
| ```/thanks``` | ```GET``` | ```text/plain``` |