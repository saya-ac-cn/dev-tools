@echo off
rem ���ý�����뻷��
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
SET Project=target
if exist %Project% (
   	rem %Project%�Ѵ��ڣ�����ɾ��.....
    echo %Project%�Ѵ��ڣ�����ɾ��.....
	rd/s/q %Project%
)
rem ����%Project%Ŀ¼.....
echo ����%Project%Ŀ¼.....
md   %Project%
::go build -ldflags "-H windowsgui" -o main.go
echo ������Ŀ��.....
go build -o dev-tools
echo ͬ�������ļ�.....
XCOPY conf\*.*  %Project%\conf\  /s /e
XCOPY web\public\*.*  %Project%\web\public\  /s /e
XCOPY web\view\*.*  %Project%\web\view\  /s /e
MOVE  dev-tools %Project%\
pause