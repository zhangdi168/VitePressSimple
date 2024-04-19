
<div align="center">
<a href="https://github.com/zhangdi168/VitePressSimple">
<img src="./docs/vpstatic/images/vpsimple.png" width="120"/></a>
</div>
<h1 align="center">VitePress Simple</h1>

<h3 align="center"> VitePress Configuration Tool </h3>

<h3 align="center">
<a href="https://github.com/zhangdi168/VitePressSimple">English
</a>  | 
<strong>简体中文</strong>
</h3>


![vpsimple](./docs/vpstatic/images/demo.png)

# 功能特性
* 基于wails2+vite+vue3+typescript开发
* 基于Webview2的轻量客户端程序
* markdown文档编辑器在线编辑文档
* 支持文档目录树，支持移动、剪切、复制、粘贴等操作
* 支持直接将图片复制粘贴到markdown文档中并自动上传到服务器
* 【原生配置解析】直接对.vitepress/config.mts进行解析和渲染
* 【多语言】支持多语言配置,可以对不同语言进行分别配置
* 【导航配置】支持导航栏可视化配置，支持对语言分别配置
* 【侧栏配置】支持侧栏可视化配置，支持多侧栏配置，支持多语言，支持自动识别侧栏
* 【搜索】支持搜索配置，可选本地配置或搜索第三方服务
* 
# 快速使用指南

提供Mac、Windows安装包免费下载:
[github下载地址](https://github.com/zhangdi168
/VitePressSimple/releases) 
| [gitee下载地址](https://gitee.com/zhangdi168/VitePressSimple/releases)

前往官网查看完整文档：[中文文档](http://vpsimple.xiaod.co/zh) |
[英文文档](http://vpsimple.xiaod.co/en)

下面提供了一个示例vitepress项目，方便大家快速体验：
## 创建VitePress项目
![创建vitepress项目](./docs/vpstatic/images/20240416/9323bce8-7c90-439d-9b1b-49aec08211ea.png)
输入项目的名称和选择项目根目录：

![4631dcde70f7427bb5d07a2bd6d80b76.png](./docs/vpstatic/images/20240416/4631dcde-70f7-427b-b5d0-7a2bd6d80b76.png)
## 在线预览
![img.png](./docs/vpstatic/images/openInDir.png)

创建完成后可以**点击文件夹图标进入到根目录**（

```bash
# 进入到项目根目录，注意不是原目录
npm install
npm run dosc:dev
```
执行完上述命令后即可运行vitepress测试，
默认地址是：http://localhost:5173 ,
在vpsimple中修改文档内容或者配置即可实现同步更新。



## 构建项目

### 运行环境要求

* Go（最新版本）
* Node.js >= 16
* NPM >= 9

### 安装wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 拉取代码

```bash
git clone https://github.com/zhangdi168/VitePressSimple --depth=1
```

### 构建前端代码

```bash
npm install --prefix ./frontend
```

### 编译运行开发版本

```bash
wails dev
```




### 关于作者
[作者个人网站：http://xiaod.co/](http://xiaod.co/) 

该项目完全免费，如果对你有所帮助，可以请作者喝杯咖啡 ☕️



