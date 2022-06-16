# J2EE实验

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

## 目录

- [项目介绍](#项目介绍) 
  - [实验一](#实验一)
  - [实验二](#实验二)
  - [实验三](#实验三)
  - [实验四](#实验四)
  - [实验五](#实验五)
- [上手指南](#上手指南)
  - [开发前的配置要求](#开发前的配置要求)
  - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [技术栈](#技术栈)
- [贡献者](#贡献者)
  - [如何参与开源项目](#如何参与开源项目)
- [作者](#作者)
- [鸣谢](#鸣谢)


## 项目介绍

### 实验一
*基于CSS+HTML+JS开发简单个人网站*
* 个人网站分为四个部分：注册页面，登录界面，主界面，以及四个界面
* 另外我用Golang实现了网站后端，基于gin框架，将用户注册的账号密码储存在MySQL中
### 实验二
*基于Vue的图书购物车*
* 点击移除按钮，则移除一行记录的显示；
* 点击“+”或者“-”按钮，则购买数量发生相应的变化，且总价也随之发生变化；
* 当页面所有记录移除完，则在页面上显示“购物车为空”。
* 当购买数量为“0”时，“-”按钮变成不可使用。
### 实验三
*Vue-router的使用*
* 通过vue-cli 脚手架搭建 vue-cli 项目
* 使用Vue-router实现路由跳转和路由嵌套
### 实验四
*基于IDEA+SpringBoot+Mave+Thymeleaf的系统实现*
* 通过IDE创建springboot项目
 * pom.xml中引入thymeleaf依赖，在classpath:/templates/下新建HTML页面(采用实验一创建的页面)，新建controller进行测试，访问页面后返回success成功数据
### 实验五
*SpringBoot整合Redis点赞功能的简单实现*
* 通过给一篇文章点赞，点赞成功+1,取消点赞-1，数据都放在reids缓存里
* 引入Springboot整合redis的依赖，并在配置文件中配置redis相关属性
  
```
spring.redis.database=0
spring.redis.host=127.0.0.1
spring.redis.port=6379
spring.redis.password=
spring.redis.jedis.pool.max-active=8
spring.redis.jedis.pool.max-wait=-1
spring.redis.jedis.pool.max-idle=10
spring.redis.jedis.pool.min-idle=2
spring.redis.timeout=6000
  
```


## 上手指南


### 开发前的配置要求

1. node.js
2. java
3. maven
4. nvm
5. vue

### **安装步骤**

1. git clone此项目

```sh
git clone https://github.com/RobKing9/J2EE.git
```
&nbsp; &nbsp; &nbsp;2.在releases直接下载源代码

## 文件目录说明
```
J2EE 
.
|-- Experiment1
|   |-- common
|   |   `-- jwt.go
|   |-- controler
|   |   `-- controler.go
|   |-- dao
|   |   `-- dao.go
|   |-- go.mod
|   |-- go.sum
|   |-- main.go
|   |-- middleware
|   |   `-- authmiddleware.go
|   |-- models
|   |   `-- user.go
|   |-- response
|   |   `-- response.go
|   |-- router
|   |   `-- router.go
|   |-- static
|   |   |-- css
|   |   |   |-- index.css
|   |   |   |-- learn.css
|   |   |   |-- login.css
|   |   |   |-- mydream.css
|   |   |   |-- mylife.css
|   |   |   |-- register.css
|   |   |   `-- resume.css
|   |   |-- images
|   |   |   |-- 001.jpg
|   |   `-- js
|   |       `-- register.js
|   |-- template
|   |   |-- index.html
|   |   |-- learn.html
|   |   |-- login.html
|   |   |-- mydream.html
|   |   |-- mylife.html
|   |   |-- register.html
|   |   `-- resume.html
|   `-- tmp
|       |-- runner-build.exe
|       `-- runner-build.exe~
|-- Experiment2
|   |-- index.html
|   |-- main.css
|   `-- main.js
|-- Experiment3
|   `-- vue-project
|       |-- README.md
|       |-- build
|       |   |-- build.js
|       |   |-- check-versions.js
|       |   |-- logo.png
|       |   |-- utils.js
|       |   |-- vue-loader.conf.js
|       |   |-- webpack.base.conf.js
|       |   |-- webpack.dev.conf.js
|       |   `-- webpack.prod.conf.js
|       |-- config
|       |   |-- dev.env.js
|       |   |-- index.js
|       |   `-- prod.env.js
|       |-- index.html
|       |-- package-lock.json
|       |-- package.json
|       |-- src
|       |   |-- App.vue
|       |   |-- assets
|       |   |   `-- logo.png
|       |   |-- components
|       |   |   |-- HelloWorld.vue
|       |   |   |-- Home.vue
|       |   |   |-- Index.vue
|       |   |   |-- Login.vue
|       |   |   `-- Register.vue
|       |   |-- images
|       |   |   |-- login.png
|       |   |   `-- register.png
|       |   |-- main.js
|       |   `-- router
|       |       `-- index.js
|       `-- static
|-- Experiment4&5
|   `-- test
|       |-- HELP.md
|       |-- mvnw
|       |-- mvnw.cmd
|       |-- pom.xml
|       |-- src
|       |   |-- main
|       |   |   |-- java
|       |   |   `-- resources
|       |   `-- test
|       |       `-- java
|       |-- target
|       |   |-- classes
|       |   |   |-- application.properties
|       |   |   |-- static
|       |   |   |-- templates
|       |   |   `-- test
|       |   |-- generated-sources
|       |   |   `-- annotations
|       |   |-- generated-test-sources
|       |   |   `-- test-annotations
|       |   `-- test-classes
|       |       `-- test
|       `-- test.iml
|-- README.md
|-- images
|   |-- exp1_index.png
|-- tree.txt
`-- 实验报告
    |-- 实验一.docx
    |-- 实验二 .docx
    |-- 实验三.docx
    |-- 实验四.docx
    |-- 实验五.docx




```

### 技术栈
- [gin](https://gin-gonic.com/zh-cn/)
- [vue](https://cn.vuejs.org/)
- [springboot](https://spring.io/projects/spring-boot)
- [redis](https://redis.io/)

### 贡献者

请阅读**README.md** 查阅为该项目做出贡献的开发者。

### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


### 作者

QQ邮箱：2768817839@qq.com


 *您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt)

### 鸣谢


- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)

<!-- links -->
[your-project-path]:shaojintian/Best_README_template
[contributors-shield]: https://img.shields.io/github/contributors/RobKing9/J2EE.svg?style=flat-square
[contributors-url]: https://github.com/RobKing9/J2EE/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/RobKing9/J2EE.svg?style=flat-square
[forks-url]: https://github.com/RobKing9/J2EE/network/members
[stars-shield]: https://img.shields.io/github/stars/RobKing9/J2EE.svg?style=flat-square
[stars-url]: https://github.com/RobKing9/J2EE/stargazers
[issues-shield]: https://img.shields.io/github/issues/RobKing9/J2EE.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg
[license-shield]: https://img.shields.io/github/license/RobKing9/J2EE.svg?style=flat-square
[license-url]: https://github.com/RobKing9/J2EE/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/RobKing