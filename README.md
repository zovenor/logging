# Logging
```bash
go get github.com/zovenor/logging/v2
```

---

### [Example 1](./examples/example_1/main.go)
Logs:
```logfile
[info]2023/11/13 16:39:49.760819 logging.go:91: info message
[success]2023/11/13 16:39:49.760907 logging.go:91: success message
[fatal]2023/11/13 16:39:49.760938 logging.go:91: error message
[warning]2023/11/13 16:39:49.760965 logging.go:91: warning message
```
Terminal:
```shell
[17:33:44] success message
[17:33:44] success message
[17:33:44] fatal message
[17:33:44] fatal message
[17:33:44] warning message
[17:33:44] warning message
[17:33:44] info message
[17:33:44] info message
[17:33:44] info message: 10000
[17:33:44] info message
[17:33:44] info message: 999 true
```