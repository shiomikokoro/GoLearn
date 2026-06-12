0. 环境配置
    1. 配置原始bin目录文件
    2. 修改GOPATH目录
        1. 直接在环境变量中创建 `GOPATH` + 你想要的文件地址
        2. 如果控制台 `go env` 对不上，可以在上述文件地址中添加 `bin`、`pkg`、`src`三个文件夹。
1. 项目初始化
    1. 官方1.10之后引入了GoModule管理依赖
    2. 创建项目路径之后使用 `go mod init shiomi.com/module[本初为自己想要设计的module名]`
    3. 清理依赖 `go mod tidy`
    4. 获取版本依赖 `go get package@version`
    5. 疑惑
        1. `%GOPATH%/bin` 中的工具包md5不对？ - Go是下载下源码，然后再编译出来的“专属”exe
        2. 我开VPN了，为什么下载失败？ - 需要修改Go环境的HTTP/HTTPS配置，但不方便修改，顾不建议。
        3. 下载不下来 - 修改Go的代理配置 `o env -w GOPROXY=https://goproxy.cn,direct` （七牛云）
        4. 我怎么知道是不是病毒？
            1. 七牛云保证了不会修改任何源码，只进行缓存下载等。
            2. Go保证了首次使用工具会进行验证，Go环境中有 `GOSUMDB=sum.golang.org` 或 `GOSUMDB=sum.golang.google.cn`
            3. 那官方验证成功了吗？ - 可以查看 `%GOPATH%/pkg/sumdb/sum.golang.org` 是否正常生成了文件，生成了代表没问题。
            4. **保证工具包本身不是病毒**，建议获取工具包依赖时输入完整包名，保证不会下载到钓鱼包
2. 项目缓存
    1. `APPDATA/local/go[xx]` 一系列文件
        1. gopls缓存 就是 `gopls` 文件夹
        2. goimports _疑似时gopls的附属生成_
        3. go-build go文件测试等生成的缓存，也可使用 `go clean -cache` 更安全的清除
3. 项目运行
    1. 有多个文件时，在项目根目录输入 `go run .`
    2. 单个文件输入 `go run [文件名]`
4. go没有继承和多态？
    1. 准确来说，没有其他语言口中的继承和多态概念，但能实现类似方法，就是组合/嵌入、接口式实现多态
    2. 继承可以向上转型，类似于继承的组合/嵌入，不能向上转型，没有“student is person”的概念，比如
    `Person p = new Student(); //java没问题`
    `c:=Class{} s:=Student{} c=s //go报错，类型不一致`
    3. 多态，java或类似语言中父类可以直接调用子类的方法，而go需要使用接口实现，来达到多态效果
    `Person p = new Student();p.study();//study是Student的方法`
    `type Person interface{study()} var p Person = Student{} p.study()//Person是接口，student需要实现study方法`
5. go天然支持高并发，因为底层代码决定了。
6. go的最大诟病就是异常处理，因为没有异常捕获机制，导致没调一个函数都要接一下函数的error数据，戏称Error是go的一等公民。
    
