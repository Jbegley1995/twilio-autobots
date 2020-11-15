# twilio-autobots
Description...

# Build the container
```
docker build -t jesrbegl/twilio-autobots .
```

# Run the server
```
docker run -d -p 8080:8080 --name twilio-autobots-container jesrbegl/twilio-autobots
```

# Deploy to test cluster
```
kubectl create -f deployment.yaml
```

# Expose the cluster
```
kubectl expose deployment twilio-autobots --type=NodePort --name=twilio-autobots-svc --target-port=8080
```

# Testing
Unit tests can be ran through the normal means.

```go
go test ./...
```