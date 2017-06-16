#### 使用
```
go build
./devcloud
````

#### 说明
- 模拟Jenkins的参数化编译，执行编译机上的shell
- 支持定时任务持续集成
- 默认运行时定时任务信息每10s会保存到./conf/croninfo文件中，重启时会从该文件读取定时任务信息
