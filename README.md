# twilio-autobots

twilio-autobots is a simple repository hosting a client and server application.

These applications are meant to provide a simple CLI on the client side, and communicate with the server in order to return information from github.

# Server

## Build the container
Docker must be installed before building, if docker isn't currently installed it can be installed [here](https://docs.docker.com/get-docker/). After it is installed make sure you are in the base directory of twilio-autobots.
```
docker build -t jesrbegl/twilio-autobots .
```

## Run the server
```
docker run -d -p 8080:8080 --name twilio-autobots-container jesrbegl/twilio-autobots
```

## Deploy to test cluster
Minikube will need to be installed and running, if minikube isn't currently installed it can be installed [here](https://minikube.sigs.k8s.io/docs/start/)

```
kubectl create -f deployment.yaml

kubectl expose deployment twilio-autobots --type=NodePort --name=twilio-autobots-svc --target-port=8080
```

## Test the cluster
Below are a couple commands you can run to make sure that the deployment is working correctly.

### Check that the service is deployed correctly
```
kubectl get services twilio-autobots-svc
```

### Check that the service is working correctly
```
minikube service twilio-autobots-svc 
```

# Client

## Building the client
First you will want to move into the client directory. Afterwords you will build a binary in the standard go process.
```
cd ./client
go build -o autobots.exe
```

## Using the client
After you have build the project, there are a few things you can do, these actions should be documented via the CLI.

### Searching
The below search will print messages to the command line for each origin passed. This will output the number of stars each repository currently has.

Using the origin flag via the CLI directly:
```
./autobots.exe search --origin=jbegley1995/twilio-autobots --origin=google/go-github
```

Passing a file into the CLI:
```
./autobots.exe search --origin-file=..path/to/file/get_repos.txt
```
The file contents should look something like:
```
jbegley1995/twilio-autobots
google/go-github
```

# Testing
Unit tests can be ran through the normal means.

```
go test ./...
```

# Roadmap

## Server
- [x] Dockerized Your code must be able to automatically generate a docker image
- [x] README Your README must include clear instructions on how to run the
- [x] Server as a docker container
- [x] Error Handling Your server must pass meaningful errors on to the client
- [x] Tests Your server must include tests and instructions on how to run your tests
- [x] Calls GitHub Your server must query GitHub’s API for a given list of GitHub repository origins (strings in the form of “organization/repository”) and return the number of stars on those repositories.

## Client
- [x] Runnable You must be able to execute the program from the Command Line
- [x] README Your README must include clear instructions on how to run the client
- [x] CLI Your program must accept arguments
- [x] Input Your client must accept an arbitrary list of GitHub repository origin strings in the form of “organization/repository”
- [x] Input Validation Your program must validate that it has received valid arguments
- [x] Calls Server Your client program must be able to call your dockerized server with the given input, and display the server’s output to the command line.
- [x] Error Handling Your client program must handle any errors that it received from the server
- [x] Tests Your client must include tests and instructions on how to run your tests

## Going the extra mile
- [x] Set up [Minikube](https://minikube.sigs.k8s.io/docs/start/), and deploy the **server** application to your local test cluster.

# TODO
There are several things I wanted to do, but felt like I might be adding a lot of complexity in, and just didn't want to go too far. For instance, the way that the client application handles it's services to connect to the API is very naive. What I should do here is add in a base service, make these urls/ports configurable (if they are not static). The way that it's setup currently would making adding more actions that connect up to the service a bit difficult. We could easily have a base service which would lay the foundation for connecting to the API, and than extend off of that per subdomain.