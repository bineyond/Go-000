# 工程管理改造


## 微服务架构

按 kratos v2 思想，架构分成。

## API 设计（包括 API 定义、错误码规范、Error 的使用）

Probuf 定义接口数据的全部字段，每个模块专注于自己的事情。

错误规范，统一错误信息结构。内部错误，不暴露到调用方，统一转换后返回。

## gRPC 的使用

调用方有Java也有Go，统一API接口，接口即文档。

## Go 项目工程化（项目结构、DI、代码分层、ORM 框架）

```
├── LICENSE
├── README.md
├── api
│   └── helloworld
│       ├── errors
│       │   ├── helloworld.pb.go
│       │   ├── helloworld.proto
│       │   └── helloworld_errors.pb.go
│       └── v1
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd
│   └── helloworld
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs
│   └── config.yaml
├── generate.go
├── go.mod
├── go.sum
└── internal
    ├── biz
    │   ├── README.md
    │   ├── biz.go
    │   └── greeter.go
    ├── conf
    │   ├── conf.pb.go
    │   └── conf.proto
    ├── data
    │   ├── README.md
    │   ├── data.go
    │   └── greeter.go
    ├── server
    │   ├── grpc.go
    │   ├── http.go
    │   └── server.go
    └── service
        ├── README.md
        ├── greeter.go
        └── service.go
```

## 并发的使用（errgroup 的并行链路请求）

目前,在推送数据中使用，及项目启动后台侦听服务。

## 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）

- ELK 日志收集，至今使用项目体量不大，还应用到实际项目中去。
- Opentracing 和ELK一样计划加入项目，用于分析项目模块性能及瓶颈。
- Prometheus 目前简单分析调用数据信息，正在逐步增加数据指标
- Kafka 未使用到。

## 缓存的使用优化（一致性处理、Pipeline 优化)

- 随机缓存失效时间，避免key集中失效
- singleflight 避免热点数据失效，高并发请求穿透到DB
- Pipeline 目前还没有应用上，分析改造中。


# 毕业总结

之前项目一直采用的是巨石架构，也自学过一些微服务的支持，在这课程里，跟着毛大深入工程管理、微服务架构、内存及高并发模型，代领我搭建一套完整的架构系统，运用到项目中去，极大地提高了项目的工程化，及业务的健壮性。

除此之外，毛大一直强调的方法论，每个人都要有一套适合自己的方法论，要学会总结，归纳吸收。以前查一个资料就是一顿搜索，搞定完事，不知其所以然，现在要会刻意地去找第一手资料看，培养自己的英语阅读能力。其次，阅读优秀的代码学习，为什么这么实现，提高自己的编程思想。

这13周课程是物超所值，干活满满，需要3遍、4遍不停地去理解巩固。
