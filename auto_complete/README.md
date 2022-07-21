# 自动选项

### 服务器启动后界面可以自定义命令选择结果

- 可以通过斜线/ 触发对应的命令选择必选和可选的逻辑， 
   目前看适合提供分支选择功能，触发对应的逻辑业务
-  -guild 选择服务器id 
```
 go run main.go -t xxxxx -guild xxxxxx rmcmd false
```