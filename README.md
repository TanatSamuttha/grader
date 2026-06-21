# Grader

A distributed competitive programming online judge. This online judge is designed using a microservices architecture.

## Feature
- Multiple worker-based grading system
- OAuth-based authentication
- Docker-based sandboxed code execution
- Real-time grading process via Websocket
- Multiple test cases evaluation
- Runtime and memory limit error detection

## Services
### Web-service
- Serve frontend staatic files
- Provide reverse-proxy

### Auth-service
