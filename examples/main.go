package main

import (
	"fmt"

	"github.com/Riki-Okunishi/compose-gen/compose"
)

func main() {
	var yml compose.ComposeFileEditor = compose.NewDefaultComposeFile()

	app := yml.Services("app")
	app.Build("./build/app")
	app.Ports("10080:80")
	app.Volumes("./volumes:/mounted")
	app.DependsOn("web")

	web := yml.Services("web")
	web.Image("nginx:1.18-alpine")
	web.Ports(map[string]interface{}{"target": 80, "published": 8080, "protocol": "tcp", "mode": "host"})
	web.Volumes("./volumes:/mounted", "./web/nginx/default.conf:/etc/nginx/conf.d/default.conf")
	web.DependsOn("db")

	db := yml.Services("db")
	db.Build("./build/mysql")
	db.Volumes("db-store:/var/lib/mysql")

	yml.Volumes("db-store")

	fmt.Print(yml)

	yml.GenerateFile("sample.yml")
}
