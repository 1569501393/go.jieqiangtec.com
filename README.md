[TOC]
# [go.jieqiangtec.com](https://www.jb51.net/books/709772.html)
go学习

# [Go 语言学习资料与社区索引](https://gobyexample.com/)

# [Go 学习之路](https://github.com/1569501393/learning-golang)

# [Go 语言设计模式](https://github.com/1569501393/golang-design-pattern)

# [go-study-index](https://github.com/1569501393/go-study-index)

# [Golang资料补给包](https://github.com/1569501393/Introduction-to-Golang)

# [Go 开发者路线图](https://github.com/Quorafind/golang-developer-roadmap-cn)

# [Go 入门的学习笔记](https://github.com/1569501393/Go-1)

## [**Go入门指南——The Way to Go（中文版）**](https://m.ituring.com.cn/book/1205)

## [Go 入门指南](https://learnku.com/docs/the-way-to-go)
```
基础信息
 镜像说明
 前言
 第一章. Go 起源，发展与普及
 1.1. 起源与发展
 1.2. 主要特性与发展的环境和影响因素
 第二章. 安装与运行环境
 2.1. 平台与架构
 2.2. Go 环境变量
 2.3. 在 Linux 上安装 Go
 2.4. 在 Mac OS x 上安装 Go
 2.5. 在 Windows 上安装 Go
 2.6. 安装目录清单
 2.7. Go 运行时（runtime）
 2.8. Go 解释器
 第三章. 编辑器、集成开发环境与其它工具
 章节说明
 3.1. Go 开发环境的基本要求
 3.2. 编辑器和集成开发环境
 3.3. 调试器
 3.4. 构建并运行 Go 程序
 3.5. 格式化代码
 3.6. 生成代码文档
 3.7. 其它工具
 3.8. Go 性能说明
 3.9. 与其它语言进行交互
 第四章. 基本结构和基本数据类型
 4.1. 文件名、关键字与标识符
 4.2. Go 程序的基本结构和要素
 4.3. 常量
 4.4. 变量
 4.5. 基本类型和运算符
 4.6. 字符串
 4.7. strings 和 strconv 包
 4.8. 时间和日期
 4.9. 指针
 第五章. 控制结构
 章节说明
 5.1. if-else 结构
 5.2. 测试多返回值函数的错误
 5.3. switch 结构
 5.4. for 结构
 5.5. Break 与 continue
 5.6. 标签与 goto
 第六章. 函数（function）
 章节说明
 6.1. 介绍
 6.2. 函数参数与返回值
 6.3. 传递变长参数
 6.4. defer 和追踪
 6.5. 内置函数
 6.6. 递归函数
 6.7. 将函数作为参数
 6.8. 闭包
 6.9. 应用闭包：将函数作为返回值
 6.10. 使用闭包调试
 6.11. 计算函数执行时间
 6.12. 通过内存缓存来提升性能
 第七章. 数组与切片
 章节说明
 7.1. 声明和初始化
 7.2. 切片
 7.3. For-range 结构
 7.4. 切片重组（reslice）
 7.5. 切片的复制与追加
 7.6. 字符串、数组和切片的应用
 第八章. Map
 章节说明
 8.1. 声明、初始化和 make
 8.2. 测试键值对是否存在及删除元素
 8.3. for-range 的配套用法
 8.4. map 类型的切片
 8.5. map 的排序
 8.6. 将 map 的键值对调
 第九章. 包（package）
 章节说明
 9.1. 标准库概述
 9.2. regexp 包
 9.3. 锁和 sync 包
 9.4. 精密计算和 big 包
 9.5. 自定义包和可见性
 9.6. 为自定义包使用 godoc
 9.7. 使用 go install 安装自定义包
 9.8. 自定义包的目录结构、go install 和 go test
 9.9. 通过 Git 打包和安装
 9.10. Go 的外部包和项目
 9.11. 在 Go 程序中使用外部库
 第十章. 结构（struct）与方法（method）
 章节说明
 10.1. 结构体定义
 10.2. 使用工厂方法创建结构体实例
 10.3. 使用自定义包中的结构体
 10.4. 带标签的结构体
 10.5. 匿名字段和内嵌结构体
 10.6. 方法
 10.7. 类型的 String() 方法和格式化描述符
 10.8. 垃圾回收和 SetFinalizer
 第十一章. 接口（interface）与反射（reflection）
 章节说明
 11.1. 接口是什么
 11.2. 接口嵌套接口
 11.3. 类型断言：如何检测和转换接口变量的类型
 11.4. 类型判断：type-switch
 11.5. 测试一个值是否实现了某个接口
 11.6. 使用方法集与接口
 11.7. 第一个例子：使用 Sorter 接口排序
 11.8. 第二个例子：读和写
 11.9. 空接口
 11.10 反射包
 11.12. 接口与动态类型
 11.13. 总结：Go 中的面向对象
 11.14. 结构体、集合和高阶函数
 第十二章. 读写数据
 章节说明
 12.1. 读取用户的输入
 12.2. 文件读写
 12.3. 文件拷贝
 12.4. 从命令行读取参数
 12.5. 用 buffer 读取文件
 12.6. 用切片读写文件
 12.7. 用 defer 关闭文件
 12.8. 使用接口的实际例子：fmt.Fprintf
 12.9. JSON 数据格式
 12.10. XML 数据格式
 12.11. 用 Gob 传输数据
 12.12. Go 中的密码学
 第十三章. 错误处理与测试
 章节说明
 13.1. 错误处理
 13.2. 运行时异常和 panic
 13.3. 从 panic 中恢复（Recover）
 13.4. 自定义包中的错误处理和 panicking
 13.5. 一种用闭包处理错误的模式
 13.6. 启动外部命令和程序
 13.7. Go 中的单元测试和基准测试
 13.8. 测试的具体例子
 13.9. 用（测试数据）表驱动测试
 13.10. 性能调试：分析并优化 Go 程序
 第十四章. 协程（goroutine）与通道（channel）
 章节说明
 14.1. 并发、并行和协程
 14.2. 协程间的信道
 14.3. 协程的同步：关闭通道-测试阻塞的通道
 14.4. 使用 select 切换协程
 14.5. 通道、超时和计时器（Ticker）
 14.6. 协程和恢复（recover）
 14.7. 新旧模型对比：任务和worker
 14.8. 惰性生成器的实现
 14.9. 实现 Futures 模式
 14.10. 多路复用 已完成
 14.11. 限制并发数 已完成
 14.12. 链式操作 已完成
 14.13. 多核运算 已完成
 14.14. 多核运算处理大量数据 已完成
 14.15. 漏桶算法 Leaky Bucket 已完成
 14.16. 标杆分析 Goroutines 已完成
 14.17. 使用 Channel 来并发读取对象 已完成
 第十五章. 网络，模板和网页应用
 章节说明
 15.1. tcp 服务器
 15.2. 一个简单的网页服务器
 15.3. 访问并读取页面
 15.4. 写一个简单的网页应用
 15.5. 让 Web 应用更加健壮 已完成
 15.6. 在 Web 应用中使用模板 已完成
 15.7. 探索 Template 扩展的功能 已完成
 15.8. 一个多功能的精致的 WebServer 已完成
 15.9. RPC 远程调用 已完成
 15.10. 使用 netchan 跨网络实现消息传递 已完成
 15.11. Websocket 通讯 已完成
 15.12. SMTP 发送邮件 已完成
 第十六章. 常见的陷阱与错误
 章节说明
 16.1. 误用短声明导致变量覆盖
 16.2. 误用字符串
 16.3 发生错误时使用defer关闭一个文件
 16.4. 何时使用 new() 和 make()
 16.5. 不需要将一个指向切片的指针传递给函数
 16.6. 使用指针指向接口类型
 16.7. 使用值类型时误用指针
 16.8 误用协程和通道
 16.9. 闭包和协程的使用
 16.10. 糟糕的错误处理
 第十七章. 模式
 17.1. 关于逗号 ok 模式
 17.2. defer 模式 已完成
 17.3.能见度模式 已完成
 17.4. 操作者模式和接口 已完成
 第十八章. 出于性能考虑的实用代码片段
 18.1. 字符串
 18.2. 数组和切片
 18.3. 映射
 18.4. 结构体
 18.5. 接口
 18.6. 函数
 18.7. 文件
 18.8. 协程（goroutine）与通道（channel）
 18.9. 网络和网页应用
 18.10. 其他
 18.11. 出于性能考虑的最佳实践和建议
 第十九章. 构建完整的应用程序
 19.1. 简介 已完成
 19.2. UrlShortener 项目介绍 已完成
 19.3. 数据结构分析 已完成
 19.4. 用户界面：Web 网页前端 已完成
 19.5. 数据存储: gob 已完成
 19.6. 使用 Goroutines 来提高性能 已完成
 19.7. 使用 Json 来存储 已完成
 19.8. 多台机器上的多线程 已完成
 19.9. 使用 ProxyStore 已完成
 19.10. 总结和优化 已完成
 第二十章. Go 在 Google App 引擎中的使用
 章节说明
 第二十一章. Go 在现实世界中的使用
 章节说明 已完成
 21.2. MROffice — Go 实现的 VOIP 系统 已完成
 21.3. Atlassian— 虚拟机集群管理系统 已完成
 21.4. Camlistore 个人住址存储系统 已完成
```

1. 基础信息

   1.  [镜像说明](https://learnku.com/docs/the-way-to-go/book-intro/3560)
   2.  [前言](https://learnku.com/docs/the-way-to-go/preface/3561)

2. 

    

   第一章. Go 起源，发展与普及

   1.  [1.1. 起源与发展](https://learnku.com/docs/the-way-to-go/origin-and-development/3562)
   2.  [1.2. 主要特性与发展的环境和影响因素](https://learnku.com/docs/the-way-to-go/characteristics-environment-and-factors/3563)

3. 

    

   第二章. 安装与运行环境

   1.  [2.1. 平台与架构](https://learnku.com/docs/the-way-to-go/31-platform-and-architecture/3564)
   2.  [2.2. Go 环境变量](https://learnku.com/docs/the-way-to-go/go-environment-variable/3565)
   3.  [2.3. 在 Linux 上安装 Go](https://learnku.com/docs/the-way-to-go/install-go-on-linux/3566)
   4.  [2.4. 在 Mac OS x 上安装 Go](https://learnku.com/docs/the-way-to-go/install-go-on-mac-os-x/3567)
   5.  [2.5. 在 Windows 上安装 Go](https://learnku.com/docs/the-way-to-go/install-go-on-windows/3568)
   6.  [2.6. 安装目录清单](https://learnku.com/docs/the-way-to-go/install-directory-listing/3569)
   7.  [2.7. Go 运行时（runtime）](https://learnku.com/docs/the-way-to-go/go-runtime-runtime/3570)
   8.  [2.8. Go 解释器](https://learnku.com/docs/the-way-to-go/go-interpreter/3571)

4. 

    

   第三章. 编辑器、集成开发环境与其它工具

   1.  [章节说明](https://learnku.com/docs/the-way-to-go/editors-integrated-development-environments-and-other-tools/3572)
   2.  [3.1. Go 开发环境的基本要求](https://learnku.com/docs/the-way-to-go/basic-requirements-for-go-development-environment/3573)
   3.  [3.2. 编辑器和集成开发环境](https://learnku.com/docs/the-way-to-go/editors-and-integrated-development-environments/3574)
   4.  [3.3. 调试器](https://learnku.com/docs/the-way-to-go/33-debugger/3575)
   5.  [3.4. 构建并运行 Go 程序](https://learnku.com/docs/the-way-to-go/build-and-run-go-programs/3576)
   6.  [3.5. 格式化代码](https://learnku.com/docs/the-way-to-go/formatting-code/3577)
   7.  [3.6. 生成代码文档](https://learnku.com/docs/the-way-to-go/generating-code-documents/3578)
   8.  [3.7. 其它工具](https://learnku.com/docs/the-way-to-go/other-tools/3579)
   9.  [3.8. Go 性能说明](https://learnku.com/docs/the-way-to-go/go-performance-description/3580)
   10.  [3.9. 与其它语言进行交互](https://learnku.com/docs/the-way-to-go/interact-with-other-languages/3581)

5. 

    

   第四章. 基本结构和基本数据类型

   1.  [4.1. 文件名、关键字与标识符](https://learnku.com/docs/the-way-to-go/file-name-keyword-and-identifier/3582)
   2.  [4.2. Go 程序的基本结构和要素](https://learnku.com/docs/the-way-to-go/the-basic-structure-and-elements-of-the-go-program/3583)
   3.  [4.3. 常量](https://learnku.com/docs/the-way-to-go/constant/3584)
   4.  [4.4. 变量](https://learnku.com/docs/the-way-to-go/variable/3585)
   5.  [4.5. 基本类型和运算符](https://learnku.com/docs/the-way-to-go/basic-types-and-operators/3586)
   6.  [4.6. 字符串](https://learnku.com/docs/the-way-to-go/character-string/3587)
   7.  [4.7. strings 和 strconv 包](https://learnku.com/docs/the-way-to-go/strings-and-strconv-packages/3588)
   8.  [4.8. 时间和日期](https://learnku.com/docs/the-way-to-go/time-and-date/3589)
   9.  [4.9. 指针](https://learnku.com/docs/the-way-to-go/pointer/3590)

6. 

    

   第五章. 控制结构

   1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3591)
   2.  [5.1. if-else 结构](https://learnku.com/docs/the-way-to-go/if-else-structure/3592)
   3.  [5.2. 测试多返回值函数的错误](https://learnku.com/docs/the-way-to-go/test-the-error-of-multiple-return-function/3593)
   4.  [5.3. switch 结构](https://learnku.com/docs/the-way-to-go/switch-structure/3594)
   5.  [5.4. for 结构](https://learnku.com/docs/the-way-to-go/for-structure/3595)
   6.  [5.5. Break 与 continue](https://learnku.com/docs/the-way-to-go/break-and-continue/3596)
   7.  [5.6. 标签与 goto](https://learnku.com/docs/the-way-to-go/label-and-goto/3597)

7. 

    

   第六章. 函数（function）

   1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3598)
   2.  [6.1. 介绍](https://learnku.com/docs/the-way-to-go/introduce/3599)
   3.  [6.2. 函数参数与返回值](https://learnku.com/docs/the-way-to-go/function-parameters-and-return-values/3600)
   4.  [6.3. 传递变长参数](https://learnku.com/docs/the-way-to-go/transfer-length-parameter/3601)
   5.  [6.4. defer 和追踪](https://learnku.com/docs/the-way-to-go/defer-and-tracking/3602)
   6.  [6.5. 内置函数](https://learnku.com/docs/the-way-to-go/built-in-function/3603)
   7.  [6.6. 递归函数](https://learnku.com/docs/the-way-to-go/recursive-function/3604)
   8.  [6.7. 将函数作为参数](https://learnku.com/docs/the-way-to-go/take-the-function-as-a-parameter/3605)
   9.  [6.8. 闭包](https://learnku.com/docs/the-way-to-go/closure/3606)
   10.  [6.9. 应用闭包：将函数作为返回值](https://learnku.com/docs/the-way-to-go/application-closure-function-as-a-return-value/3607)
   11.  [6.10. 使用闭包调试](https://learnku.com/docs/the-way-to-go/use-closure-debugging/3608)
   12.  [6.11. 计算函数执行时间](https://learnku.com/docs/the-way-to-go/computing-function-execution-time/3609)
   13.  [6.12. 通过内存缓存来提升性能](https://learnku.com/docs/the-way-to-go/improve-performance-through-memory-caching/3610)

8. 

    

   第七章. 数组与切片

   1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3611)
   2.  [7.1. 声明和初始化](https://learnku.com/docs/the-way-to-go/declarations-and-initialization/3612)
   3.  [7.2. 切片](https://learnku.com/docs/the-way-to-go/72-slice/3613)
   4.  [7.3. For-range 结构](https://learnku.com/docs/the-way-to-go/73-for-range-structure/3614)
   5.  [7.4. 切片重组（reslice）](https://learnku.com/docs/the-way-to-go/74-slice-recombination-reslice/3615)
   6.  [7.5. 切片的复制与追加](https://learnku.com/docs/the-way-to-go/duplication-and-addition-of-75-slices/3616)
   7.  [7.6. 字符串、数组和切片的应用](https://learnku.com/docs/the-way-to-go/the-application-of-76-strings-arrays-and-slices/3617)

9. 

    

   第八章. Map

   1.  [章节说明](https://learnku.com/docs/the-way-to-go/8-chapters/3618)
   2.  [8.1. 声明、初始化和 make](https://learnku.com/docs/the-way-to-go/81-declarations-initialization-and-make/3619)
   3.  [8.2. 测试键值对是否存在及删除元素](https://learnku.com/docs/the-way-to-go/82-test-key-values-for-the-existence-and-deletion-of-elements/3620)
   4.  [8.3. for-range 的配套用法](https://learnku.com/docs/the-way-to-go/the-matching-usage-of-83-for-range/3621)
   5.  [8.4. map 类型的切片](https://learnku.com/docs/the-way-to-go/slicing-of-the-84-map-type/3622)
   6.  [8.5. map 的排序](https://learnku.com/docs/the-way-to-go/the-sorting-of-85-map/3623)
   7.  [8.6. 将 map 的键值对调](https://learnku.com/docs/the-way-to-go/86-adjust-the-key-values-of-map/3624)

10. 

     

    第九章. 包（package）

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3625)
    2.  [9.1. 标准库概述](https://learnku.com/docs/the-way-to-go/overview-of-the-91-standard-library/3626)
    3.  [9.2. regexp 包](https://learnku.com/docs/the-way-to-go/93-locks-and-sync-packages/3627)
    4.  [9.3. 锁和 sync 包](https://learnku.com/docs/the-way-to-go/94-precision-calculation-and-big-package/3628)
    5.  [9.4. 精密计算和 big 包](https://learnku.com/docs/the-way-to-go/95-custom-packages-and-visibility/3629)
    6.  [9.5. 自定义包和可见性](https://learnku.com/docs/the-way-to-go/92-regexp-package/3630)
    7.  [9.6. 为自定义包使用 godoc](https://learnku.com/docs/the-way-to-go/96-uses-godoc-for-custom-packages/3631)
    8.  [9.7. 使用 go install 安装自定义包](https://learnku.com/docs/the-way-to-go/97-installs-custom-packages-using-go-install/3632)
    9.  [9.8. 自定义包的目录结构、go install 和 go test](https://learnku.com/docs/the-way-to-go/98-custom-package-directory-structure-go-install-and-go-test/3633)
    10.  [9.9. 通过 Git 打包和安装](https://learnku.com/docs/the-way-to-go/99-is-packaged-and-installed-through-git/3634)
    11.  [9.10. Go 的外部包和项目](https://learnku.com/docs/the-way-to-go/external-packages-and-projects-of-910-go/3635)
    12.  [9.11. 在 Go 程序中使用外部库](https://learnku.com/docs/the-way-to-go/911-uses-an-external-library-in-the-go-program/3636)

11. 

     

    第十章. 结构（struct）与方法（method）

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3637)
    2.  [10.1. 结构体定义](https://learnku.com/docs/the-way-to-go/101-structure-definition/3638)
    3.  [10.2. 使用工厂方法创建结构体实例](https://learnku.com/docs/the-way-to-go/102-creates-an-instance-of-a-structure-using-a-factory-method/3639)
    4.  [10.3. 使用自定义包中的结构体](https://learnku.com/docs/the-way-to-go/103-uses-the-structure-in-the-custom-package/3640)
    5.  [10.4. 带标签的结构体](https://learnku.com/docs/the-way-to-go/107-type-string-method-and-formatted-descriptor/3641)
    6.  [10.5. 匿名字段和内嵌结构体](https://learnku.com/docs/the-way-to-go/104-tag-structure/3642)
    7.  [10.6. 方法](https://learnku.com/docs/the-way-to-go/105-anonymous-field-and-embedded-structure/3643)
    8.  [10.7. 类型的 String() 方法和格式化描述符](https://learnku.com/docs/the-way-to-go/106-method/3644)
    9.  [10.8. 垃圾回收和 SetFinalizer](https://learnku.com/docs/the-way-to-go/108-garbage-collection-and-setfinalizer/3645)

12. 

     

    第十一章. 接口（interface）与反射（reflection）

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3646)
    2.  [11.1. 接口是什么](https://learnku.com/docs/the-way-to-go/what-is-the-111-interface/3647)
    3.  [11.2. 接口嵌套接口](https://learnku.com/docs/the-way-to-go/112-interface-nesting-interface/3648)
    4.  [11.3. 类型断言：如何检测和转换接口变量的类型](https://learnku.com/docs/the-way-to-go/113-type-assertion-how-to-detect-and-transform-the-type-of-interface-variables/3649)
    5.  [11.4. 类型判断：type-switch](https://learnku.com/docs/the-way-to-go/114-type-judgment-type-switch/3650)
    6.  [11.5. 测试一个值是否实现了某个接口](https://learnku.com/docs/the-way-to-go/115-tests-whether-a-value-implements-an-interface/3651)
    7.  [11.6. 使用方法集与接口](https://learnku.com/docs/the-way-to-go/116-usage-set-and-interface/3652)
    8.  [11.7. 第一个例子：使用 Sorter 接口排序](https://learnku.com/docs/the-way-to-go/the-first-example-of-117-sorting-using-the-sorter-interface/3653)
    9.  [11.8. 第二个例子：读和写](https://learnku.com/docs/the-way-to-go/118-second-examples-reading-and-writing/3654)
    10.  [11.9. 空接口](https://learnku.com/docs/the-way-to-go/119-empty-interface/3655)
    11.  [11.10 反射包](https://learnku.com/docs/the-way-to-go/1110-reflector/3656)
    12.  [11.12. 接口与动态类型](https://learnku.com/docs/the-way-to-go/1112-interface-and-dynamic-type/3657)
    13.  [11.13. 总结：Go 中的面向对象](https://learnku.com/docs/the-way-to-go/1113-summary-object-oriented-in-go/3658)
    14.  [11.14. 结构体、集合和高阶函数](https://learnku.com/docs/the-way-to-go/1114-structures-sets-and-higher-order-functions/3659)

13. 

     

    第十二章. 读写数据

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3660)
    2.  [12.1. 读取用户的输入](https://learnku.com/docs/the-way-to-go/121-reads-user-input/3661)
    3.  [12.2. 文件读写](https://learnku.com/docs/the-way-to-go/122-file-reading-and-writing/3662)
    4.  [12.3. 文件拷贝](https://learnku.com/docs/the-way-to-go/copy-of-123-file/3663)
    5.  [12.4. 从命令行读取参数](https://learnku.com/docs/the-way-to-go/124-reads-parameters-from-the-command-line/3664)
    6.  [12.5. 用 buffer 读取文件](https://learnku.com/docs/the-way-to-go/125-reads-files-with-buffer/3665)
    7.  [12.6. 用切片读写文件](https://learnku.com/docs/the-way-to-go/126-read-and-write-files-with-sliced-slices/3666)
    8.  [12.7. 用 defer 关闭文件](https://learnku.com/docs/the-way-to-go/127-closes-files-with-defer/3667)
    9.  [12.8. 使用接口的实际例子：fmt.Fprintf](https://learnku.com/docs/the-way-to-go/an-actual-example-of-128-using-the-interface-fmtfprintf/3668)
    10.  [12.9. JSON 数据格式](https://learnku.com/docs/the-way-to-go/129-json-data-format/3669)
    11.  [12.10. XML 数据格式](https://learnku.com/docs/the-way-to-go/1210-xml-data-format/3670)
    12.  [12.11. 用 Gob 传输数据](https://learnku.com/docs/the-way-to-go/1211-uses-gob-to-transmit-data/3671)
    13.  [12.12. Go 中的密码学](https://learnku.com/docs/the-way-to-go/cryptography-in-1212-go/3672)

14. 

     

    第十三章. 错误处理与测试

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3673)
    2.  [13.1. 错误处理](https://learnku.com/docs/the-way-to-go/131-error-handling/3674)
    3.  [13.2. 运行时异常和 panic](https://learnku.com/docs/the-way-to-go/132-runtime-exception-and-panic/3675)
    4.  [13.3. 从 panic 中恢复（Recover）](https://learnku.com/docs/the-way-to-go/133-recovery-from-panic-recover/3676)
    5.  [13.4. 自定义包中的错误处理和 panicking](https://learnku.com/docs/the-way-to-go/error-handling-and-panicking-in-the-134-custom-package/3677)
    6.  [13.5. 一种用闭包处理错误的模式](https://learnku.com/docs/the-way-to-go/135-a-pattern-that-handles-errors-with-closures/3678)
    7.  [13.6. 启动外部命令和程序](https://learnku.com/docs/the-way-to-go/136-starts-external-commands-and-programs/3679)
    8.  [13.7. Go 中的单元测试和基准测试](https://learnku.com/docs/the-way-to-go/unit-testing-and-benchmarking-in-137-go/3680)
    9.  [13.8. 测试的具体例子](https://learnku.com/docs/the-way-to-go/specific-examples-of-138-testing/3681)
    10.  [13.9. 用（测试数据）表驱动测试](https://learnku.com/docs/the-way-to-go/139-driven-test-with-test-data-table/3682)
    11.  [13.10. 性能调试：分析并优化 Go 程序](https://learnku.com/docs/the-way-to-go/1310-performance-debugging-analysis-and-optimization-of-go-programs/3683)

15. 

     

    第十四章. 协程（goroutine）与通道（channel）

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3684)
    2.  [14.1. 并发、并行和协程](https://learnku.com/docs/the-way-to-go/141-concurrency-parallel-and-co-process/3685)
    3.  [14.2. 协程间的信道](https://learnku.com/docs/the-way-to-go/142-covariance-channel/3686)
    4.  [14.3. 协程的同步：关闭通道-测试阻塞的通道](https://learnku.com/docs/the-way-to-go/synchronization-of-143-cooperations-closing-channels-testing-blocked-channels/3687)
    5.  [14.4. 使用 select 切换协程](https://learnku.com/docs/the-way-to-go/144-switching-co-process-using-select/3688)
    6.  [14.5. 通道、超时和计时器（Ticker）](https://learnku.com/docs/the-way-to-go/145-channels-timeouts-and-timers-ticker/3689)
    7.  [14.6. 协程和恢复（recover）](https://learnku.com/docs/the-way-to-go/146-co-process-and-recovery-recover/3690)
    8.  [14.7. 新旧模型对比：任务和worker](https://learnku.com/docs/the-way-to-go/comparison-of-147-new-and-old-models-task-and-worker/3691)
    9.  [14.8. 惰性生成器的实现](https://learnku.com/docs/the-way-to-go/implementation-of-148-inert-generator/3692)
    10.  [14.9. 实现 Futures 模式](https://learnku.com/docs/the-way-to-go/149-implementation-of-futures-mode/3693)
    11.  [14.10. 多路复用 **已完成**](https://learnku.com/docs/the-way-to-go/1410-multiplexing/3694)
    12.  [14.11. 限制并发数 **已完成**](https://learnku.com/docs/the-way-to-go/1411-limiting-the-number-of-requests-processed-concurrently/3695)
    13.  [14.12. 链式操作 **已完成**](https://learnku.com/docs/the-way-to-go/1412-chaining-goroutines/3696)
    14.  [14.13. 多核运算 **已完成**](https://learnku.com/docs/the-way-to-go/1413-parallelizing-a-computation-over-a-number-of-cores/3697)
    15.  [14.14. 多核运算处理大量数据 **已完成**](https://learnku.com/docs/the-way-to-go/1414-parallelizing-a-computation-over-a-large-amount-of-data/3698)
    16.  [14.15. 漏桶算法 Leaky Bucket **已完成**](https://learnku.com/docs/the-way-to-go/1415-the-leaky-bucket-algorithm/3699)
    17.  [14.16. 标杆分析 Goroutines **已完成**](https://learnku.com/docs/the-way-to-go/1416-benchmarking-goroutines/3700)
    18.  [14.17. 使用 Channel 来并发读取对象 **已完成**](https://learnku.com/docs/the-way-to-go/1417-concurrent-acces-to-objects-by-using-a-channel/3701)

16. 

     

    第十五章. 网络，模板和网页应用

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3702)
    2.  [15.1. tcp 服务器](https://learnku.com/docs/the-way-to-go/151-tcp-server/3703)
    3.  [15.2. 一个简单的网页服务器](https://learnku.com/docs/the-way-to-go/152-a-simple-web-server/3704)
    4.  [15.3. 访问并读取页面](https://learnku.com/docs/the-way-to-go/153-access-and-read-pages/3705)
    5.  [15.4. 写一个简单的网页应用](https://learnku.com/docs/the-way-to-go/154-writes-a-simple-web-page-application/3706)
    6.  [15.5. 让 Web 应用更加健壮 **已完成**](https://learnku.com/docs/the-way-to-go/155-making-a-web-application-robust/3707)
    7.  [15.6. 在 Web 应用中使用模板 **已完成**](https://learnku.com/docs/the-way-to-go/156-writing-a-web-application-with-templates/3708)
    8.  [15.7. 探索 Template 扩展的功能 **已完成**](https://learnku.com/docs/the-way-to-go/157-exploring-the-template-package/3709)
    9.  [15.8. 一个多功能的精致的 WebServer **已完成**](https://learnku.com/docs/the-way-to-go/158-a-versatile-and-delicate-webserver/3710)
    10.  [15.9. RPC 远程调用 **已完成**](https://learnku.com/docs/the-way-to-go/159-rpc-remote-call/3711)
    11.  [15.10. 使用 netchan 跨网络实现消息传递 **已完成**](https://learnku.com/docs/the-way-to-go/1510-uses-netchan-across-network-to-implement-message-delivery/3712)
    12.  [15.11. Websocket 通讯 **已完成**](https://learnku.com/docs/the-way-to-go/1511-websocket-communication/3713)
    13.  [15.12. SMTP 发送邮件 **已完成**](https://learnku.com/docs/the-way-to-go/1512-smtp-sends-mail/3714)

17. 

     

    第十六章. 常见的陷阱与错误

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3715)
    2.  [16.1. 误用短声明导致变量覆盖](https://learnku.com/docs/the-way-to-go/163-closes-a-file-with-defer-when-an-error-occurs/3716)
    3.  [16.2. 误用字符串](https://learnku.com/docs/the-way-to-go/161-misuse-short-declarations-leading-to-variable-coverage/3717)
    4.  [16.3 发生错误时使用defer关闭一个文件](https://learnku.com/docs/the-way-to-go/when-164-uses-new-and-make/3718)
    5.  [16.4. 何时使用 new() 和 make()](https://learnku.com/docs/the-way-to-go/168-misuse-co-process-and-channel/3719)
    6.  [16.5. 不需要将一个指向切片的指针传递给函数](https://learnku.com/docs/the-way-to-go/162-misuse-string/3720)
    7.  [16.6. 使用指针指向接口类型](https://learnku.com/docs/the-way-to-go/165-does-not-need-to-transfer-a-pointer-to-a-slice-to-a-function/3721)
    8.  [16.7. 使用值类型时误用指针](https://learnku.com/docs/the-way-to-go/the-use-of-169-closures-and-association-courses/3722)
    9.  [16.8 误用协程和通道](https://learnku.com/docs/the-way-to-go/166-uses-pointers-to-point-to-the-type-of-interface/3723)
    10.  [16.9. 闭包和协程的使用](https://learnku.com/docs/the-way-to-go/167-misuse-pointers-when-using-value-types/3724)
    11.  [16.10. 糟糕的错误处理](https://learnku.com/docs/the-way-to-go/1610-bad-error-handling/3725)

18. 

     

    第十七章. 模式

    1.  [17.1. 关于逗号 ok 模式](https://learnku.com/docs/the-way-to-go/171-on-the-comma-ok-pattern/3726)
    2.  [17.2. defer 模式 **已完成**](https://learnku.com/docs/the-way-to-go/172-the-defer-pattern/3727)
    3.  [17.3.能见度模式 **已完成**](https://learnku.com/docs/the-way-to-go/173-the-visibility-pattern/3728)
    4.  [17.4. 操作者模式和接口 **已完成**](https://learnku.com/docs/the-way-to-go/174-the-operator-pattern-and-interface/3729)

19. 

     

    第十八章. 出于性能考虑的实用代码片段

    1.  [18.1. 字符串](https://learnku.com/docs/the-way-to-go/181-string/3730)
    2.  [18.2. 数组和切片](https://learnku.com/docs/the-way-to-go/182-array-and-slicing/3731)
    3.  [18.3. 映射](https://learnku.com/docs/the-way-to-go/183-mapping/3732)
    4.  [18.4. 结构体](https://learnku.com/docs/the-way-to-go/184-structure/3733)
    5.  [18.5. 接口](https://learnku.com/docs/the-way-to-go/185-interface/3734)
    6.  [18.6. 函数](https://learnku.com/docs/the-way-to-go/186-function/3735)
    7.  [18.7. 文件](https://learnku.com/docs/the-way-to-go/187-file/3736)
    8.  [18.8. 协程（goroutine）与通道（channel）](https://learnku.com/docs/the-way-to-go/188-co-range-goroutine-and-channel-channel/3737)
    9.  [18.9. 网络和网页应用](https://learnku.com/docs/the-way-to-go/189-network-and-web-application/3738)
    10.  [18.10. 其他](https://learnku.com/docs/the-way-to-go/1810-other/3739)
    11.  [18.11. 出于性能考虑的最佳实践和建议](https://learnku.com/docs/the-way-to-go/1811-best-practices-and-suggestions-for-performance-considerations/3740)

20. 

     

    第十九章. 构建完整的应用程序

    1.  [19.1. 简介 **已完成**](https://learnku.com/docs/the-way-to-go/191-introduction/3741)
    2.  [19.2. UrlShortener 项目介绍 **已完成**](https://learnku.com/docs/the-way-to-go/192-introducing-project-urlshortener/3742)
    3.  [19.3. 数据结构分析 **已完成**](https://learnku.com/docs/the-way-to-go/193-data-structure/3743)
    4.  [19.4. 用户界面：Web 网页前端 **已完成**](https://learnku.com/docs/the-way-to-go/194-our-user-interface-a-web-server-frontend/3744)
    5.  [19.5. 数据存储: gob **已完成**](https://learnku.com/docs/the-way-to-go/195-persistent-storage-gob/3745)
    6.  [19.6. 使用 Goroutines 来提高性能 **已完成**](https://learnku.com/docs/the-way-to-go/196-using-goroutines-for-performance/3746)
    7.  [19.7. 使用 Json 来存储 **已完成**](https://learnku.com/docs/the-way-to-go/197-using-json-for-storage/3747)
    8.  [19.8. 多台机器上的多线程 **已完成**](https://learnku.com/docs/the-way-to-go/198-multiprocessing-on-many-machines/3748)
    9.  [19.9. 使用 ProxyStore **已完成**](https://learnku.com/docs/the-way-to-go/199-using-a-proxystore/3749)
    10.  [19.10. 总结和优化 **已完成**](https://learnku.com/docs/the-way-to-go/1910-summary-and-enhancements/3750)

21. 

     

    第二十章. Go 在 Google App 引擎中的使用

    1.  [章节说明](https://learnku.com/docs/the-way-to-go/chapter-description/3751)

22. 

     

    第二十一章. Go 在现实世界中的使用

    1.  [章节说明 **已完成**](https://learnku.com/docs/the-way-to-go/chapter-description/3752)
    2.  [21.2. MROffice — Go 实现的 VOIP 系统 **已完成**](https://learnku.com/docs/the-way-to-go/the-voip-system-implemented-by-212-mroffice-go/3754)
    3.  [21.3. Atlassian— 虚拟机集群管理系统 **已完成**](https://learnku.com/docs/the-way-to-go/213-atlassian-virtual-machine-cluster-management-system/3755)
    4.  [21.4. Camlistore 个人住址存储系统 **已完成**](https://learnku.com/docs/the-way-to-go/214-camlistore-personal-address-storage-system/3756)

## Go入门指南(The Way to Go) 完整版PDF
```
对于学习 Go 编程语言的爱好者来说，这本书无疑是最适合你的一本书籍，这里包含了当前最全面的学习资源。本书通过对官方的在线文档、名人博客、书籍、相关文章以及演讲的资料收集和整理，并结合我自身在软件工程、编程语言和数据库开发的授课经验，将这些零碎的知识点组织成系统化的概念和技术分类来进行讲解。 我非常想要向发明这门精湛的语言的 Go 开发团队表示真挚的感谢，尤其是团队的领导者 Rob Pike、Russ Cox 和 Andrew Gerrand，他们阐述的例子和说明都非常的完美。同时，我还要感谢 Miek Gieben、Frank Muller、Ryanne Dolan 和 Satish V.J. 给予我巨大的帮助，还有那些 Golang-nuts 邮件列表里的所有的成员。 欢迎来到 Go 语言开发的奇妙世界！ 本翻译版本已获得原作者（Ivo Balbaert）本人授权，并表示支持开源事业的发展！ Ivo Balbaert 是一名已经拥有超过20年商业软件开发经验的物理学博士、程序员和项目经理。现居住位于比利时安特卫普的研究所，并在 CVO-安特卫普教授编程和数据库领域的知识。

目录
第一部分：学习 Go 语言
第1章：Go 语言的起源，发展与普及
第2章：安装与运行环境
第3章：编辑器、集成开发环境与其它工具
第二部分：语言的核心结构与技术
第4章：基本结构和基本数据类型
第5章：控制结构
第6章：函数（function）
第7章：数组与切片
第8章：Map
第9章：包（package）
第10章：结构（struct）与方法（method）
第11章：接口（interface）与反射（reflection）
第三部分：Go 高级编程
第12章：读写数据
第13章：错误处理与测试
第14章：协程（goroutine）与通道（channel）
第15章：网络、模版与网页应用
第四部分：实际应用
第16章：常见的陷阱与错误
第17章：模式
第18章：出于性能考虑的实用代码片段
```
