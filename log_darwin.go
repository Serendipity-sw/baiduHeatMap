package main

// +build darwin
import "github.com/smtc/glog"

func logInit(debug bool) {
	var option = make(map[string]interface{})

	option["typ"] = "file"

	glog.InitLogger(glog.DEV, option)
}
