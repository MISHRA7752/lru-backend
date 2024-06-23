# LRU-Backend
Backend code for LRU Cache (go version >= 1.18)
# Install dependencies
-> Run `go mod tidy` to install dependencies.
# Start Project
-> Use `go run main.go` to start the server.
# To change project PORT
-> To change the port, open `main.go`. in case if busy
# Connect Frontend and Backend
-> Ensure the port in `main.go` for the backend matches the port specified in `RuntimeConfig.ts` in the frontend repository.
# Info
-> By default LRU size is set to 10 you can change it in `handler.go`
-> Make sure to start backend server before frontend
