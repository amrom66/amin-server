## 基于Gin + Vue + Element UI的前后端分离权限管理系统 

系统初始化极度简单，只需要配置文件中，修改数据库连接，系统启动后会自动初始化数据库信息以及必须的基础数据

## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库 

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具

- 表单构建工具

- 多命令模式

- TODO: 单元测试


## 🎁 内置

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9.  登录日志：系统登录日志记录查询包含登录异常。
10. 系统接口：根据业务代码自动生成相关的api接口文档。
11. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全部可视化编程，基本业务可以0代码实现。
12. 表单构建：自定义页面样式，拖拉拽实现页面布局。
13. 服务监控：查看一些服务器的基本信息。

## 准备工作

你需要在本地安装 [go] [gin] [node](http://nodejs.org/) 和 [git](https://git-scm.com/) 


## 🗞 系统架构

## 📦 本地开发

### 开发目录创建

```bash

# 创建开发目录
mkdir amin
cd amin
```

### 获取代码

> 重点注意：两个项目必须放在同一文件夹下；

```bash
# 获取后端代码
git clone git@github.com:linjinbao/amin-server.git

# 获取前端代码
git clone git@github.com:linjinbao/amin-ui.git

```

### 启动说明

#### 服务端启动说明

```bash
# 进入 admin 后端项目
cd ./admin

# 编译项目
go build

# 修改配置 
# 文件路径  config/settings.yml
vi ./config/setting.yml 

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

#### 初始化数据库，以及服务启动
```
# 首次配置需要初始化数据库资源信息
./amin migrate -c config/settings.yml


# 启动项目，也可以用IDE进行调试
./amin server -c config/settings.yml -p 8000

```

#### 使用docker 编译启动

```shell
# 编译镜像
docker build -t admin .

# 启动容器
docker run --name admin -p 8000:8000 -d admin
```

#### 文档生成

```bash
swag init  

# 如果没有swag命令 go get安装一下即可
go get -u github.com/swaggo/swag/cmd/swag
```

#### 交叉编译
```bash
env GOOS=windows GOARCH=amd64 go build main.go

# or

env GOOS=linux GOARCH=amd64 go build main.go
```

### UI交互端启动说明

```bash
# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run dev
```