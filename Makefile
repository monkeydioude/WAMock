watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go run cmd/wamock/main.go examples/single_files_config/config-1.json -x 2'
watch-dir:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/wamock/main.go examples/routes_config/
watch-test:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/wamock/main.go examples/single_files_config/empty.json