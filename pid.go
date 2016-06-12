package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// 将进程pid写入文件中
func writePid() {
	pid := os.Getpid()
	f, err := os.OpenFile(pidFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		panic(fmt.Sprintf("Open pid file %s failed: %s\n", pidFile, err.Error()))
	}

	_, err = f.Write([]byte(fmt.Sprint(pid)))
	if err != nil {
		panic(fmt.Sprintf("Write Pid file %s failed: %s\n", pidFile, err.Error()))
	}
	f.Close()
}

// 检查pid文件是否存在，pid文件中的进程是否存在
// 返回：true： 文件不存在或者进程不存在
//      false: 进程已存在
func checkPid() bool {
	f, err := os.Open(pidFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(fmt.Sprintf("Open Pid file %s failed: %s\n", pidFile, err.Error()))
		}
	}
	defer f.Close()
	// 读取文件内容
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("Read Pid file %s failed: %s\n", pidFile, err.Error()))
	}
	pid, err := strconv.Atoi(string(buf))
	if err != nil {
		panic(fmt.Sprintf("Convert pid %s failed: %s\n", pid, err.Error()))
	}
	// 进程是否存在
	exist := isProcessExist(pid)
	if exist == false {
		return false
	}
	fmt.Printf("Process with Pid %d is running, exit.\n", pid)
	return true
}

// 删除pid文件
func rmPidFile() {
	os.Remove(pidFile)
}
