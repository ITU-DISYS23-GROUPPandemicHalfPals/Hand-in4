# Distributed Mutual Exclusion
## Running the application
### Node
- Change directory to `Hand-in4/node`.
- Use `go run . -port <port>`, this will start a node with given port. (Only one of each port between 5000-5002 work, since we have hardcoded the port values)
- The nodes will now find a coordinator and begin requesting access to the critical section.
- It is possible to add and remove clients while the program is running.