### URL-shortner Sevice in GO

#### Tech stack
* GoLang - web server 
    * gin http library
    * swaggo documentation
* postgres DB
* redis cache
* docker and docker-compose for devlopment and deployment
-----
#### How to run with docker
* create `.env` file at the root of the project with following vars.
```
    PORT=4444
    DB_HOST=postgres
    DB_PORT=5432
    DB_USER=postgres
    DB_PASS=pass
    DB_NAME=postgres
    REDIS_HOST=redis
    REDIS_PORT=6379
    START_SEQ=1000
    API_VERSION=v1
    SERVER_DOMAIN=localhost:4444/
```

* Install docker and  [docker-compose](https://docs.docker.com/compose/install/) for your distro.
* make sure docker daemon is up by running `docker` in cmd.
* go to project root dir and run `docker build -t go-dev:1` (or any name for the image).
* run `docker-compose up`.
* now ther service is accessible at port 4444.
----
#### Docs
* go get [swag](https://github.com/swaggo/swag).
* `swag init` (dont forget to export your go bin path).
* visit localhost:4444/docs/doc.json 
