package main

import (
	"github.com/larwef/cert-monitor/pkg/webapp"
)

func main() {
	webapp.New(GetConf()).Run()
}
