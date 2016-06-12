package main

import "syscall"

// +build linux
const recordPath = "/tmp/import/"
const pidFile = "./sceneportal.pid" // /var/run 防止程序没有权限创建该文件

func isProcessExist(pid int) bool {
	// 进程是否存在
	err := syscall.Kill(pid, 0)
	if err == nil {
		return true
	}
	return false
}
