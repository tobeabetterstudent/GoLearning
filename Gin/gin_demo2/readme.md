实现了路由分组 v1版本、v2版本。
实现了生成签名和验证验证。
实现了在配置文件中读取路由配置如port mode。

比如我们的接口地址是这样的：

/v1/product/add
/v1/member/add
/v2/product/add
/v2/member/add
假设需求是这样的，接口支持多种请求方式，v1 不需签名验证，v2 需要签名验证。