package main

// +build windows
import "github.com/smtc/glog"

func logInit(debug bool) {
	var option = make(map[string]interface{})

	option["typ"] = "file"
	if debug {
		glog.InitLogger(glog.DEV, option)
	} else {
		glog.InitLogger(glog.PRO, option)
	}
}
