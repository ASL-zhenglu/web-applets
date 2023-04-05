package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goPC/dao"
	"goPC/model"
	"net/http"
	"sort"
	"strconv"
)

func PCLogin(c *gin.Context) {
	//使用map来获取请求的参数
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)

	//获取参数
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	username := requestMap["username"]
	password := requestMap["password"]
	fmt.Println(username)
	fmt.Println(password)

	//var r []model.Admin
	//dao.DB.Find(&r, "name=? AND password = ?", username, password)
	//dao.DB.Create(&model.Admin{Name: "666", Password: "666"})
	var r []model.Admin
	dao.DB.Find(&r, "name=? AND password = ?", username, password)
	//fmt.Printf("%#v", result)
	//fmt.Printf("%#v\n", result)
	//fmt.Println("???????????????????????")

	//err := db.Where("name = ? AND password = ?", username, password).Find(&Admin{}).Error
	//fmt.Printf("%#v", r)
	//fmt.Println("???????????????????")
	if len(r) == 0 {
		fmt.Println("查询不到数据")
	} else {
		fmt.Println("查找成功")
		c.JSON(http.StatusOK, gin.H{
			"msg": "1",
		})
		c.Redirect(http.StatusMovedPermanently, "CanteenPC/home")
	}
}

// Reserve 预约申请界面
func Reserve(c *gin.Context) {
	var r []model.Reserve
	dao.DB.Find(&r, "status=?", 1)
	c.JSON(http.StatusOK, r)
}

//AgreeReserve 同意预约
func AgreeReserve(c *gin.Context) {
	var requestMap = make(map[string]int)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	id := requestMap["id"]
	dao.DB.Model(&model.Reserve{}).Where("id = ?", id).Update("status", 2)
}

//DisagreeReserve 拒绝预约
func DisagreeReserve(c *gin.Context) {
	var requestMap = make(map[string]int)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	id := requestMap["id"]
	dao.DB.Model(&model.Reserve{}).Where("id = ?", id).Update("status", 0)
}

// AdminApply 管理员申请界面
func AdminApply(c *gin.Context) {
	var ad []model.User
	dao.DB.Find(&ad, "status=?", 1)
	c.JSON(http.StatusOK, ad)
}

// AgreeApply 同意成为管理员
func AgreeApply(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	var wx = requestMap["WxId"]
	fmt.Println(wx)
	dao.DB.Model(&model.User{}).Where("wx_id = ?", wx).Update("status", 2)
	var user = model.User{WxId: wx}
	dao.DB.First(&user)
	dao.DB.Create(&model.Admin{Name: wx, Password: user.Password})
}

// DisagreeApply 不同意申请
func DisagreeApply(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	var wx = requestMap["WxId"]
	dao.DB.Model(&model.User{}).Where("wx_id = ?", wx).Update("status", 0)
}

// Opinion 显示意见
func Opinion(c *gin.Context) {
	var advise []model.Advise
	dao.DB.Find(&advise)
	//fmt.Printf("%#v", advise)
	c.JSON(http.StatusOK, advise)
}

// DeleteOpinion 删除意见
func DeleteOpinion(c *gin.Context) {
	var requestMap = make(map[string]int)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	var id = requestMap["Id"]
	//fmt.Println(id)
	//fmt.Println("????????????????????????")
	dao.DB.Where("id = ?", id).Delete(model.Advise{})
}

// MeetingRoom 显示会议室
func MeetingRoom(c *gin.Context) {
	var r []model.Room
	dao.DB.Find(&r)
	c.JSON(http.StatusOK, r)
}

// DeleteRoom 删除会议室
func DeleteRoom(c *gin.Context) {
	var requestMap = make(map[string]int)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	var id = requestMap["id"]
	fmt.Println(id)
	fmt.Println("?????????????????")
	dao.DB.Where("id = ?", id).Delete(model.Room{})
}

// EditRoom 修改会议室
func EditRoom(c *gin.Context) {
	var requestMap = make(map[string]int)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	id := requestMap["id"]
	value := requestMap["value"]
	fmt.Println("+", id, "+", value, "+")
	dao.DB.Model(&model.Room{}).Where("id = ?", id).Update("size", value)
}

// AddRoom 添加会议室
func AddRoom(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	name := requestMap["name"]
	s := requestMap["size"]
	size, _ := strconv.Atoi(s)
	remark := requestMap["remark"]
	var r []model.Room
	dao.DB.Find(&r, "room_name=?", name)

	if len(r) > 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "0"})
	} else {
		room := model.Room{RoomName: name, Size: size, Remark: remark}
		dao.DB.Create(&room)
		c.JSON(http.StatusOK, gin.H{"msg": "1"})
	}
}

//AppointmentRecord 预约历史记录
func AppointmentRecord(c *gin.Context) {
	var r []model.Reserve
	//dao.DB.Find(&r, "status IN ?", []int{0, 2})
	dao.DB.Where("status = ?", 0).Or("status=?", 2).Find(&r)
	fmt.Printf("%#v", r)
	c.JSON(http.StatusOK, gin.H{
		"msg":    r,
		"length": len(r),
	})
}

//RecordSearch 搜索历史预约记录
func RecordSearch(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	name := requestMap["name"]
	var r []model.Reserve
	dao.DB.Where("status = ? AND room_name=?", 0, name).Or("status = ? AND room_name=?", 2, name).Find(&r)
	dao.DB.Where("status = ? AND room_name=?", 0, name).Or("status = ? AND room_name=?", 2, name).Find(&r)
	c.JSON(http.StatusOK, gin.H{
		"msg":    r,
		"length": len(r),
	})
}

// RoomAnalysis 预约室分析
func RoomAnalysis(c *gin.Context) {
	var room []model.Reserve
	dao.DB.Where("status = ?", 0).Or("status = ?", 2).Find(&room)
	mp := map[string]int{}
	for _, x := range room {
		mp[x.RoomName]++
	}
	name := make([]string, 0)
	cnt := make([]int, len(mp))
	for k := range mp {
		name = append(name, k)
	}
	sort.Slice(name, func(i, j int) bool {
		a, b := name[i], name[j]
		return mp[a] < mp[b]
	})
	for i, x := range name {
		cnt[i] = mp[x]
	}
	//fmt.Println(name, cnt)
	//fmt.Println("??????????????????????????")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"cnt":  cnt,
	})
}

// PeriodAnalysis 时间段分析
func PeriodAnalysis(c *gin.Context) {
	var t []model.Reserve
	dao.DB.Where("status = ?", 0).Or("status = ?", 2).Find(&t)
	mp := map[string]int{}

	for _, x := range t {
		mp[x.Period]++
	}
	p := []string{"8:00-10:00", "10:15-12:15", "14:00-16:00", "16:15-18:15", "20:00-22:00"}
	cnt := make([]int, 5)
	for i := 0; i < 5; i++ {
		cnt[i] = mp[p[i]]
	}
	//fmt.Println(p, cnt)
	//fmt.Println("??????????????????????????")
	c.JSON(http.StatusOK, gin.H{
		"name": p,
		"cnt":  cnt,
	})
}
