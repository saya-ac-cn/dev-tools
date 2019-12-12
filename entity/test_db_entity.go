package entity

// 发布到测试版本实体类
type TestDbEntity struct {
	Id          int    `json:"Id" xorm:"id"`                    // 测试版本号
	VersionId   string `json:"VersionId" xorm:"version_id"`     // 测试版本号
	UserId      int8   `json:"UserId" xorm:"user_id"`           // 发布userId
	Status      int8   `json:"Status" xorm:"status"`            // 发布状态
	PublishTime string `json:"PublishTime" xorm:"publish_time"` // 发布时间

	TestDevId int8 `json:"TestDevId" xorm:"test_dev_id"` // 发布到测试详情
	DevId     int8 `json:"DevId" xorm:"dev_id"`          // 测试版本id
	TestId    int8 `json:"TestId" xorm:"test_id"`        // 开发id

	TestUser string `json:"TestUser" xorm:"test_user"` // 发布到测试用户
	DevUser  string `json:"DevUser" xorm:"dev_user"`   // 开发变更用户

	DBName   string `json:"DBName" xorm:"db_name"`     // 变更库
	DBItem   string `json:"DBItem" xorm:"db_item"`     // 变更内容
	DBReason string `json:"DBReason" xorm:"db_reason"` // 变更原因
}
