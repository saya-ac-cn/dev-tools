package dao

import (
	"bytes"
	"dev-tools/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// 开发数据库DAO层

var (
	devDbOnce     sync.Once
	devDbInstance *DevDbDao
)

type DevDbDao struct {
	engine *xorm.Engine
}

func NewDevDbDao(engine *xorm.Engine) *DevDbDao {
	devDbOnce.Do(func() {
		devDbInstance = &DevDbDao{
			engine: engine,
		}
		fmt.Println("DevDbDao,instance...")
	})
	return devDbInstance
}

// 查看还没有发布到测试的变更内容
func (d *DevDbDao) GetDbChange(userid int) []entity.DevDbEntity {
	datalist := make([]entity.DevDbEntity, 0)
	err := error(nil)
	sql := "select c.id,b.db_name,c.db_item,c.db_reason,c.create_time,c.`status`,c.user_id,d.user_name from d_user_mana_db a inner join d_db_owner b on a.user_id = b.id inner join d_dev_db_change c on c.db_id=b.id left join d_user d on d.id = c.user_id where a.user_id = ? and c.`status`=1"
	err = d.engine.SQL(sql, userid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 查看还没有发布到测试的变更内容
func (d *DevDbDao) GetDbChangeByDb(userid int, dbid int) []entity.DevDbEntity {
	datalist := make([]entity.DevDbEntity, 0)
	err := error(nil)
	sql := "select c.id,b.db_name,c.db_item,c.db_reason,c.create_time,c.`status`,c.user_id,d.user_name from d_user_mana_db a inner join d_db_owner b on a.user_id = b.id inner join d_dev_db_change c on c.db_id=b.id left join d_user d on d.id = c.user_id where b.id = ? and a.user_id = ? and c.`status`=1"
	err = d.engine.SQL(sql, dbid, userid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 添加开发变更
func (d *DevDbDao) DevInsertChange(db entity.DevDbEntity) int64 {
	err := error(nil)
	sql := "insert into d_dev_db_change(`user_id`,`db_id` ,`db_item`,`db_reason`,`create_time`) values(?,?,?,?,now())"
	result, err := d.engine.Exec(sql, db.UserId, db.DBId, db.DBItem, db.DBReason)
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

// 移除开发变更
func (d *DevDbDao) DevRemoveChange(id int) int64 {
	err := error(nil)
	sql := "delete from d_dev_db_change where id=?"
	result, err := d.engine.Exec(sql, id)
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

func (d *DevDbDao) DevEditStatus(devids []int8) int64 {
	lenth := len(devids)
	if lenth <= 0 {
		return -1
	}
	args := make([]interface{}, 0)
	// 封装语句
	var sql bytes.Buffer
	sql.WriteString("update d_dev_db_change set `status` = 2 where id in (")
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
		args = append(args, devids[i])
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
