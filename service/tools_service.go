package service

import (
	"bytes"
	"dev-tools/entity"
	"dev-tools/tools"
	"fmt"
	"sync"
)

var (
	toolsServiceOnce     sync.Once
	toolsServiceInstance ToolsService
)

type ToolsService interface {
	// 根据数据库表自动生成Java实体对象
	MappingJavaEntity(parma entity.DBEntity) *tools.ResultUtil
}

type toolsService struct {
}

func NewToolsService() ToolsService {
	toolsServiceOnce.Do(func() {
		toolsServiceInstance = &toolsService{}
	})
	return toolsServiceInstance
}

// 根据数据库表自动生成Java实体对象
func (s *toolsService) MappingJavaEntity(parma entity.DBEntity) *tools.ResultUtil {
	log := tools.GetLoggerInstance().Logger
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	list, err := tools.GetDBConnetion(parma.UserName, parma.PassWord, parma.IpAddr, parma.Port, parma.DBName, parma.TableName)
	if err != nil {
		log.Errorf("读取数据失败：%s", err)
		return tools.NewResultError(-2, fmt.Sprintf("Java实体模型生成失败：%s", err))
	}
	var outStr bytes.Buffer
	outStr.WriteString("public class ")
	outStr.WriteString(tools.CreateEntityName("liar"))
	outStr.WriteString(" {\n")
	// 生成字段
	for _, val := range list {
		// 循环处理数据库字段 -> columnComment:编号 columnName:id dataType:int
		outStr.WriteString("\n    /**\n")
		outStr.WriteString("     * ")
		outStr.WriteString(fmt.Sprintf("%s", val["columnComment"]))
		outStr.WriteString("\n")
		outStr.WriteString("     */\n")
		outStr.WriteString("    private ")
		outStr.WriteString(tools.FieldTypeChange(fmt.Sprintf("%s", val["dataType"])))
		outStr.WriteString(" ")
		outStr.WriteString(tools.CreateFieldName(fmt.Sprintf("%s", val["columnName"])))
		outStr.WriteString(";\n")
	}
	outStr.WriteString("\n    public ")
	// 生成构造函数
	outStr.WriteString(tools.CreateEntityName("liar"))
	outStr.WriteString("() {\n")
	outStr.WriteString("    }\n")
	outStr.WriteString("\n}")
	return tools.NewResultSuccess(outStr.String())
}
