# Ecomm GO API

## Getting Started
To start the server run `go run go run cmd/*.go`
To keep the server running during development install air and run `air`. The .air.toml file is already setup

## Understanding the project
The project is structured to composable. Everything runs as a layer. The cmd file houses the http.
The api.go has the server info and everything is configured and run inside the main.go

The internal folder contains all the info related to services in the app. 
