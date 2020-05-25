# GOSS - Server Status

  English Version : placeholder

## GOSS

  GOSS是Server Status的一个Golang版本实现，用于检测服务器及运行服务状态.

  从RBQ Proj.分离下来的一个模块，目前没有做详细测试，基本能用.

## 功能

- 监测CPU使用率/内存使用/硬盘使用/网络IO使用/带宽使用
- 监测指定服务运行状态

![demo](.\doc\demo.png)

## 部署

  Linux和Windows版本部署大致流程相同

  Step 1.在Release中下载系统对应版本的服务端

  Step 2.在服务器上首次运行，修改配置文件`./goss-server-config.json`(以下为服务端配置文件模板)

```json
{
    "secKey": "MUST CHANGE THIS!", // 通讯密钥, 自定义修改
    "port": 80,	// GOSS服务端的端口
	"enableFront": true, // 是否启用前端(暂时没用)
	"enableBuiltInFront": true, // 是否启用内置前端(iViewUI的简单前端实现)
	"enableDatabase": false, // 是否启用数据记录(暂时没用)
	"enableCache": true, // 是否启用服务器状态结果的缓存
	"cacheLife": 3, // 缓存时间, 秒
	"enableHTTPS": false, // 是否启用HTTPS
	"certFile": "Your cert file path(You should have permission)", // 证书文件路径(GOSS服务端需要有权限访问)
	"keyFile": "Your key file path(You should have permission)" // Key文件路径(GOSS服务端需要有权限访问)
}
```

  Step 3.修改内置前端的配置文件`./front/config.js`(不启用忽略此步)

```json
window.g = {
	apiUrl : "http://127.0.0.1:1551", // 后端API的地址
	timeDuration : 5, // 获取服务器状态的时间间隔, 秒
	timeLimit : 5, // 认为被监测服务器多久未更新状态为断连, 秒
}
```

  Step 4.在欲被监测的服务器上下载对应系统的客户端

  Step 5.首次运行修改配置文件，或者将已经修改过的配置文件放入直接启动

```json
{
    "secKey": "This is your server's secret key", // 通讯密钥, 自定义修改
    "name": "Client's name (UNIQUE)", // 客户端的名称, 唯一
	"description": "Anything....", // 客户端服务器的描述, 可以不写
    "reportServer": "http://example.com/report", // 服务端的API地址
	"reportInterval": 5, // 回报信息的频率, 秒
    "reportService": [
		"nginx",
        "process name",
		"Windows process should add .exe like this.exe"
    ] // 需要监测的服务, 在Linux上为进程名, Win上需要添加.exe后缀
}
```

  Step 6. 没有Step 6了.

## 魔改(二次开发)

​    作者天天摸鱼，所以就别指望我们能新添什么功能了.

### 目录描述

```
├─client 客户端源码
├─doc 文档所需
├─front 前端源码(Vue)
│  ├─public
│  └─src
│      ├─assets
│      ├─components
│      └─plugins
└─server 服务端源码
    └─front 生成好的内置前端
        ├─css
        ├─fonts
        ├─img
        └─js
```

### 前端开发

- 原有前端的修改
    - 直接摸`./front/src`的前端源码就行
- 开发个新前端
    - 自己摸一套前端，然后和后台API通讯就可以
    - 需要修改一下服务端里的CORS

### 显示IP

  如果没记错的话，嗯，没记错的话. 前端是有IP显示的，后端没有传数据，在服务端的`ClientReportInfo`里面定义一下IP，然后添加 :

```
reportInfo.ip := ctx.RemoteAddr()
```

  当然也可以等我不摸的时候完善一下这里，原来是有的，但是因为每个人(会有人用?)的部署环境和需求不一样，就删了.

### API通讯

- POST方法 客户端回报状态 Report `/report`

    ```json
    {
        SecKey             string          `json:"secKey"` // 通讯密钥
    	Name               string          `json:"name"` // 客户端名称
    	Description        string          `json:"description"` // 客户端描述
    	CpuUsage           int             `json:"cpuUsage"` // CPU使用率
    	MemUsage           int             `json:"memUsage"` // 内存使用
    	MaxMemLimit        int             `json:"maxMemLimit"` // 内存上限
    	DiskUsage          int             `json:"diskUsage"` // 硬盘使用
    	MaxDiskLimit       int             `json:"maxDiskLimit"` // 硬盘上限
    	InBound            float64         `json:"inBound"` // 当前网络传入(计算拟合得出)
    	InBoundTotalUsage  float64         `json:"inBoundTotalUsage"` // 总计网络输入
    	OutBound           float64         `json:"outBound"` // 当前网络传出(计算拟合得出)
    	OutBoundTotalUsage float64         `json:"outBoundTotalUsage"` // 总计网络传出
    	ServiceStatus      map[string]bool `json:"serviceStatus"` // 服务状态
    }
    ```

- POST方法 获取服务端中的客户端信息 Status `/status`

    ```json
    {
        status : [stauts]
    }
    ```

## 其他

基于`Iris V12`，`Viper`，`gopsutil`开发.

内置前端基于`Vuejs`，`iViewUI`.

摸了.

哦对，不建议部署在大型的生产环境中，玩玩就行，反正我只是把它拿来做梯子可用监测的.