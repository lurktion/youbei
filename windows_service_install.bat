@echo off
set workdir=%~dp0
set ServerName=Youbei
set Instsrv=%workdir%WindowsService\instsrv.exe
set Srvany=%workdir%WindowsService\srvany.exe
set workpath=%workdir:\=\\%
set Child=HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%ServerName%\Parameters
set InstallCommand=%Instsrv% %ServerName% %Srvany%
set RemoveCommand=%Instsrv% %ServerName% remove
echo ################################################################################
echo. 
echo 本脚会将youbei.exe添加至本地服务，请使用管理员权限开启
echo. 
echo 如不需要以服务模式启动，直接双击youbei.exe运行即可
echo. 
echo ################################################################################
echo. 
echo 1.安装并启动
echo 2.卸载(请先停止服务)
echo 3.退出
echo.
set /p POP=请选择:

if "%POP%"=="1" goto install
if "%POP%"=="2" (goto remove) else goto end

:install
%InstallCommand%
reg add %Child%
reg add %Child% /v Application  /t REG_SZ /d "%workpath%"youbei.exe
reg add %Child% /v AppDirectory  /t REG_SZ /d "%workpath%"
sc start %ServerName%
goto end

:remove
sc stop %ServerName%
reg delete HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%ServerName% /f
%RemoveCommand%

goto end


:end
pause