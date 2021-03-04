package main

import (
	"fmt"
	"github.com/Riki-Okunishi/compose-gen/compose"
)

func main() {
	var yml compose.ComposeFileEditor = compose.NewDefaultComposeFile()
	
	yml.Services("app").Build("./build/app")
	yml.Services("app").Volumes("./volumes:/mounted")
	yml.Services("app").DependsOn("web")
	
	yml.Services("web").Image("nginx:1.18-alpine")
	yml.Services("web").Volumes("./volumes:/mounted", "./web/nginx/default.conf:/etc/nginx/conf.d/default.conf")
	yml.Services("web").DependsOn("db")

	yml.Services("db").Build("./build/mysql")
	yml.Services("db").Volumes("db-store:/var/lib/mysql")

	yml.Volumes("db-store")

	fmt.Print(yml)
}