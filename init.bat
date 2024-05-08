@echo off
echo Ejecutando...

start /MIN cmd /c "cd ./client && npm start"
start "Server" /MIN cmd /c "cd ./server && go run ./main.go"


rem Esperar a que se cierre la ventana principal
pause >nul

rem Cerrar las consolas secundarias
taskkill /FI "WINDOWTITLE eq Windows PowerShell"
taskkill /FI "WINDOWTITLE eq Server"
