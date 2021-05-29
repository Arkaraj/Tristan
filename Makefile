dev:
	go run src/main.go

# best script!!
start: 
	nodemon --exec go run src/main.go --signal SIGTERM