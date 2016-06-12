package main

import "syscall"

// +build darwin
const recordPath = "./logs"
const pidFile = "./sceneportal.pid"

func isProcessExist(pid int) bool {
	// 进程是否存在
	err := syscall.Kill(pid, 0)
	if err == nil {
		return true
	}
	return false
}
