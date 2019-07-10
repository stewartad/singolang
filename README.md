[![Go Report Card](https://goreportcard.com/badge/github.com/stewartad/singolang)](https://goreportcard.com/report/github.com/stewartad/singolang)

# Singolang
Singolang is a library to interact with Singularity containers in Go. It is modeled from Spython. Designed for use with Singularity 3+

## Currently Supported Features
* Pulling images from Dockerhub or Singularity Hub
* Starting and Stopping instances of built images
* Executing commands in running instances

## Usage
### Create a Client
To start using Singolang, create a new client

```go
client, teardown := client.NewClient()
defer teardown()
```

`NewClient()` returns both a Client struct and a teardown function. The teardown function stops all instances and is intended to be deferred

### Pull an image

```go
pullOpts := &client.PullOptions{
    Name: "",
    Pullfolder: filepath.Join("/tmp", "pull"),
    Force: false,
}

imgPath, err := cl.Pull("docker://godlovedc/lolcow", pullOpts)

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(imgPath)
}
```