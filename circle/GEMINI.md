## API文档

这里是`Circle`官方的[开发者文档](https://developers.circle.com/)，请从以下路径检索不同API的`Request`、`Response`的Schema以及其他行为。

- Web3 Service: <https://developers.circle.com/api-reference/w3s/common/ping>
- Stablecoin: <https://developers.circle.com/api-reference/stablecoins/common/get-public-keys-v-2>

### 错误响应

`Circle`的API处理时遇到错误时，会返回以下的`json`结构体:

```json
{
    "code": <your_error_code>,
    "message": <your_error_message>
}
```

我们这里使用 @types.go 中的封装该error信息

```go
type PartialResponse struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
```

因此一个完整的响应结构体应该这么定义

```go
type CreateSubscriptionResponse struct {
		PartialResponse
		Data *Subscription `json:"data,omitempty"`
	}
```

## 封装原则

### HTTP Client

本库使用`github.com/jacexh/requests`进行封装，它简化了内置`net/http`的一些行为

## 流程定义

请按照以下流程完成代码的封装

- 搜索未封装的API，进行封装
- 编写新API的单元测试用例，执行测试，并修复错误
- 循环以上步骤


## Command

- 执行单元测试: `go test -v ./...`
