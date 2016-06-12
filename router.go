package main

import (
	"github.com/gin-gonic/gin"
	"github.com/smtc/glog"
	"net/http"
	"path"
	"strings"
)

func assetsFiles(c *gin.Context) {
	r := c.Request
	pth := c.Param("pth")
	if pth == "" {
		glog.Error("assetsFiles: path is empty: %s\n", r.URL.Path)
		c.Data(200, "text/plain", []byte(""))
		return
	}

	fp, err := getAssetFilePath(pth)
	if err != nil {
		glog.Error("assetsFiles: %s\n", err)
		c.Data(200, "text/plain", []byte(""))
		return
	}

	http.ServeFile(c.Writer, c.Request, fp)
}

func getAssetFilePath(pth string) (string, error) {
	entrys := strings.Split(pth, "/")
	sentrys := []string{contentDir}
	for _, s := range entrys {
		s = strings.TrimSpace(s)
		if s != "" {
			sentrys = append(sentrys, s)
		}
	}
	return path.Join(sentrys...), nil
}
