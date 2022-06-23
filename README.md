# ThePorter
A Golang based PortDomainService service that either creates a new record in a database, or updates the existing one.

# Steps to use this binary
1. checkout the repository.
2. make sure golang 1.17 is installed on your system.
3. go the directory and run `go build main.go`
4. run the compiled binary following way : ./main --config ~/Documents/theporter/config.yml --portdata ~/tmp/ports.json

# Other way of running this binary 
1. checkout the repository
2. compile the go-binary
3. run dockerbuild command using the -f option, provide the porter.Dockerfile as input
4. start a redis-container on the same machin. update the config.yml file and provide its credentials as input.
5. docker-exec the container to run the application.

# Todo
- add docker compose file for automatic container orchestration.
- add makefile to automate the binary building process.
- daemonize the binary, so that it keeps waiting at a path for file, as soon as file is uploaded, it will update the database.
-

