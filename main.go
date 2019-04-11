package main

import (
	"autossh/core"
	"os"
	"os/user"
)

const template = `
[
  {
    "name": "vagrant",
    "ip": "192.168.33.10",
    "port": 22,
    "user": "root",
    "password": "vagrant",
    "method": "password"
  },
  {
    "name": "ssh-pem",
    "ip": "192.168.33.11",
    "port": 22,
    "user": "root",
    "password": "your pem file password or empty",
    "method": "pem",
    "key": "your pem file path"
  }
]
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	userHome := user.HomeDir
	rootPath := userHome + "/.autossh"
	exist, err := core.PathExists(rootPath)
	if err == nil && !exist {
		err = os.Mkdir(rootPath, os.ModePerm)
		if err == nil {
			config, err := os.OpenFile(rootPath+"/servers.json", os.O_RDWR|os.O_CREATE, 0766)
			if err == nil {
				config.Write([]byte(template))
			}
		}
	}

	app := core.App{ServersPath: rootPath + "/servers.json"}
	app.Exec()
}
