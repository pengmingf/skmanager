# GFast-V3
gf23 build ./main.go -n admin-http -a amd64 -s linux -p ./bin -version http



## 平台简介
* 基于全新Go Frame 2.3+Vue3+Element Plus开发的全栈前后端分离的管理系统
* 前端采用vue-next-admin 、Vue、Element UI。
* 本项目由<a href="http://www.qjit.cn/" target="_blank"><b>奇讯科技</b></a>团队开发。
* 如您想购买云服务器请先领取阿里云优惠券：[点我进入](https://www.aliyun.com/minisite/goods?userCode=fcor2omk )，腾讯云优惠券：[点我领取](https://cloud.tencent.com/act/cps/redirect?redirect=1062&cps_key=20b1c3842f74986b2894e2c5fcde7ea2&from=console )
* 目前我们基于gfast开发了：吸vo(￣▽￣)ｄ毒人员风控平台、政府数据化招商平台、档案资源普查信息系统、扶贫资产管理系统、厂房消防安全风控平台、重点项目督办平台等。

## 特征
* 高生产率：几分钟即可搭建一个后台管理系统
* 模块化：单应用多系统的模式，将一个完整的应用拆分为多个系统，后续扩展更加便捷，增加代码复用性。
* 插件化： 可通过插件的方式扩展系统功能
* 认证机制：采用gtoken的用户状态认证及casbin的权限认证
* 路由模式：得利于goframe2.0提供了规范化的路由注册方式,无需注解自动生成api文档 
* 面向接口开发 


## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 在线用户：当前系统中活跃用户状态监控。
11. 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
12. 代码生成：前后端代码的生成。
13. 服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
14. 在线构建器：拖动表单元素生成相应的HTML代码。
15. 文件上传,缓存标签等。

> 项目使用模块化与插件化的方式开发，目前正在升级的插件有：cms系统、万能模型、微信管理、工单系统、问答系统、工作流引擎...，同时也欢迎大家把写好的插件上传到商城，我们来帮您实现技术变现，获取收益！

> GFast开源以来得到了大家的很多支持，如果您愿意为GFast贡献代码或提供建议，请加微信：qixun007(备注：gfast)

## 演示地址
[http://v3.g-fast.cn/sys](http://v3.g-fast.cn/sys)
账号：demo  密码：123456
## 配置
项目数据库文件 `resource/data/db.sql` 创建数据库导入后修改配置 `manifest/config/config.yaml.bak` 复制改为`manifest/config/config.yaml`  

其中gfToken配置
```yaml
gfToken:
  cacheKey: "gfToken_"   #缓存前缀
  timeOut: 10800         #token超时时间（秒）
  maxRefresh: 5400       #token自动刷新时间（秒）
  multiLogin: true       #是否允许一个账号多人同时登录
  encryptKey: "49c54195e750b04e74a8429b17896586"    #加密key (32位)
  cacheModel: "redis"    #存储引擎 （memory使用内存|redis使用redis）
  excludePaths:          #排除不做登录验证的路由地址
    - "/api/v1/system/login"
```

## 文档地址
> [http://doc.g-fast.cn/docs/gfast32](http://doc.g-fast.cn/docs/gfast32)