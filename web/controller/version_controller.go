package controller

import (
	"dev-tools/entity"
	"dev-tools/service"
	"dev-tools/tools"
	log "github.com/cihub/seelog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type VersionController struct {
	Ctx              iris.Context
	DBVersionService service.DbVersionService
	//当使用单例控制器时，由开发人员负责访问安全,所有客户端共享相同的控制器实例。
	//注意任何控制器的方法,是每个客户端，但结构的字段可以在多个客户端共享（如果是结构）
	//没有任何依赖于的动态struct字段依赖项
	//并且所有字段的值都不为零，在这种情况下我们使用uint64，它不是零（即使我们没有设置它手动易于理解的原因）因为它的值为＆{0}
	//以上所有都声明了一个Singleton，请注意，您不必编写一行代码来执行此操作，Iris足够聪明。
	//见`Get`
	Visits uint64
	//当前请求 Session
	//它的初始化是由我们添加到路由的依赖函数发生的。
	Session *sessions.Session
	Logger  log.LoggerInterface
}

// 数据库版本管理主页
// 访问/tools/view/home
func (c *VersionController) GetHome() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	// 判断对象是否为用户
	user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
	//c.Logger.Errorf("读取数据失败：%s", user.UserName)
	if isUser {
		data := c.DBVersionService.HomeData(user.Id)
		c.Ctx.View("home.html", iris.Map{
			"Title":    "数据库版本管理",
			"FuncName": "操作主页",
			"User":     user,
			"Data":     data,
		})
	} else {
		c.Logger.Error("获取用户会话信息失败")
		c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
	}
}

// 在开发库中修改（返回已修改但未发布）
func (c *VersionController) GetDev() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	dbid, err := c.Ctx.URLParamInt("db")
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		// 判断对象是否为用户
		user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
		//c.Logger.Errorf("读取数据失败：%s", user.UserName)
		if isUser {
			data := c.DBVersionService.DevUnPublish(user.Id, dbid)
			c.Ctx.View("dev.html", iris.Map{
				"Title":    "数据库版本管理",
				"FuncName": "修改开发库",
				"DbId":     dbid,
				"User":     user,
				"Data":     data,
			})
		} else {
			c.Logger.Error("获取用户会话信息失败")
			c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
		}
	}
}

// 在开发库中修改（返回已修改但未发布）
func (c *VersionController) GetTest() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	dbid, err := c.Ctx.URLParamInt("db")
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		// 判断对象是否为用户
		user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
		//c.Logger.Errorf("读取数据失败：%s", user.UserName)
		if isUser {
			data := c.DBVersionService.DevRecentlyData(user.Id, dbid)
			c.Ctx.View("test.html", iris.Map{
				"Title":    "数据库版本管理",
				"FuncName": "发布到测试",
				"DbId":     dbid,
				"User":     user,
				"Data":     data,
			})
		} else {
			c.Logger.Error("获取用户会话信息失败")
			c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
		}
	}
}

// 在测试库中修改（返回已修改但未发布）
func (c *VersionController) GetPro() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	dbid, err := c.Ctx.URLParamInt("db")
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		data := c.DBVersionService.ProRecentlyData(dbid)
		c.Ctx.View("pro.html", iris.Map{
			"Title":    "数据库版本管理",
			"FuncName": "发布到线上",
			"DbId":     dbid,
			"Data":     data,
		})
	}
}

// 保存开发中的修改
func (c *VersionController) PostDev() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	parmas := &[]entity.DevDbEntity{}
	err := c.Ctx.ReadJSON(parmas)
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		datalist := make([]entity.DevDbEntity, 0)
		user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
		if isUser {
			for _, param := range *parmas {
				datalist = append(datalist, param)
			}
			changed := c.DBVersionService.DevInsertChange(datalist, user.Id)
			if changed {
				c.Ctx.JSON(tools.NewResult())
			} else {
				c.Ctx.JSON(tools.NewResultError(-3, "添加失败"))
			}
		} else {
			c.Logger.Error("获取用户会话信息失败")
			c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
		}
	}
}

// 移除开发中的修改
func (c *VersionController) DeleteDev() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	parmas := &entity.DevDbEntity{}
	err := c.Ctx.ReadJSON(parmas)
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		changed := c.DBVersionService.DevRemoveChange(parmas.Id)
		if changed {
			c.Ctx.JSON(tools.NewResult())
		} else {
			c.Ctx.JSON(tools.NewResultError(-3, "移除失败"))
		}
	}
}

// 发布到测试
func (c *VersionController) PostTest() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	versionId := c.Ctx.URLParam("versionId")
	if "" == versionId {
		c.Ctx.JSON(tools.NewResultError(-1, "非法参数请求"))
	}
	parmas := &[]entity.TestDbEntity{}
	err := c.Ctx.ReadJSON(parmas)
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		datalist := make([]entity.TestDbEntity, 0)
		user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
		if isUser {
			for _, param := range *parmas {
				datalist = append(datalist, param)
			}
			changed := c.DBVersionService.TestPublish(datalist, user.Id, versionId)
			if changed {
				c.Ctx.JSON(tools.NewResult())
			} else {
				c.Ctx.JSON(tools.NewResultError(-3, "发布失败"))
			}
		} else {
			c.Logger.Error("获取用户会话信息失败")
			c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
		}
	}
}

// 发布到线上
func (c *VersionController) PostPro() {
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	versionId := c.Ctx.URLParam("versionId")
	if "" == versionId {
		c.Ctx.JSON(tools.NewResultError(-1, "非法参数请求"))
	}
	parmas := &[]entity.ProDbEntity{}
	err := c.Ctx.ReadJSON(parmas)
	if nil != err {
		c.Logger.Errorf("读取用户请求数据失败：%s", err)
		c.Ctx.JSON(tools.NewResultError(-2, "非法参数请求"))
	} else {
		datalist := make([]entity.ProDbEntity, 0)
		user, isUser := (c.Session.Get("user")).(*entity.UserEntity)
		if isUser {
			for _, param := range *parmas {
				datalist = append(datalist, param)
			}
			changed := c.DBVersionService.ProPublish(datalist, user.Id, versionId)
			if changed {
				c.Ctx.JSON(tools.NewResult())
			} else {
				c.Ctx.JSON(tools.NewResultError(-3, "发布失败"))
			}
		} else {
			c.Logger.Error("获取用户会话信息失败")
			c.Ctx.JSON(tools.NewResultError(-5, "获取用户会话信息失败"))
		}
	}
}
