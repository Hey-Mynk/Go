@echo off
if "%~1"=="" (
    echo Usage: %0 project_name
    exit /b 1
)

set PROJECT_NAME=%~1

mkdir %PROJECT_NAME%
cd %PROJECT_NAME%
mkdir cmd pkg internal api configs docs scripts build deploy web tests

echo package main > cmd\main.go
echo. >> cmd\main.go
echo import "fmt" >> cmd\main.go
echo. >> cmd\main.go
echo func main() { >> cmd\main.go
echo. >> cmd\main.go
echo     fmt.Println("Hello, world!") >> cmd\main.go
echo } >> cmd\main.go

go mod init %PROJECT_NAME%

echo Project %PROJECT_NAME% created successfully.
