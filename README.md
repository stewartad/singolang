# Singolang

[![Go Report Card](https://goreportcard.com/badge/github.com/stewartad/singolang)](https://goreportcard.com/report/github.com/stewartad/singolang)
Singolang is a library to interact with Singularity containers in Go. It is modeled from Spython. Designed for use with Singularity 3+

## Currently Supported Features

* Pulling images from Dockerhub or Singularity Hub
* Starting and Stopping instances of built images
* Executing commands in running instances

## Usage

### Create a Client

To start using Singolang, create a new client

```go
client, teardown := singolang.NewClient()
defer teardown()
```

`NewClient()` returns both a Client struct and a teardown function. The teardown function stops all instances and is intended to be deferred

### Pull an image

Pulling an image requires a struct to be filled out detailing the options with which to perform the pull.

Name is the filename to save the image as.

Pullfolder is the folder to place the final image

Force, if true, will overwrite any existing files of the same name

```go
opts := &client.PullOptions{
    Name: "",
    Pullfolder: filepath.Join("/tmp", "pull"),
    Force: false,
}

imgPath, err := client.Pull("docker://godlovedc/lolcow", pullOpts)

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(imgPath)
}
```

### Create an Instance

```go
err := client.NewInstance("lolcow_latest.sif", "lolcow3")
if err != nil {
    fmt.Println(err)
}
```

### Execute a Command

```go
opts := singolang.DefaultExecOptions()
stdout, stderr, code, err := client.Execute("lolcow3", []string{"which", "fortune"}, opts)
fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)
```
