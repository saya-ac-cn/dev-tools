@echo off
rem 设置交叉编译环境
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
SET Project=target
if exist %Project% (
   	rem %Project%已存在，正在删除.....
    echo %Project%已存在，正在删除.....
	rd/s/q %Project%
)
rem 创建%Project%目录.....
echo 创建%Project%目录.....
md   %Project%
::go build -ldflags "-H windowsgui" -o main.go
echo 编译项目中.....
go build -o dev-tools
echo 同步依赖文件.....
XCOPY conf\*.*  %Project%\conf\  /s /e
XCOPY web\public\*.*  %Project%\web\public\  /s /e
XCOPY web\view\*.*  %Project%\web\view\  /s /e
MOVE  dev-tools %Project%\
pause