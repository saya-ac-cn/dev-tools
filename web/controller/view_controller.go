package controller

import (
	"github.com/kataras/iris"
)

// 视图控制器
type ViewController struct {
	Ctx iris.Context
	//MapService service.MapService
	//当使用单例控制器时，由开发人员负责访问安全,所有客户端共享相同的控制器实例。
	//注意任何控制器的方法,是每个客户端，但结构的字段可以在多个客户端共享（如果是结构）
	//没有任何依赖于的动态struct字段依赖项
	//并且所有字段的值都不为零，在这种情况下我们使用uint64，它不是零（即使我们没有设置它手动易于理解的原因）因为它的值为＆{0}
	//以上所有都声明了一个Singleton，请注意，您不必编写一行代码来执行此操作，Iris足够聪明。
	//见`Get`
	Visits uint64
}

// 生成java实体类页面
// 访问/tools/view/javaentity
func (c *ViewController) GetJavaentity() {
	c.Ctx.View("entity.html", iris.Map{
		"Title":    "映射实体对象",
		"FuncName": "生成实体",
	})
}

// 登录页面
// 访问/tools/view/login
func (c *ViewController) GetLogin() {
	c.Ctx.View("login.html", iris.Map{
		"Title":    "数据库版本管理",
		"FuncName": "身份认证",
	})
}
