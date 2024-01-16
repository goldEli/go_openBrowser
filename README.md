
# build
```shell
# build liunx
env GOOS=linux GOARCH=amd64 go build -o openBrowser main.go

# build windows
 go build -o openBrowser.exe main.go
```

# execute
```shell
go run -url='https://oa.dev.staryuntech.com/getAttendanceStatistics?date=2024-01-15&companyId=1504303190303051778'

```