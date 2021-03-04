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
    yml := compose.NewComposeFile("3.9")

    yml.Service("app").Build("./build/path")
    yml.Service("app").Volumes("./volume1/local:./volume1/container", "./volume2/local:./volume2/container")
}
```

### Output as file
coming soon...

### License
MIT License
