# Status checker for all services

Lightweight services monitoring for the association's infrastructure

## Architecture

```
.
├── cmd/server/         # Entry point
├── internal/
│   ├── config/         # YAML config loading
│   ├── checker/        # Checks implementation (ping, http)
│   ├── scheduler/      # Periodic orchestration
│   ├── storage/        # Results memory storage
│   └── api/            # RESTful http server
├── configs/            # configuration files
└── Dockerfile
```

## API

- `GET /status` - State of all services
- `GET /health` - Healthcheck

## Installation

```
git clone git@github.com:DaVinciCodeCTF/status-checker.git
```

## Configuration

### env file

There's only one environment variable that must be set. It's `ENCRYPTION_KEY` which must be a 32 bytes string.
It is used to cipher api output.

> [!NOTE]
> The same key must be used on whatever other service that'll reach this api in order to decipher its ouptut.

You can also optionally override the `PORT` environment variable which is defaulted to "8080".

### services configuration

 - Edit `configs/services.yaml.example` following the example.
> [!TIP]
> When finished editing, remove the *.example* extension.


## Run

First build the container, preferably using **podman** (docker will work as well):

```sh
podman build -t status-checker .
```

Then just run it while providing the **.env** file location & the forwarding port to use:

```sh
podman run -d --name status-checker --env-file <full_path_of_env_file> -p 0.0.0.0:<desired_port>:8080 status-checker:latest
```
