# dev-tools
## 项目开发工具包
* 开发语言 Go
* 采用的MVC框架Iris（单例控制器）
* 前端采用原生Html开发
* 项目全局采用单例模式
* 模板引擎参考：https://blog.csdn.net/sryan/article/details/52353937
## 部署&运行方式
* go build ./main.go
* Mac&Linux直接./main 即可运行。可以在运行时加端口参数p，指定运行在哪个端口
* Windows双夹即可运行
* linux后台运行nohup ./dev-tools >./warn.log 2>&1 &
## 当前功能
* 自动生成Java实体对象
