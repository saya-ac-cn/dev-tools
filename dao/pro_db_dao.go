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
	proDbOnce     sync.Once
	proDbInstance *ProDbDao
)

type ProDbDao struct {
	engine *xorm.Engine
}

func NewProDbDao(engine *xorm.Engine) *ProDbDao {
	proDbOnce.Do(func() {
		proDbInstance = &ProDbDao{
			engine: engine,
		}
		fmt.Println("ProDbDao,instance...")
	})
	return proDbInstance
}

// 查看线上版本信息
func (d *ProDbDao) GetProInfo(where entity.ProDbEntity) []entity.ProDbEntity {
	args := make([]interface{}, 0)
	datalist := make([]entity.ProDbEntity, 0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("select a.id,a.version_id,a.user_id,f.user_name as pro_test_user,a.publish_time,b.test_id,b.online_id,i.version_id as test_version_id,g.user_name as dev_test_user,d.db_item,d.db_reason,h.user_name as dev_user,e.db_name from d_online_publish_version a inner join d_online_publish_info b on a.id=b.online_id inner join d_test_publish_info c on b.test_id=c.id inner join d_dev_db_change d on c.dev_id = d.id inner join d_db_owner e on d.db_id = e.id inner join d_test_publish_version i on i.id = c.test_id left join d_user f on a.user_id = f.id left join d_user g on i.user_id = g.id left join d_user h on d.user_id = h.id where 1=1")
	if where.Id != 0 {
		sql.WriteString(" and a.id = ?")
		args = append(args, where.Id)
	}
	if where.VersionId != "" {
		sql.WriteString(" and a.version_id = ?")
		args = append(args, where.VersionId)
	}
	err = d.engine.SQL(sql.String(), args...).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 查看最近发布到线上的5个版本
func (d *ProDbDao) GetRecently5(dbid int) []entity.ProDbEntity {
	datalist := make([]entity.ProDbEntity, 0)
	err := error(nil)
	sql := "select a.id,a.version_id,a.publish_time from d_online_publish_version a left join d_online_publish_info b on a.id = b.online_id left join d_test_publish_info c on b.test_id = c.id left join d_test_publish_version d on d.id = c.test_id left join d_dev_db_change e on c.id = c.dev_id where e.db_id = ? group by a.id order by a.id desc limit 5"
	err = d.engine.SQL(sql, dbid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 发布到线上(最后返回主键)
func (d *ProDbDao) TestPublish(userId int, versionId string) int64 {
	err := error(nil)
	sql := "insert into `d_online_publish_version` ( `version_id`, `user_id`, `publish_time`) values ( ?, ?, now());"
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

// 添加线上发布详情
func (d *ProDbDao) TestPublishInfo(testId int8, onlineId int64) int64 {
	err := error(nil)
	sql := "insert into `d_online_publish_info` ( `test_id`, `online_id`) values ( ?, ?);"
	result, err := d.engine.Exec(sql, testId, onlineId)
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
