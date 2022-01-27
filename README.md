### Go语言开发博客。<br>
Go+Gin+Gorm
可查看我的博客：http://blog.caixiaoxin.cn

### 功能
```diff
+ <span style="color:	#1E90FF;font-size:15px">创建管理员<br>
+ <span style="color:	#1E90FF;font-size:15px">账号列表<br>
+ <span style="color:	#1E90FF;font-size:15px">删除账号<br>
+ <span style="color:	#1E90FF;font-size:15px">删除账号<br>
+ <span style="color:	#1E90FF;font-size:15px">账号详情<br>
+ <span style="color:	#1E90FF;font-size:15px">修改账号信息<br>
+ <span style="color:	#1E90FF;font-size:15px">禁用<br>
+ <span style="color:	#1E90FF;font-size:15px">启用<br>
+ <span style="color:	#1E90FF;font-size:15px">创建标签<br>
+ <span style="color:	#1E90FF;font-size:15px">标签列表<br>
+ <span style="color:	#1E90FF;font-size:15px">修改标签名字<br>
+ <span style="color:	#1E90FF;font-size:15px">删除标签<br>
+ <span style="color:	#1E90FF;font-size:15px">创建分类<br>
+ <span style="color:	#1E90FF;font-size:15px">分类列表<br>
+ <span style="color:	#1E90FF;font-size:15px">修改分类名字<br>
+ <span style="color:	#1E90FF;font-size:15px">删除分类<br>
+ <span style="color:	#1E90FF;font-size:15px">写文章<br>
+ <span style="color:	#1E90FF;font-size:15px">文章列表<br>
+ <span style="color:	#1E90FF;font-size:15px">文章修改<br>
+ <span style="color:	#1E90FF;font-size:15px">删除文章<br>
+ <span style="color:	#1E90FF;font-size:15px">评论<br>
+ <span style="color:	#1E90FF;font-size:15px">评论列表<br>
+ <span style="color:	#1E90FF;font-size:15px">删除评论<br>
+ <span style="color:	#1E90FF;font-size:15px">添加友联<br>
+ <span style="color:	#1E90FF;font-size:15px">修改友联<br>
+ <span style="color:	#1E90FF;font-size:15px">删除友联<br>

```
### 设计目录结构
```shell
目录结构
	constants                    # 定义常量目录
		state.go 
	db                              # 数据库连接初始化
		core.go
	handler                      # 控制器
		......
	middleware               # 中间件
		......
	models                      # 模型
		......
	routers                      # 路由
		router.go
	service                      # 服务层，业务逻辑处理，与models交互操作数据库
		......
	types                        # 状态码，状态信息
		code.go
	util                           # 工具包
		encrypt.go
		result.go
	.gitignore
	gin-blog
	go.mod 
	main.go
	README.md
``` 
### 本机部署
```git
git clone git@github.com:jeffcail/gin-blog.git
```
```shell
➜ cd gin-blog
➜ gowatch gin-blog
```
![img.png](go-gowatch.png)

### 测试部署是否成功
![img_1.png](login.png)
