# twilio-autobots

twilio-autobots is a simple repository hosting a client and server application.

These applications are meant to provide a simple CLI on the client side, and communicate with the server in order to return information from github.

# Server

## Build the container
```
docker build -t jesrbegl/twilio-autobots .
```

## Run the server
```
docker run -d -p 8080:8080 --name twilio-autobots-container jesrbegl/twilio-autobots
```

## Deploy to test cluster
```
kubectl create -f deployment.yaml
```

## Expose the cluster
```
kubectl expose deployment twilio-autobots --type=NodePort --name=twilio-autobots-svc --target-port=8080
```

# Client

## Building the client
First you will want to move into the client directory. Afterwords you will build a binary in the standard go process.
```
go build -o autobots.exe
```

## Using the client
After you have build the project, there are a few things you can do, these actions should be documented via the CLI.

### Searching
The below search will print messages to the command line for each origin passed. This will output the number of stars each repository currently has.

```
./autobots.exe search --origin=jbegley1995/twilio-autobots --origin=github/github-go
```

# Testing
Unit tests can be ran through the normal means.

```
go test ./...
```