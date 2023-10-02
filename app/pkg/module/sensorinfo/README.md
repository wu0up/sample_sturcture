Go Playground Validator

https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835

https://yeqown.xyz/2018/08/29/Golang%E9%80%82%E7%94%A8%E7%9A%84DTO%E5%B7%A5%E5%85%B7/


如果你已经使用 Swagger 注解标记了路由的 payload 为 `CreateMemberDto`，并且使用了 Swaggo 自动生成 API 文档，那么通常情况下不需要再手动使用 `g.Bind(&member)`。

Swaggo 会根据你的注解生成 API 文档，并指定请求的 payload 类型。当你在 Swagger UI 或其他 Swagger 文档查看工具中查看 API 文档时，它应该会显示请求的 payload 格式为 `CreateMemberDto`。

在这种情况下，Gin 框架会自动将请求中的 JSON 数据解析为 `CreateMemberDto` 结构体，而无需显式使用 `g.Bind(&member)`。你可以直接在控制器方法中访问 `CreateMemberDto` 类型的参数，如下所示：

```go
func (c *MemberController) Create(g *gin.Context) {
    // 直接访问 CreateMemberDto 类型的参数
    var member dto.CreateMemberDto

    // 处理 member 数据并调用服务方法
    err := c.MemberService.Create(&member)

    // 处理错误和返回响应
    if err != nil {
        c.HandleError(g, err)
    } else {
        c.SuccessRes(g, nil)
    }
}
```

这种方式会使代码更加清晰，不需要额外的绑定步骤，因为 Gin 已经根据路由的注解将请求绑定到了指定的结构体上。如果你的 Swagger 文档和注解配置正确，请求和响应的 payload 应该会正确地反映在生成的文档中。


swaggo詳解
https://www.cnblogs.com/mayanan/p/17094955.html

使用swag fmt--> 優化format
