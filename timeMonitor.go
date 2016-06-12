package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/howeyc/fsnotify"
	"github.com/smtc/glog"
	"html/template"
	"strings"
	"sync"
	"time"
)

var (
	jsTmr    *time.Timer
	funcName = template.FuncMap{
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"safeurl": func(s string) template.URL {
			return template.URL(s)
		},
	}
)

func init() {
	deferinit.AddInit(func() {
		tempDir = config.GetStringDefault("tempDir", "template/")
		if !strings.HasSuffix(tempDir, "/") {
			tempDir += "/"
		}
	}, nil, 40)
	deferinit.AddRoutine(notifyTemplates)
	deferinit.AddRoutine(watchFuncDir)
}

/**
定时运行程序
创建人:邵炜
创建时间:2016年3月7日09:51:42
输入参数: 终止命令  计数器对象
*/
func watchFuncDir(ch chan struct{}, wg *sync.WaitGroup) {
	go func() {
		<-ch

		jsTmr.Stop()
		wg.Done()
	}()

	jsTmr = time.NewTimer(time.Minute)
	for {
		//需要定时执行的方法
		jsTmr.Reset(time.Minute)
		<-jsTmr.C
	}
}

/**
加载模版
创建人:邵炜
创建时间:2016年2月26日11:34:12
输入参数: gin对象
*/
func loadTemplates(e *gin.Engine) {
	t, err := template.New("tmpls").Funcs(funcName).ParseGlob(tempDir + "*")

	if err != nil {
		glog.Error("loadTemplates failed: %s %s \n", tempDir, err.Error())
		return
	}

	e.SetHTMLTemplate(t)
}

/**
监视文件夹目录如发生任何修改,重新载入
创建人:邵炜
创建时间:2016年3月7日09:47:50
输入参数: 终止命令 计数器对象
*/
func notifyTemplates(ch chan struct{}, wg *sync.WaitGroup) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		glog.Error("notifyTemplates: create new watcher failed: %v\n", err)
		return
	}

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				glog.Debug("notifyTemplates: event: %v\n", ev)
				loadTemplates(rt)
			case err := <-watcher.Error:
				glog.Error("notifyTemplates: error: %v\n", err)
			}
		}
	}()

	err = watcher.Watch(tempDir)
	if err != nil {
		glog.Error("notifyTemplates: watch dir %s failed: %v \n", tempDir, err)
	}

	// Hang so program doesn't exit
	<-ch

	/* ... do clean stuff ... */
	watcher.Close()
	wg.Done()
}
