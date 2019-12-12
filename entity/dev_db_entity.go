package entity

// 开发数据库变更实体类
type DevDbEntity struct {
	Id         int    `json:"Id" xorm:"id"`              // 数据库变更编号(开发)
	DBName     string `json:"DBName" xorm:"db_name"`     // 变更库
	DBItem     string `json:"DBItem" xorm:"db_item"`     // 变更内容
	DBReason   string `json:"DBReason" xorm:"db_reason"` // 变更缘由
	CreateTime string `json:"CreateTime" xorm:"create_time"`
	Status     int8   `json:"Status" xorm:"status"`      // 发布状态
	UserName   string `json:"UserName" xorm:"user_name"` // 变更者
	UserId     int    `json:"UserId" xorm:"user_id"`     // 变更者userId
	DBId       int    `json:"DBId" xorm:"db_id""`        // 变更库
}
