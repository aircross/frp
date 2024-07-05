# frp

[![Build Status](https://circleci.com/gh/fatedier/frp.svg?style=shield)](https://circleci.com/gh/fatedier/frp)
[![GitHub release](https://img.shields.io/github/tag/fatedier/frp.svg?label=release)](https://github.com/aircross/frp/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/aircross/frp)](https://goreportcard.com/report/github.com/aircross/frp)
[![GitHub Releases Stats](https://img.shields.io/github/downloads/fatedier/frp/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=fatedier&repository=frp)

[README](https://github.com/aircross/frp/blob/dev/README_en.md) | [中文文档](https://github.com/aircross/frp/blob/dev/README.md)

### 增加支持客户端直接配置IP使用，解决了使用域名绑定国内IP会因为未备案而被阻断的问题

#### Docker快速部署

###### 安装Docker

```
    #国外服务器使用以下命令安装Docker
    curl -fsSL https://get.docker.com | sh
    # 设置开机自启
    sudo systemctl enable docker.service
    # 根据实际需要保留参数start|restart|stop
    sudo service docker start|restart|stop
```

国内的请参照下面这个教程安装，需要配合能访问download.docker.com的服务器服用

###### [和谐之后如何在国内安装Docker及拉取镜像使用](https://vps.la/2024/07/01/%e5%92%8c%e8%b0%90%e4%b9%8b%e5%90%8e%e5%a6%82%e4%bd%95%e5%9c%a8%e5%9b%bd%e5%86%85%e5%ae%89%e8%a3%85docker%e5%8f%8a%e6%8b%89%e5%8f%96%e9%95%9c%e5%83%8f%e4%bd%bf%e7%94%a8/)

###### 运行Frps的镜像

```
    mkdir -p /opt/docker/frp/
    
    tee /opt/docker/frp/frps.toml <<-'EOF'
    bindPort = 7000
    # 默认为 127.0.0.1，如果需要公网访问，需要修改为 0.0.0.0。
    webServer.addr = "0.0.0.0"
    webServer.port = 7500
    # dashboard 用户名密码，可选，默认为空
    webServer.user = "admin"
    webServer.password = "admin"
    # auth token
    auth.token = "your token Here"
    EOF
    docker run --name frps -d --network host --restart=unless-stopped -v /opt/docker/frp/frps.toml:/etc/frp/frps.toml  aircross/frps
```

###### 运行Frpc的镜像

```mkdir -p /opt/docker/frp/
    
    tee /opt/docker/frp/frpc.toml <<-'EOF'
    serverAddr = "your frps ip here"
    serverPort = 7000
    # auth token
    auth.token = "your token Here"
    
    [[proxies]]
    name = "Node1_SSH"
    type = "tcp"
    localIP = "127.0.0.1"
    localPort = 22
    remotePort = 40022
    EOF
    docker run --name frpc -d --network host --restart=unless-stopped -v /opt/docker/frp/frpc.toml:/etc/frp/frpc.toml  aircross/frpc
``` 

#### 推荐服务器

如果你觉得本项目对你有用,而且你也恰巧有这方面的需求,你也可以选择通过我的购买链接赞助我

-   [搬瓦工GIA高端线路](https://bandwagonhost.com/aff.php?aff=38140),仅推荐购买GIA套餐
-   [Spartan三网4837性价比主机](https://billing.spartanhost.net/aff.php?aff=1156)
-   [Dmit](https://www.dmit.io/aff.php?aff=9771)
-   [Linode](https://www.linode.com/lp/refer/?r=107a1afa3e657b37fc273df95803557588e7dcc5)
-   [Vultr](https://www.vultr.com/?ref=7130790)
-   [Cloudcone性价比主机提供商](https://app.cloudcone.com/?ref=2227)

frp 是一个专注于内网穿透的高性能的反向代理应用，支持 TCP、UDP、HTTP、HTTPS 等多种协议，且支持 P2P 通信。可以将内网服务以安全、便捷的方式通过具有公网 IP 节点的中转暴露到公网。

## 为什么使用 frp ？

通过在具有公网 IP 的节点上部署 frp 服务端，可以轻松地将内网服务穿透到公网，同时提供诸多专业的功能特性，这包括：

* 客户端服务端通信支持 TCP、QUIC、KCP 以及 Websocket 等多种协议。
* 采用 TCP 连接流式复用，在单个连接间承载更多请求，节省连接建立时间，降低请求延迟。
* 代理组间的负载均衡。
* 端口复用，多个服务通过同一个服务端端口暴露。
* 支持 P2P 通信，流量不经过服务器中转，充分利用带宽资源。
* 多个原生支持的客户端插件（静态文件查看，HTTPS/HTTP 协议转换，HTTP、SOCK5 代理等），便于独立使用 frp 客户端完成某些工作。
* 高度扩展性的服务端插件系统，易于结合自身需求进行功能扩展。
* 服务端和客户端 UI 页面。

## 开发状态

frp 目前已被很多公司广泛用于测试、生产环境。

master 分支用于发布稳定版本，dev 分支用于开发，您可以尝试下载最新的 release 版本进行测试。

我们正在进行 v2 大版本的开发，将会尝试在各个方面进行重构和升级，且不会与 v1 版本进行兼容，预计会持续较长的一段时间。

现在的 v0 版本将会在合适的时间切换为 v1 版本并且保证兼容性，后续只做 bug 修复和优化，不再进行大的功能性更新。

### 关于 v2 的一些说明

v2 版本的复杂度和难度比我们预期的要高得多。我只能利用零散的时间进行开发，而且由于上下文经常被打断，效率极低。由于这种情况可能会持续一段时间，我们仍然会在当前版本上进行一些优化和迭代，直到我们有更多空闲时间来推进大版本的重构，或者也有可能放弃一次性的重构，而是采用渐进的方式在当前版本上逐步做一些可能会导致不兼容的修改。

v2 的构想是基于我多年在云原生领域，特别是在 K8s 和 ServiceMesh 方面的工作经验和思考。它的核心是一个现代化的四层和七层代理，类似于 envoy。这个代理本身高度可扩展，不仅可以用于实现内网穿透的功能，还可以应用于更多领域。在这个高度可扩展的内核基础上，我们将实现 frp v1 中的所有功能，并且能够以一种更加优雅的方式实现原先架构中无法实现或不易实现的功能。同时，我们将保持高效的开发和迭代能力。

除此之外，我希望 frp 本身也成为一个高度可扩展的系统和平台，就像我们可以基于 K8s 提供一系列扩展能力一样。在 K8s 上，我们可以根据企业需求进行定制化开发，例如使用 CRD、controller 模式、webhook、CSI 和 CNI 等。在 frp v1 中，我们引入了服务端插件的概念，实现了一些简单的扩展性。但是，它实际上依赖于简单的 HTTP 协议，并且需要用户自己启动独立的进程和管理。这种方式远远不够灵活和方便，而且现实世界的需求千差万别，我们不能期望一个由少数人维护的非营利性开源项目能够满足所有人的需求。

最后，我们意识到像配置管理、权限验证、证书管理和管理 API 等模块的当前设计并不够现代化。尽管我们可能在 v1 版本中进行一些优化，但确保兼容性是一个令人头疼的问题，需要投入大量精力来解决。

非常感谢您对 frp 的支持。

## 文档

完整文档已经迁移至 [https://gofrp.org](https://gofrp.org)。

## 为 frp 做贡献

frp 是一个免费且开源的项目，我们欢迎任何人为其开发和进步贡献力量。

* 在使用过程中出现任何问题，可以通过 [issues](https://github.com/aircross/frp/issues) 来反馈。
* Bug 的修复可以直接提交 Pull Request 到 dev 分支。
* 如果是增加新的功能特性，请先创建一个 issue 并做简单描述以及大致的实现方法，提议被采纳后，就可以创建一个实现新特性的 Pull Request。
* 欢迎对说明文档做出改善，帮助更多的人使用 frp，特别是英文文档。
* 贡献代码请提交 PR 至 dev 分支，master 分支仅用于发布稳定可用版本。
* 如果你有任何其他方面的问题或合作，欢迎发送邮件至 fatedier@gmail.com 。

**提醒：和项目相关的问题请在 [issues](https://github.com/aircross/frp/issues) 中反馈，这样方便其他有类似问题的人可以快速查找解决方法，并且也避免了我们重复回答一些问题。**

## 关联项目

* [gofrp/plugin](https://github.com/gofrp/plugin) - frp 插件仓库，收录了基于 frp 扩展机制实现的各种插件，满足各种场景下的定制化需求。
* [gofrp/tiny-frpc](https://github.com/gofrp/tiny-frpc) - 基于 ssh 协议实现的 frp 客户端的精简版本(最低约 3.5MB 左右)，支持常用的部分功能，适用于资源有限的设备。

## 赞助

如果您觉得 frp 对你有帮助，欢迎给予我们一定的捐助来维持项目的长期发展。

### Sponsors

长期赞助可以帮助我们保持项目的持续发展。

您可以通过 [GitHub Sponsors](https://github.com/sponsors/fatedier) 赞助我们。

国内用户可以通过 [爱发电](https://afdian.net/a/fatedier) 赞助我们。

企业赞助者可以将贵公司的 Logo 以及链接放置在项目 README 文件中。

### 知识星球

如果您想了解更多 frp 相关技术以及更新详解，或者寻求任何 frp 使用方面的帮助，都可以通过微信扫描下方的二维码付费加入知识星球的官方社群：

![zsxq](/doc/pic/zsxq.jpg)
