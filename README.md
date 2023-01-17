# demoX-rk_worker
demo for worker using rk_boot + redis/kafka

## Installation Instructions

### Requirements
- git
- golang 1.19

### Initial setup
- Install dependencies:
    - Run: `go mod download`

### Run application

- **RUN**:
  - `Initialize`: Only when using kafka - creating initial topic
    - `go run cmd/admin/main.go`
  - `Produce`: Only when using kafka - producing kafka test messages
    - `go run cmd/producer/main.go <message>`
  - `Run Worker`:
    - `go run cmd/worker/main.go`
      - for kafka messge queue, you need to run Initialize first
      - for redis, you can skip both of the upper steps, just run as it is

- **Note**:
    - You may need to run the database and config them first to run the app
    - Check `docker-compose.yaml` and `config.yaml`

### Docker

- Build: `docker build -t rk_worker:latest .`
- To build and run the entire stack (rk_worker, kafka or redis):
    - Run: `docker-compose up -d`

### Testing

- Run: `go test ./... -cover -bench=Benmark`

### Template Structure

- `cmd`: This contains all the run commands for the app
    - `admin/main.go`: to initialize kafka topic (only in kafka mode)
    - `producer/main.go`: to generate message for kafka topic (only in kafka mode)
    -  `worker/main.go`: run the worker instance
- `pkg`: Contains all the pkg, helper modules to build Echo REST API
    - `config`: Contains helper function for app configurations
    - `conn`: Contains helper functions to connect to message queue (kafka, redis)
- `internal`: Contain endpoint and logic for each API. The structure of this module is as follows:
    - `handler`: Handle the all the endpoints of the application (take request, get input and pass it to logic layer)
    - `service`: Logic Layer, all the logic of the service are handled here

