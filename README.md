# new-newt

Domain availability checker

## Usage

### docker-compose

```yaml
version: "3.9"
services:
  redis:
    image: redis:alpine
    volumes:
      - ./data/redis/:/data
    ports:
      - 6379:6379

  new-newt-crawler:
    build:
      dockerfile: Dockerfile
    command: crawler
    env_file:
      - ./.env
    depends_on:
      - redis

  new-newt-web:
    build:
      dockerfile: Dockerfile
    command: web
    env_file:
      - ./.env
    depends_on:
      - redis

```

### add domains

#### single domain

```bash
docker-compose exec new-newt-crawler new-newt add example.de
```

To delete the domain again use:

```bash
docker-compose exec new-newt-crawler new-newt delete example.de
```

#### bulk import

`./example-domains.txt`

```text
domain.de
domain1.de
domain2.de
domain3.de
domain4.de
```

```bash
docker cp example-domains.txt new-newt_new-newt-crawler_1:./example-domains.txt
docker-compose exec new-newt-crawler new-newt import file example-domains.txt
```

## clear

clear all domains

```bash
docker-compose exec new-newt-crawler clear
```

## Development

### Make Targets

```bash
❯ make
build-docker                   build docker image
run-docker                     run project
test                           test go code
```
