package service

import (
	"crypto/md5"
	"dev-tools/dao"
	"dev-tools/datasource"
	"dev-tools/entity"
	"encoding/hex"
	"fmt"
	"sync"
)

// 数据库版本管理
var (
	dbVersionServiceOnce     sync.Once
	dbVersionServiceInstance DbVersionService
)

type DbVersionService interface {
	UserLogin(parma entity.UserEntity) *entity.UserEntity
	HomeData(userid int) interface{}
	DevUnPublish(userid int, dbid int) []entity.DevDbEntity
	DevInsertChange(db []entity.DevDbEntity, userid int) bool
	DevRemoveChange(id int) bool
	DevRecentlyData(dbid int, userid int) interface{}
	TestPublish(devList []entity.TestDbEntity, userId int, versionId string) bool
	ProRecentlyData(dbid int) interface{}
	ProPublish(testList []entity.ProDbEntity, userId int, versionId string) bool
	GetProInfo(where entity.ProDbEntity) []entity.ProDbEntity
	GetTestInfo(where entity.TestDbEntity) []entity.TestDbEntity
}

type dbVersionService struct {
	userDao   *dao.UserDao
	dbDao     *dao.DbDao
	devDbDao  *dao.DevDbDao
	testDbDao *dao.TestDbDao
	proDbDao  *dao.ProDbDao
}

func NewUserService() DbVersionService {
	dbVersionServiceOnce.Do(func() {
		dbVersionServiceInstance = &dbVersionService{
			userDao:   dao.NewUserDao(datasource.InstanceMaster()),
			dbDao:     dao.NewDbDao(datasource.InstanceMaster()),
			devDbDao:  dao.NewDevDbDao(datasource.InstanceMaster()),
			testDbDao: dao.NewTestDbDao(datasource.InstanceMaster()),
			proDbDao:  dao.NewProDbDao(datasource.InstanceMaster()),
		}
		fmt.Println("NewUserService,instance...")
	})
	return dbVersionServiceInstance
}

// 登录
func (s *dbVersionService) UserLogin(parma entity.UserEntity) *entity.UserEntity {
	h := md5.New()
	h.Write([]byte(parma.UserPassword)) // 需要加密的字符串为 123456
	md5Pwd := hex.EncodeToString(h.Sum(nil))
	//fmt.Printf("%s\n", md5Pwd) // 输出加密结果
	user := s.userDao.GetByAccount(parma.UserAccount)
	if user.UserPassword != md5Pwd {
		// 密码错误
		return nil
	} else {
		return user
	}
}

// 返回用户主页数据
func (s *dbVersionService) HomeData(userid int) interface{} {
	db := make(map[string]interface{})
	// 查询该用户可管理的数据库
	dbList := s.dbDao.GetOwnerDB(userid)
	db["db"] = dbList
	devList := s.devDbDao.GetDbChange(userid)
	db["devdb"] = devList
	testList := s.testDbDao.GetDbChange(userid)
	db["test"] = testList
	return db
}

// 返回开发已修改，但在测试还没有发布
func (s *dbVersionService) DevUnPublish(userid int, dbid int) []entity.DevDbEntity {
	return s.devDbDao.GetDbChangeByDb(userid, dbid)
}

// 添加开发变更
func (s *dbVersionService) DevInsertChange(db []entity.DevDbEntity, userid int) bool {
	for _, item := range db {
		item.UserId = userid
		s.devDbDao.DevInsertChange(item)
	}
	return true
}

// 移除开发变更
func (s *dbVersionService) DevRemoveChange(id int) bool {
	rows := s.devDbDao.DevRemoveChange(id)
	if rows <= 0 {
		return false
	}
	return true
}

// 查看最近发布到测试和线上的5个版本以及开发未发布的变更
func (s *dbVersionService) DevRecentlyData(dbid int, userid int) interface{} {
	db := make(map[string]interface{})
	// 开发还未发发布的子项
	unPublish := s.devDbDao.GetDbChangeByDb(userid, dbid)
	db["dev"] = unPublish
	// 查看最近发布到测试的5个版本
	dbList := s.testDbDao.GetRecently5(dbid)
	db["test"] = dbList
	// 查看最近发布到线上的5个版本
	devList := s.proDbDao.GetRecently5(dbid)
	db["pro"] = devList
	return db
}

// 发布到测试
func (s *dbVersionService) TestPublish(devList []entity.TestDbEntity, userId int, versionId string) bool {
	// 先在开发中设置发布状态
	devIds := make([]int8, 0)
	for _, param := range devList {
		devIds = append(devIds, param.DevId)
	}
	devRows := s.devDbDao.DevEditStatus(devIds)
	if devRows <= 0 {
		return false
	}
	// 添加一条测试版本信息
	testId := s.testDbDao.TestPublish(userId, versionId)
	if testId <= 0 {
		return false
	}
	for _, param := range devList {
		s.testDbDao.TestPublishInfo(param.DevId, testId)
	}
	return true
}

// 查看最近发布到线上的5个版本以及测试未发布的变更
func (s *dbVersionService) ProRecentlyData(dbid int) interface{} {
	db := make(map[string]interface{})
	// 查看最近发布到测试的5个版本
	dbList := s.testDbDao.GetUnPublish(dbid)
	db["test"] = dbList
	// 查看最近发布到线上的5个版本
	devList := s.proDbDao.GetRecently5(dbid)
	db["pro"] = devList
	return db
}

// 发布到线上
func (s *dbVersionService) ProPublish(proList []entity.ProDbEntity, userId int, versionId string) bool {
	// 先在测试中设置发布状态
	testIds := make([]int8, 0)
	for _, param := range proList {
		testIds = append(testIds, param.TestId)
	}
	testRows := s.testDbDao.TestEditStatus(testIds)
	if testRows <= 0 {
		return false
	}
	// 添加一条线上版本信息
	onlineId := s.proDbDao.TestPublish(userId, versionId)
	if onlineId <= 0 {
		return false
	}
	for _, param := range proList {
		s.proDbDao.TestPublishInfo(param.TestId, onlineId)
	}
	return true
}

// 查看线上版本明细
func (s *dbVersionService) GetProInfo(where entity.ProDbEntity) []entity.ProDbEntity {
	return s.proDbDao.GetProInfo(where)
}

// 查看测试版本明细
func (s *dbVersionService)GetTestInfo(where entity.TestDbEntity) []entity.TestDbEntity{
	return s.testDbDao.GetTestInfo(where)
}
