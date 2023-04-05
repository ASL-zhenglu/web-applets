package model

// Admin 管理员
type Admin struct {
	Name     string
	Password string
}

// User 用户
type User struct {
	UserId    string `gorm:"autoIncrement"`
	WxId      string `gorm:"not null; unique"`
	ApplyTime string `form:"time" json:"time"`
	Username  string
	Password  string
	Address   string
	Phone     string
	Reason    string
	Status    int //0代表普通用户 1代表正在申请用户 2代表管理员
}

// Advise 建议
type Advise struct {
	ID     uint `gorm:"primary_key"`
	Time   string
	Advise string
}

// Room 会议室
type Room struct {
	ID       uint `gorm:"primary_key"`
	RoomName string
	Size     int
	Remark   string
}

//Reserve 预约表
type Reserve struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	Date      string
	Period    string
	RoomName  string
	WxId      string
	Reason    string
	Status    int    //0代表未占用 1表示正在审核 2表示占用
	ApplyTime string //申请时间
	Phone     string
}

// Empty 空字段
type Empty struct {
}

type Res struct {
	RoomName string
	Status   string
	Date     string
	Period   string
}
