# Building Docker images

- Build the Docker image and run the container

    ```bash
    docker build -t sample-web-go -t sample-web-go:bookworm .
    docker run -p 8080:8080 sample-web-go
    ```

    or the alpine version

    ```bash
    docker build -t sample-web-go:alpine -f Dockerfile.alpine .
    docker run -p 8080:8080 sample-web-go:alpine
    ```

- Open a web browser and navigate to http://localhost:8080 to see the "Hello, world!" message.

- Navigate to http://localhost:8080/api to see the JSON response.


## Image size comparison

Bookworm | Alpine
-------- | ------
  108 MB |  18 MB
