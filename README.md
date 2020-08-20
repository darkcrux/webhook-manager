
This is a prototype of a callback manager

## Relevant Docs

- [High Level Design Diagram](https://drive.google.com/file/d/1W2YoWduYUCeCu1H_IqT0eWtww7ZOnFj-/view?usp=sharing)
- [High Level Design Diagram - messagebus](https://drive.google.com/file/d/19ZDiUA9L2XLS_BwJMh0VHgldwSWVpXgw/view?usp=sharing)
- [ERD](https://drive.google.com/file/d/1QchuqKYSEu0D_vJBPYayy33KU3xEfO5d/view?usp=sharing)

## Getting Started

To clone the repo:

```shell
git clone https://github.com/darkcrux/webhook-manager
cd webhook-manager
```

## Prerequisire

### Workspace

- Go
- Docker

### Deployment

- Docker Compose

### Important Documents

- [API Spec](https://github.com/darkcrux/webhook-manager/blob/master/openapi.yaml)

## Build

Run the following to build the binary:

```shell
make build
```

## Run

To run the application (eg. databases, service, etc.):

Copy the `.env.sample` to `.env` and change configuration as needed.

```
make run
```

The following services will be up after a few seconds/minutes:

- Postgres: http://localhost:5432
- Swagger UI: http://localhost:9999
- Swagger Editor: http://localhost:9998

To build the binary:

```
make build
```

## Packaging Image

To create the docker image:

```
make package
```

To publish the image to docker hub:

```
make publish
```

Note that publishing the image requires access to the registry.

## Development

// TODO

### Project Structure

```
webhook-manager
|- build/               # build artifacts are generated here
|- cmd/                 # command line commands live here. Checkout cobra library
|- config/              # configuration files are here
|- db/                  # for database migration files
|- internal/            # for internal go packages 
| |- component/         # the actual domains / business logic live here including the ports for integration
| |- entrypoint/        # entrypoint for the application, the REST API handlers are here
| |- infrastructure/    # the third party infrastructure the service needs (eg. database, etc)
|- pkg/                 # for public go packages
|- .dockerignore        # ignore list for docker
|- .gitignore           # ignore list for git
|- go.mod               # dependencies for project
|- go.sum               # checksum for dependencies, do not manually change
|- main.go              # the main go file
|- Makefile             # build scripts
|- README.md            # this file
```

### Adding Dependencies

To add dependencies, run the following:

```shell
go get -u {dependency}
make deps
```

## TODO

// Also TODO

