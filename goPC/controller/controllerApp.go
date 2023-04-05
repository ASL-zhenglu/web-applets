package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goPC/dao"
	"goPC/model"
	"net/http"
	"time"
)

// AddAdvice 提意见
func AddAdvice(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	//time := requestMap["time"]
	advise := requestMap["advise"]
	//fmt.Printf("%#v", time)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%#v", advise)
	//time := c.PostForm("time") //对应前端的name--username
	//advise := c.PostForm("advise")
	ad := model.Advise{Time: timeStr, Advise: advise}
	dao.DB.Create(&ad)
	//data := []int{1, 2, 3}
	//header := []int{2, 3, 4}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

type rm struct {
	Xx string
	Yy int
	Zz string
}

// GetRoom 显示小程序界面
func GetRoom(c *gin.Context) {

	var r []model.Room
	dao.DB.Find(&r)
	res := make([]rm, len(r))
	for i, v := range r {
		res[i].Xx = v.RoomName
		res[i].Yy = v.Size
		res[i].Zz = v.Remark
	}
	//fmt.Printf("%#v", res)
	//dao.DB.Select("room_name,size,remark").Find(&r)
	c.JSON(http.StatusOK, res)
}

// ApplyReserve 申请预约
func ApplyReserve(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	date := requestMap["date"]
	period := requestMap["period"]
	rName := requestMap["room_name"]
	reason := requestMap["reason"]
	name := requestMap["name"]
	phone := requestMap["phone"]
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//var res []model.Reserve
	//dao.DB.Find(&res,"date=? AND period = ? AND room_name=? AND wx_id=? status = ?",date,period,rName,)
	r := model.Reserve{ApplyTime: timeStr, Date: date, Period: period, RoomName: rName, WxId: name, Reason: reason, Phone: phone, Status: 1}
	dao.DB.Create(&r)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// GetUser 显示个人信息
func GetUser(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	wx := requestMap["name"]
	var r []model.User
	dao.DB.Find(&r, "wx_id = ?", wx)
	if len(r) == 0 {
		user := model.User{WxId: wx}
		dao.DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"msg": "create",
		})
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// UpdateUser 更新个人信息
func UpdateUser(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	wx := requestMap["name"]
	username := requestMap["username"]
	address := requestMap["address"]
	phone := requestMap["phone"]
	var u []model.User
	dao.DB.Find(&u, "wx_id = ?", wx)
	if len(u) == 0 {
		user := model.User{WxId: wx}
		dao.DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"msg": "create",
		})
	}
	dao.DB.Model(&u).Where("wx_id = ?", wx).Updates(model.User{Username: username, Address: address, Phone: phone})
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	//fmt.Printf("%#v", u)

}

//Apply 申请成为管理员
func Apply(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	wx := requestMap["name"]
	reason := requestMap["reason"]
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	var u model.User
	dao.DB.First("wx_id=?", wx)
	dao.DB.Model(&u).Where("wx_id=?", wx).Updates(model.User{ApplyTime: timeStr, Status: 1, Reason: reason})
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

//GetRecord 获得历史记录
func GetRecord(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	wx := requestMap["name"]
	var u []model.Reserve
	dao.DB.Find(&u, "wx_id=?", wx)
	r := make([]model.Res, len(u))
	for i, v := range u {
		r[i].RoomName = v.RoomName
		if v.Status == 0 {
			r[i].Status = "未通过"
		} else if v.Status == 1 {
			r[i].Status = "审核中"
		} else {
			r[i].Status = "通过"
		}
		r[i].Date = v.Date
		r[i].Period = v.Period
	}
	c.JSON(http.StatusOK, r)
}
