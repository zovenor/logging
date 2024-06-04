# Logging
```bash
go get github.com/zovenor/logging/v2
```

---

### [Example 1](./examples/example_1/main.go)
Logs:
```logfile
[success]2024/06/04 18:36:34 success message
[success]2024/06/04 18:36:34 success message
[fatal]2024/06/04 18:36:34 
	> /examples/example_1/main.go:13
	fatal message
[fatal]2024/06/04 18:36:34 
	> /examples/example_1/main.go:15
	fatal message
[warning]2024/06/04 18:36:34 warning message
[warning]2024/06/04 18:36:34 warning message
[info]2024/06/04 18:36:34 info message
[info]2024/06/04 18:36:34 info message

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