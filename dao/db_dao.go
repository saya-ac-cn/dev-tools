package dao

import (
	"dev-tools/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// 数据库DAO层

var (
	dbOnce     sync.Once
	dbInstance *DbDao
)

type DbDao struct {
	engine *xorm.Engine
}

func NewDbDao(engine *xorm.Engine) *DbDao {
	dbOnce.Do(func() {
		dbInstance = &DbDao{
			engine: engine,
		}
		fmt.Println("DbDao,instance...")
	})
	return dbInstance
}

// 查看自己可管理的数据库
func (d *DbDao) GetOwnerDB(userid int) []entity.DbAttributeEntity {
	datalist := make([]entity.DbAttributeEntity, 0)
	err := error(nil)
	sql := "select b.* from d_user_mana_db a inner join d_db_owner b on a.user_id = ?"
	err = d.engine.SQL(sql, userid).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}
