package adapter

import (
	"os"

	"github.com/astaxie/beego"
)

var envOs = os.Getenv("GOENV")

// CallAdapter ...
func CallAdapter(call string) string {
	var str string
	assignAdapter(&str, call)
	return str
}

func assignAdapter(str *string, call string) {
	if envOs == "local" {
		*str = beego.AppConfig.String(call + "::local")
	}
}
