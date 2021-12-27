# Go Svelte Project

This is just a little project I am working on in my spare time primarily for learning purposes.

I want to improve my working knowledge of:

  - Go, and Go HTTP requests.
  - Docker containers.
  - Postgres (or some other NoSQL Db, haven't decided yet)
  - Auth with JWT.
  - Svelte (I haven't used this yet, but might be run to try - it may end up being React/Vue)
  - Some CSS framework (probably not Bootstrap or Tailwind)

## Docker

A big reason for this was to improve my experience of Docker.

To build the docker container: `docker build --tag <image name> .`

To run the docker image: `docker run --publish 8080:8080 <image name>` where the ports are `[host-port]:[container-port]` so for example `3000:8000` would expose port `8000` inside the container to `3000` outside the container.