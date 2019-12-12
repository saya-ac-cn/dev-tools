package entity

// 数据库实体类
type DbAttributeEntity struct {
	Id         int    `json:"Id" xorm:"id"`
	DBName     string `json:"DBName" xorm:"db_name"`
	OwnerName  string `json:"OwnerName" xorm:"owner_name"`
	CreateTime string `json:"CreateTime" xorm:"create_time"`
	UpdateTime string `json:"UpdateTime" xorm:"update_time"`
}
