## 个人快速建立一个简单的项目

[![Build Status](https://app.travis-ci.com/catbugdemo/project_order.svg?branch=master)](https://app.travis-ci.com/catbugdemo/project_order)
[![Go Report Card](https://goreportcard.com/badge/github.com/catbugdemo/project_order)](https://goreportcard.com/report/github.com/catbugdemo/project_order)

### 1.项目结构体
```text
project_order
    |- configs      # 配置包
    |   |- local.toml   # 配置文件
    |- inits        # 初始化包
    |   |- config.go    # 配置初始化
    |   |- db.go        # 数据库初始化
    |   |- http.go      # http初始化
    |   |- redis.go     # 缓存初始化
    |   |- rocketmq.go  # 消息队列初始化
    |- pkg          # 放置各种接口和模块的包
    |- utils            # 工具包
    |   |- logger       # 自定义 log
    |   |   |- log.go   # log 初始化
    |- vendor       # 第三方包放置位置
    |- go.mod       # 项目依赖文件
    |- main.go      # 主函数
    |- README.md    # 说明 
```

### 2.第三方包
- http
    - github.com/gin-gonic/gin
    
- database
    - github.com/jinzhu/gorm 
    
- redis 
    - github.com/gomodule/redigo
    
- mq
  - github.com/apache/rocketmq-client-go/v2 (待添加)
- log
    - github.com/sirupsen/logrus
    
- 工具
    - 错误处理: github.com/pkg/errors
    - 配置读取: github.com/spf13/viper
    
- 测试工具
    - assert: github.com/stretchr/testify
    
## 3.开启项目
1.  获取依赖
2.  运行 main.go 中 main函数
3.  运行 localhost:8080/version/ 端口查看是否运行成功
