# go-websocket
    简单的go websocket 并发编程，实时推送

# govendor go依赖包管理使用

    命令	            功能
    init	        初始化 vendor 目录
    list	        列出所有的依赖包
    add	            添加包到 vendor 目录，如 govendor add +external 添加所有外部包
    add PKG_PATH	添加指定的依赖包到 vendor 目录
    update	        从 $GOPATH 更新依赖包到 vendor 目录
    remove	        从 vendor 管理中删除依赖
    status	        列出所有缺失、过期和修改过的包
    fetch	        添加或更新包到本地 vendor 目录
    sync	        本地存在 vendor.json时拉取依赖包，存储在本项目的vendor目录下
    get	            类似 go get 目录，拉取依赖包到 vendor 目录
