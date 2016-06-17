package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smtc/glog"
	"net/http"
	"path"
	"strconv"
	"strings"
)

//查询类接口返回参数
type discntInfo struct {
	ResultCode string        //返回码
	Message    string        //返回消息信息
	Data       []interface{} //数据载体
}

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

/**
查询区域分析集合数据路由方法
创建人:邵炜
创建时间:2016年6月17日10:49:28
输入参数:gin 对象
*/
func selectRegionStreamRawRouter(c *gin.Context) {
	userHttpInfo := userReqInfo(c.Request)
	jsonDataIn, err := selectRegionStreamRaw()
	if err != nil {
		glog.Error("selectRegionStreamRawRouter selectRegionStreamRaw run error! userHttpInfo: %s err: %s \n", userHttpInfo, err.Error())
		jsonPRequest(c, true, "数据请求发生错误")
		return
	}
	glog.Info("selectRegionStreamRawRouter run success! userHttpInfo: %s \n", userHttpInfo)
	jsonPRequest(c, false, *jsonDataIn)
}

/**
查询热力图显示区域数据路由方法
创建人:邵炜
创建时间:2016年6月17日10:51:24
输入参数:gin 对象
*/
func selectRegionHeatMapHistRouter(c *gin.Context) {
	userHttpInfo := userReqInfo(c.Request)
	jsonDataIn, err := selectRegionHeatMapHist()
	if err != nil {
		glog.Error("selectRegionHeatMapHistRouter selectRegionHeatMapHist run error! userHttpInfo: %s err: %s \n", userHttpInfo, err.Error())
		jsonPRequest(c, true, "数据请求发生错误")
		return
	}
	glog.Info("selectRegionHeatMapHistRouter run success! userHttpInfo: %s \n", userHttpInfo)
	jsonPRequest(c, false, *jsonDataIn)
}

/**
查询区域分析集合数据及热力图显示区域数据路由方法
创建人:邵炜
创建时间:2016年6月17日13:41:55
输入参数: gin 对象
*/
func selectMapHistAndStreamRawRouter(c *gin.Context) {
	userHttpInfo := userReqInfo(c.Request)
	jsonDataStreamRawIn, err := selectRegionStreamRaw()
	if err != nil {
		glog.Error("selectMapHistAndStreamRawRouter selectRegionStreamRaw run error! userHttpInfo: %s err: %s \n", userHttpInfo, err.Error())
		jsonPRequest(c, true, "区域数据请求发生错误")
		return
	}
	jsonDataHeatMapHistIn, err := selectRegionHeatMapHist()
	if err != nil {
		glog.Error("selectRegionHeatMapHistRouter selectRegionHeatMapHist run error! userHttpInfo: %s err: %s \n", userHttpInfo, err.Error())
		jsonPRequest(c, true, "热力图数据请求发生错误")
		return
	}
	/**
	数据反馈对象构建
	创建人:邵炜
	创建时间:2016年6月17日13:44:53
	*/
	type resultHeatMapStruct struct {
		StreamRaw   string `json:"streamRaw"`   //区域分析集合json字符串
		HeatMapHist string `json:"heatMapHist"` //热力图区域数据集合json字符串
	}
	resultHeatMapStructClass := resultHeatMapStruct{
		StreamRaw:   *jsonDataStreamRawIn,
		HeatMapHist: *jsonDataHeatMapHistIn,
	}
	resultHeatMapStructByte, err := json.Marshal(resultHeatMapStructClass)
	if err != nil {
		glog.Error("selectRegionHeatMapHistRouter jsonData can't marshal! userHttpInfo: %s err: %s \n", userHttpInfo, err.Error())
		jsonPRequest(c, true, "返回对象json字符串序列化失败")
		return
	}
	resultHeatMapStructStr := string(resultHeatMapStructByte)
	glog.Info("selectMapHistAndStreamRawRouter run success! userHttpInfo: %s \n", string(resultHeatMapStructStr))
	jsonPRequest(c, false, resultHeatMapStructStr)
}

/**
首页路由方法
创建人:邵炜
创建时间:2016年6月17日13:59:38
*/
func index(c *gin.Context) {
	glog.Info("index http success! userHttpInfo: %s \n")
	c.HTML(http.StatusOK, "index.html", nil)
}

/**
JSON请求数据返回
创建人:邵炜
创建时间:2016年1月4日20:23:36
输入参数: gin指针 bo判断是否使用错误返回对象 param泛型参数
输出参数: 无
数据反馈由gin进行
*/
func jsonPRequest(c *gin.Context, bo bool, param interface{}) {

	var cb string

	if c.Request.Method == "GET" {
		cb = c.Query("callback")
	} else {
		cb = c.PostForm("callback")
	}

	jsonResP := &discntInfo{
		ResultCode: "00000",
		Message:    "",
	}

	switch paramType := param.(type) {
	case string:
		if bo {
			jsonResP.ResultCode = "00001"
		}
		jsonResP.Message = param.(string)
		if cb != "" {
			b, _ := json.Marshal(jsonResP)
			c.Data(http.StatusOK, "application/javascript", []byte(fmt.Sprintf("%s(%s)", cb, b)))
		} else {
			c.JSON(http.StatusOK, jsonResP)
		}
	case int32:
		jsonResP.Message = strconv.Itoa(int(paramType))
		if cb != "" {
			b, _ := json.Marshal(jsonResP)
			c.Data(http.StatusOK, "application/javascript", []byte(fmt.Sprintf("%s(%s)", cb, b)))
		} else {
			c.JSON(http.StatusOK, jsonResP)
		}
	case int64:
		jsonResP.Message = strconv.Itoa(int(paramType))
		if cb != "" {
			b, _ := json.Marshal(jsonResP)
			c.Data(http.StatusOK, "application/javascript", []byte(fmt.Sprintf("%s(%s)", cb, b)))
		} else {
			c.JSON(http.StatusOK, jsonResP)
		}
	default:
		if cb != "" {
			b, _ := json.Marshal(paramType)
			c.Data(http.StatusOK, "application/javascript", []byte(fmt.Sprintf("%s(%s)", cb, b)))
		} else {
			c.JSON(http.StatusOK, param)
		}
	}
}

/**
报告用户请求相关信息
创建人:邵炜
创建时间:2016年6月17日10:57:24
输入参数: request对象
输出参数: 用户相关信息字符串
*/
func userReqInfo(req *http.Request) (info string) {
	info += fmt.Sprintf("ipaddr: %s user-agent: %s referer: %s",
		req.RemoteAddr, req.UserAgent(), req.Referer())
	return info
}
