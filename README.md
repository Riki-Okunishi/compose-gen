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
    yml.Services("app").Ports("80:10080", "81:10081")
    yml.Services("app").Ports(compose.PortsShortSyntax("82:10082"))

    web := yml.Services("web")
    web.Image("nginx:1.18-alpine")
    web.Ports(map[string]interface{}{"target": 80, "published": 8080, "protocol": "tcp", "mode": "host"})

}
```

The code in `example/main.go` will generate the following file as `docker-compose.yml`.

```yml
version: "3.5"
services:
  app:
    build: ./build/app
    ports:
      - "10080:80"
    volumes:
      - ./volumes:/mounted
    depends_on:
      - web
  web:
    image: nginx:1.18-alpine
    ports:
      - target: 80
        published: 8080
        protocol: tcp
        mode: host
    volumes:
      - ./volumes:/mounted
      - ./web/nginx/default.conf:/etc/nginx/conf.d/default.conf
    depends_on:
      - db
  db:
    build: ./build/mysql
    volumes:
      - db-store:/var/lib/mysql
volumes:
  db-store:
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
