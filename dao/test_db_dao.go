package dao

import (
	"bytes"
	"dev-tools/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// 发布到测试DAO层

var (
	testDbOnce     sync.Once
	testDbInstance *TestDbDao
)

type TestDbDao struct {
	engine *xorm.Engine
}

func NewTestDbDao(engine *xorm.Engine) *TestDbDao {
	testDbOnce.Do(func() {
		testDbInstance = &TestDbDao{
			engine: engine,
		}
		fmt.Println("TestDbDao,instance...")
	})
	return testDbInstance
}

// 查看该用户所属或参与到还没有发布到线上的变更内容
func (d *TestDbDao) GetDbChange(userid int) []entity.TestDbEntity {
	datalist := make([]entity.TestDbEntity, 0)
	err := error(nil)
	sql := "select a.id,a.version_id,a.user_id,a.publish_time,b.`status`,b.id as test_dev_id,b.dev_id,b.test_id,e.user_name as test_user,f.user_name as dev_user,d.db_name,c.db_item from d_test_publish_version a inner join d_test_publish_info b on a.id=b.test_id inner join d_dev_db_change c on b.dev_id = c.id inner join d_db_owner d on c.db_id=d.id left join d_user e on a.user_id = e.id left join d_user f on c.user_id = f.id where b.`status`=1 and (a.user_id = ? or c.user_id=?)"
	err = d.engine.SQL(sql, userid, userid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 查看最近发布到测试的5个版本
func (d *TestDbDao) GetRecently5(dbid int) []entity.TestDbEntity {
	datalist := make([]entity.TestDbEntity, 0)
	err := error(nil)
	sql := "select a.id,a.version_id,a.publish_time from d_test_publish_version a left join d_test_publish_info b on a.id=b.test_id left join d_dev_db_change c on c.id=b.dev_id where c.db_id=? group by a.id order by a.id limit 5"
	err = d.engine.SQL(sql, dbid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 发布到测试(最后返回主键)
func (d *TestDbDao) TestPublish(userId int, versionId string) int64 {
	err := error(nil)
	sql := "insert into `d_test_publish_version` ( `version_id`, `user_id`, `publish_time`) values ( ?, ?, now());"
	result, err := d.engine.Exec(sql, versionId, userId)
	if nil != err {
		log.Println(err)
		return -1
	}
	num, err := result.LastInsertId()
	if err != nil {
		panic(err)
		log.Println(err)
		return -2
	}
	return num
}

// 添加测试发布详情
func (d *TestDbDao) TestPublishInfo(devId int8, testId int64) int64 {
	err := error(nil)
	sql := "insert into `d_test_publish_info` ( `dev_id`, `test_id`) values ( ?, ?);"
	result, err := d.engine.Exec(sql, devId, testId)
	if nil != err {
		log.Println(err)
		return -1
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(err)
		log.Println(err)
		return -2
	}
	return num
}

// 查看还未发布到测试的5个版本明细
func (d *TestDbDao) GetUnPublish(dbid int) []entity.TestDbEntity {
	datalist := make([]entity.TestDbEntity, 0)
	err := error(nil)
	sql := "select a.id,b.id as test_dev_id,a.version_id,a.publish_time,c.db_item,c.db_reason,d.db_name from d_test_publish_version a left join d_test_publish_info b on a.id=b.test_id left join d_dev_db_change c on c.id=b.dev_id left join d_db_owner d on d.id = c.db_id where c.db_id=? and b.`status`=1 order by a.version_id"
	err = d.engine.SQL(sql, dbid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 修改测试明细的发布状态
func (d *TestDbDao) TestEditStatus(testIds []int8) int64 {
	lenth := len(testIds)
	if lenth <= 0 {
		return -1
	}
	args := make([]interface{}, 0)
	// 封装语句
	var sql bytes.Buffer
	sql.WriteString("update d_test_publish_info set `status` = 2 where id in (")
	for i := 0; i < lenth; i++ {
		if i == (lenth - 1) {
			sql.WriteString(" ?")
			break
		} else {
			sql.WriteString(" ?,")
		}
	}
	sql.WriteString(" ) and `status` = 1")
	// 封装参数
	args = append(args, sql.String())
	for i := 0; i < lenth; i++ {
		args = append(args, testIds[i])
	}
	err := error(nil)
	result, err := d.engine.Exec(args...)
	if nil != err {
		log.Println(err)
		return -2
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(err)
		log.Println(err)
		return -3
	}
	return num
}

// 查看测试版本详情
func (d *TestDbDao) GetTestInfo(where entity.TestDbEntity) []entity.TestDbEntity {
	args := make([]interface{}, 0)
	datalist := make([]entity.TestDbEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("select a.id,a.version_id,a.publish_time,e.user_name as test_user,c.db_item,c.db_reason,f.user_name as dev_user,d.db_name from d_test_publish_version a left join d_test_publish_info b on a.id = b.test_id left join d_dev_db_change c on c.id = b.dev_id left join d_db_owner d on d.id = c.db_id left join d_user e on a.user_id = e.id left join d_user f on c.user_id = f.id where 1 = 1")
	if where.Id != 0 {
		sql.WriteString(" and a.id = ?")
		args = append(args, where.Id)
	}
	if where.VersionId != "" {
		sql.WriteString(" and a.version_id = ?")
		args = append(args, where.VersionId)
	}
	sql.WriteString(" order by a.id desc")
	err = d.engine.SQL(sql.String(), args...).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}
