# ErgoVest Backend

Minimal notes to run the backend locally.

1. Copy `.env.example` to `.env` and edit values, or set environment variables.

2. Build & run:

```bash
go mod vendor
go build ./...
./ergovest-backend
```

Or run directly:

```bash
export JWT_SECRET=secret
export MQTT_BROKER=ssl://broker:8883
go run ./main.go
```
