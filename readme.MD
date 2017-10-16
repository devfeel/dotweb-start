# DotWeb-Start
基于dotweb的示范项目结构

## 注意事项
* 使用配置文件方式，包含app.conf、dotweb.conf、dotlog.conf、dottask.conf
* bin目录为特定目录，建议在开发环境时，将编译生成的文件输出到该目录
* 确保bin目录下4个配置文件存在


## 目录说明
* config：全局配置文件定义及使用代码
* const：全局常量
* core：全局核心对象
* document：文档目录
* global：全局访问对象定义及入口
* protected：业务逻辑处理模块，包含数据模型、数据操作等
* resources：开发\测试\生产不同环境资源、配置文件
* server：web服务目录，包含常规处理逻辑、路由、静态资源等
* task：调度任务模块
* util：通用方法
