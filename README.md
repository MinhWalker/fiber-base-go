# Build a REST API from scratch with Go and Docker
start:
```
DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build
```

stop:
```
docker-compose down --volumes --remove-orphans
```