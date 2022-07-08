@echo off
if exist ..\000-microserver\microserver.exe goto okserver

echo compile microserver...
echo.
SET PATHLOC=%~dp0
cd ..\000-microserver
call compile.bat
cd %PATHLOC%

:okserver

echo start microserver...
echo.
..\000-microserver\microserver.exe
