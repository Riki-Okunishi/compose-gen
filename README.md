# docker-compose.yml generator from golang code
This library generate docker-compose.yml from golang code easily

## Installation
```bash
go get -u github.com/Riki-Okunishi/compose-gen
```

## Supported

### Compose Version
+ version 3.5

### Configure options
+ services
  + build
  + image
  + ports
    + Short Syntax
    + Long Syntax
  + volumes
  + depends_on
+ networks
+ volumes

## Usage

### Edit Key-Value
You just call the function with the name of the key you want to add, and give its value as an argument.

```go
import "github.com/Riki-Okunishi/compose-gen/compose"

func main() {
    yml := compose.NewComposeFile("3.5")

    yml.Services("app").Build("./build/path")
    yml.Services("app").Volumes("./volume1/local:./volume1/container", "./volume2/local:./volume2/container")
    yml.Services("app").Ports("10080:80", "10081:81")

    web := yml.Services("web")
    web.Image("nginx:1.18-alpine")
    web.Ports(map[string]interface{}{"target": 80, "published": 8080, "protocol": "tcp", "mode": "host"})

}
```

### Output

**To stdio**

```go
    fmt.Print(yml)
```

**To File**
```go
    yml.GenerateFile("docker-compose.yml")
```

### License
MIT License
