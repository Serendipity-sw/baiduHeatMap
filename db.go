package main

import (
	"encoding/json"
	"fmt"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/smtc/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

/**
热力图显示对象 region_heatmap_hist
创建人:邵炜
创建时间:2016年6月16日14:07:10
*/
type regionHeatMapHist struct {
	Id             string `bson:"_id"`             //mongoDB自动生成的唯一标识
	Count          int32  `bson:"Count"`           //人数
	Archive_date   string `bason:"archive_date"`   //数据日期
	Update_time    string `bason:"update_time"`    //更新时间
	Region_id      string `bason:"region_id"`      //区域ID
	Archive_minute int32  `bason:"archive_minute"` //数据日期精确到分钟
	Archive_second int32  `bason:"archive_second"` //数据日期精确到秒
	Archive_hour   int32  `bason:"archive_hour"`   //数据日期精确到小时
	Lat            string `bason:"lat"`            //经度
	Lng            string `bason:"lng"`            //纬度
}

/**
区域分析数据对象 region_stream_raw
创建人:邵炜
创建时间:2016年6月16日14:29:36
*/
type regionStreamRaw struct {
	Id        string `bson:"_id"`       //唯一标示
	Region_id string `bson:"region_id"` //区域ID
	User_id   string `bson:"user_id"`   //用户号码
	Lng       string `bson:"lng"`       //纬度
	Lat       string `bson:"lat"`       //经度
	Age       string `bson:"age"`       //年龄
	Gender    string `bson:"gender"`    //性别
	Locale    string `bson:"locale"`    //归属地区
	Arpu      string `bson:"arpu"`      //月平均消费
}

/**
构造函数
创建人:邵炜
创建时间:2016年6月16日14:02:07
*/
func init() {
	deferinit.AddInit(mongoDBConnection, mongoDBClose, 999)
}

/**
mongoDB数据库连接打开
创建人:邵炜
创建时间:2016年6月16日13:58:56
*/
func mongoDBConnection() {
	var err error
	dbHost := config.GetString("dbhost")
	dbPort, bo := config.GetInt("dbport")
	if !bo {
		glog.Error("mongonDBConntion get my dbport can't obtain! dbport: %d \n", dbPort)
		return
	}
	connectionStr := fmt.Sprintf("%s:%d", dbHost, dbPort)
	session, err = mgo.Dial(connectionStr)
	if err != nil {
		glog.Error("mongonDBConntion DB conntion error. conntionStr: %s  err: %s \n", connectionStr, err.Error())
		return
	}
	session.SetMode(mgo.Monotonic, true)
}

/**
mongoDB数据关闭
创建人:邵炜
创建时间:2016年6月16日14:01:48
*/
func mongoDBClose() {
	session.Close()
}

/**
查询区域分析集合数据
创建人:邵炜
创建时间:2016年6月16日14:33:18
*/
func selectRegionStreamRaw() (*[]regionStreamRaw, error) {
	var regionStreamRawList []regionStreamRaw
	c := session.DB("mwc").C("region_stream_raw")
	err := c.Find(bson.M{"region_id": "BLZX"}).All(&regionStreamRawList)
	if err != nil {
		glog.Error("selectRegionStreamRaw select data is error! err: %s \n", err.Error())
		return nil, err
	}
	dataJsonStr, err := json.Marshal(regionStreamRawList)
	if err != nil {
		glog.Error("selectRegionStreamRaw data can't marshal! err: %s \n", err.Error())
		return nil, err
	}
	go glog.Info("selectRegionStreamRaw select success! dataJsonStr: %s \n", string(dataJsonStr))
	return &regionStreamRawList, nil
}

/**
查询热力图显示区域数据
创建人:邵炜
创建时间:2016年6月16日14:33:18
*/
func selectRegionHeatMapHist() (*[]regionHeatMapHist, error) {
	var regionHeatMapHistList []regionHeatMapHist
	c := session.DB("mwc").C("region_heatmap_hist")
	err := c.Find(bson.M{"region_id": "BLZX"}).All(&regionHeatMapHistList)
	if err != nil {
		glog.Error("selectRegionHeatMapHist select data is error! err: %s \n", err.Error())
		return nil, err
	}
	dataJsonStr, err := json.Marshal(regionHeatMapHistList)
	if err != nil {
		glog.Error("selectRegionHeatMapHist data can't marshal! err: %s \n", err.Error())
		return nil, err
	}
	go glog.Info("selectRegionHeatMapHist select success! dataJsonStr: %s \n", string(dataJsonStr))
	return &regionHeatMapHistList, nil
}
