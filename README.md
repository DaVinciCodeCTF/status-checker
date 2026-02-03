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

 - Edit `configs/services.yaml.example` following the example.
> [!TIP]
> When finished editing, remove the *.example* extension.
- Optionally create a `.env` file to override the `PORT` environment variable which is defaulted to "8080".
