# fisgo

fisgo is a command line tool that push local files to remote server.this is another advantage of the [fis3](https://github.com/fex-team/fis3), which is push faster and more steady

## Installation
To install `fisgo` use the `go get` command:

```go
go get github.com/JathamJ/fisgo
```

Then you can add `fisgo` binary to PATH environment variable in your ~/.bashrc or ~/.bash_profile file:

>If you already have `fisgo` installed, updating `fisgo` is simple:

```
go get -u github.com/JathamJ/fisgo
```
## Usage
```
cd /path/to/myapp
```
Start fisgo:

```
fisgo
```


fisgo will watch for file events, and every time you create/modify/delete a file it will push the file to remote server.

### Support Options

- -m : Not required, mode of fisgo,support: pusher or server
- -c : Not required, config file of pusher or server, default ./fis-conf.yaml
- -p: Not required, port of receiver server, only server mode is available, default 8899

example:

`fisgo -c ./example/fis-conf.yaml`

### Configuration file

Create a `fis-conf.yaml` file in your project root directory:

```
# fis-conf.yaml configuration example

#local file path of project
root: /Users/xxx/data1/golang/fisgo

#remote receiver api    
receiver: http://127.0.0.1:8899/upload

#remote server root path
to: /Users/xxx/tmp/receiver
```

