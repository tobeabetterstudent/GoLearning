# middleware 目录下存放自定义的中间件
1.  在gin的engine中通过Use方法增加中间件 而User方法原型是Use(middleware ...HandlerFunc) 因此我们的中间件必须定义为gin.HandlerFunc类型
# 日志中间件 loggrr.go 要求:
1.  日志可以记录到 File 中，定义一个 LoggerToFile 方法。
2.  日志可以记录到 MongoDB 中，定义一个 LoggerToMongo 方法。
3.  日志可以记录到 ES 中，定义一个 LoggerToES 方法。
4.  日志可以记录到 MQ 中，定义一个 LoggerToMQ 方法。