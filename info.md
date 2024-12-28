1.你是一位资历深厚的Golang后端框架工程师，现在需要进行一个ai项目的开发，你将负责项目的架构设计，代码编写，以及项目部署。
2.现在需要做一个简单的聊天机器人，用户可以输入文本，机器人可以回复文本。
3.项目需要使用到以下技术栈：
    - 使用Golang语言
4.使用豆包大模型，使用http协议进行通信
5.做到企业级部署模板，包括但不限于：函数模块化，将service和config分开打包，并且将返回的结果传入数据库，使用到queue队列，并设置二十秒的过期时间
6.将敏感的数据存放在.env文件中，比如：
豆包的
ARK_API_KEY="3495ee71-c097-4211-a7c0-b17cff2bf340"
ENDPOINT_ID="ep-20241228172531-zjn6f"
BASE_URL="https://ark.cn-beijing.volces.com/api/v3"
数据库的host，port，password，user
7.单独存放一个prompt.json文件，里面存放system prompt，并且与user prompt进行组合
即为用户在前端请求时，将user prompt与system prompt进行组合，然后发送给后端
8.保证单一性原则，不要写多个无用文件，确保项目中只有一个main.go文件
9.项目模块化，将项目分为多个模块，每个模块负责一个功能，比如：
    - 用户管理模块
    - 聊天机器人模块
    - 数据库模块
    - 队列模块
    - 缓存模块
    - 日志模块
    - 配置模块
    - 监控模块
    - 部署模块
等
10.项目目录结构
    - main.go
    - config
    - service
    - bootstrap
    - pkg
    - routes
    - prompt
    - storage/logs
    - .env
