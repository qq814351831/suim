# connect 客户端通讯协议文档

connect 支持两种协议和客户端通讯 websocket， tcp。

## websocket

**请求URL**

ws://DOMAIN/ws

**HTTP请求方式**

Websocket（JSON Frame），请求和返回协议一致

**请求和返回json**

```json
{
    "ver": 102,
    "op": 10,
    "seq": 10,
    "body": {"data": "xxx"}
}
```

``` protobuf
// 上行数据
    message Input{
    string Version = 1; //版本 
    PackageType type = 2; // 包的类型
    int64 request_id = 3; // 请求id
    bytes data = 4; // 数据
}

// 下行数据
message Output {
    PackageType type = 1; // 包的类型
    int64 request_id = 2; // 请求id
    int32 code = 3; // 错误码
    string message = 4; // 错误信息
    bytes data = 5; // 数据
}

PackageType {
PT_UNKNOWN = 0; // 未知
PT_SIGN_IN = 1; // 设备登录请求
PT_SYNC = 2; // 消息同步触发
PT_HEARTBEAT = 3; // 心跳
PT_MESSAGE = 4; // 消息投递
PT_SUBSCRIBE_ROOM = 5; // 订阅房间
}
```

**请求和返回参数说明**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
| ver        | true  | int | 协议版本号 |
| op         | true  | int    | 指令 |
| seq        | true  | int    | 序列号（服务端返回和客户端发送一一对应） |
| body          | true | string | 授权令牌，用于检验获取用户真实用户Id |

## tcp

**请求URL**

tcp://DOMAIN

**协议格式**

二进制，请求和返回协议一致

**请求&返回参数**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
| package length        | true  | int32 bigendian | 包长度 |
| header Length         | true  | int16 bigendian    | 包头长度 |
| ver        | true  | int16 bigendian    | 协议版本 |
| operation          | true | int32 bigendian | 协议指令 |
| seq         | true | int32 bigendian | 序列号 |
| body         | false | binary | $(package lenth) - $(header length) |

## 指令

| 指令     | 说明  | 
| :-----     | :---  |
| 2 | 客户端请求心跳 |
| 3 | 服务端心跳答复 |
| 5 | 下行消息 |
| 7 | auth认证 |
| 8 | auth认证返回 |

PT_UNKNOWN = 0; // 未知 PT_SIGN_IN = 1; // 设备登录请求 PT_SYNC = 2; // 消息同步触发 PT_HEARTBEAT = 3; // 心跳 PT_MESSAGE = 4; // 消息投递
PT_SUBSCRIBE_ROOM = 5; // 订阅房间
