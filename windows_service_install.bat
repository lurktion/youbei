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
echo ���ŻὫyoubei.exe��������ط�����ʹ�ù���ԱȨ�޿���
echo. 
echo �粻��Ҫ�Է���ģʽ������ֱ��˫��youbei.exe���м���
echo. 
echo ################################################################################
echo. 
echo 1.��װ������
echo 2.ж��(����ֹͣ����)
echo 3.�˳�
echo.
set /p POP=��ѡ��:

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