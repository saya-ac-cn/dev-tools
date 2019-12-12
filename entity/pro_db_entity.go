package entity

// 发布到线上版本实体类
type ProDbEntity struct {
	Id          int    `json:"Id" xorm:"id"`                    // 测试版本id
	VersionId   string `json:"VersionId" xorm:"version_id"`     // 测试版本号
	UserId      int8   `json:"UserId" xorm:"user_id"`           // 发布userId
	Status      int8   `json:"Status" xorm:"status"`            // 发布状态
	PublishTime string `json:"PublishTime" xorm:"publish_time"` // 发布时间
	ProTestUser string `json:"ProTestUser" xorm:"pro_test_user"`

	OnlineId int8 `json:"OnlineId" xorm:"online_id"` // 测试版本id
	TestId   int8 `json:"TestId" xorm:"test_id"`     // 测试id

	TestVersionId string `json:"TestVersionId" xorm:"test_version_id"` // 测试版本号
	DevTestUser   string `json:"DevTestUser" xorm:"dev_test_user"`     // 发布到测试用户

	DBName   string `json:"DBName" xorm:"db_name"`     // 变更库
	DBItem   string `json:"DBItem" xorm:"db_item"`     // 变更内容
	DbReason string `json:"DbReason" xorm:"db_reason"` // 变更原因
	DevUser  string `json:"DevUser" xorm:"dev_user"`   // 操作者
}
