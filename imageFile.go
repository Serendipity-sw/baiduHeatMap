package main

import (
	"encoding/base64"
	"fmt"
	"github.com/smtc/glog"
	"io"
	"os"
	"strings"
	"time"
)

/**
将图片流字符串解析保存到给定的目录中
PS:图片流格式:data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGIAAABiCAYAAACrpQYOAAAAGX.........
创建人:邵炜
创建时间:2016年3月8日13:44:19
输入参数: imageByte 图片二进制流字符串,  path需要保存到的目录
输出参数: 图片目录   错误对象
*/
func imageFileSave(imageByte *string, path string) (string, error) {

	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	path = path + "/%d.%s"

	images := strings.Split(*imageByte, ",")

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(images[1]))

	imageName := strings.Split(strings.Split(images[0], ";")[0], "/")[1]

	fileName := fmt.Sprintf(path, time.Now().Unix(), imageName)

	file, err := os.Create(fileName)

	if err != nil {
		glog.Error("image create is error, err: %s \n", err.Error())
		return "", err
	}

	_, err = io.Copy(file, reader)

	defer file.Close()

	if err != nil {
		glog.Error("image copyu is error, err: %s \n", err.Error())
		return "", err
	}

	return fileName, nil
}
