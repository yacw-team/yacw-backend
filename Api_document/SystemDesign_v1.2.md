# 系统设计说明书

[toc]

## 一、引言

### 1.1 编写目的

### 1.2 背景

### 1.3 参考资料

### 1.4 术语表

## 二、体系结构设计

### 2.1 软件部署架构设计

本套软件有两种部署方式。

#### Web 应用部署

程序界面将以 Web 页面的形式在浏览器中执行，通过 HTTP 协议与后端服务器进行通信。

![部署架构设计1](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/系统架构设计/部署架构设计1.png)

#### Web 容器应用部署

程序将被包装进 Tauri 等 Web 容器应用中，通过本地 API 与后端服务器进行通信。

![部署架构设计2](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/系统架构设计/部署架构设计2.png)

### 2.2 软件技术架构设计

![技术架构图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/系统架构设计/软件架构设计.png)

## 三、系统功能模块

### 3.1 总体功能模块

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/功能模块/功能模块图.png)

### 3.1.1 对话功能模块

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/功能模块/对话功能模块.png)

### 3.1.2 翻译功能模块

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/功能模块/翻译功能模块.png)

### 3.1.3 游戏功能模块

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/功能模块/游戏功能模块.png)

### 3.1.4 心理咨询功能模块

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/功能模块/心理咨询功能模块.png)

## 四、类图

## 五、用例图

![用例图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/用例图.png)

## 六、泳道图

### 6.1 对话功能

    用户通过输入apikey才进入对话界面，之后需要选择prompt和character进行使用，
    其中这两个都可以在商店中寻找或者自己进行新的创建，选择后在后台形成完整的会话对象，
    数据库用来存储会话对象，之后openai会先判断apikey是否有效，
    无效则生成错误提示，有效就会生成回答，并将回答储存到数据库，在前台显示结果。

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/泳道图/对话泳道图.png)

### 6.2 翻译功能

    用户点击进入翻译界面，然后完成输入框的填入以及进行语言和语境的选择，后台便会接收到数据，
    将输入框的信息添加入数据库，之后调用api接口，实现openai的翻译，再将翻译结果储存在数据库中，
    并显示在前台界面。

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/泳道图/translate.png)

### 6.3 游戏功能

    用户点击进入游戏界面，首先需要选择游戏背景，点击确定后判定是否选择游戏背景，成功则将背景储存进数据库中，
    openai开始生成故事选项，然后openai生成四个选项，供玩家选择，玩家选择后由openai生成新的剧情，
    之后openai判断新的剧情中玩家是否结束游戏或者通关，如果是，则不需要在生成选项，若否，则继续生成四个选项供玩家选择。

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/泳道图/game.png)

### 6.4 心理咨询功能

    用户输入apikey进入咨询方向界面，从后台查询数据库获取咨询方向信息，并在前台显示有什么咨询方向，
    用户选择咨询方向后可以选择是否要导入评估报告，如果是，则会获取历史报告，并选择评估报告，否则可以直接输入咨询的对话，形成会话对象。
    这时候openai便会判断apikey是否有效，有效则会生成回答，并存储这次的数据，并显示结果。

![功能模块图](/Users/huyanjie/Documents/git/团队项目/Alpha冲刺/yacw-backend/api_document/imgs/系统设计/泳道图/心理咨询泳道图.png)

## 七、接口设计

使用 RESTful 规范设计 API，使用 JSON 作为数据交换格式。

### 7.1 对话界面接口设计

入口：/v1/chat

#### 用户输入apiKey

URL：POST/v1/chat/apikey

Req:

```
{
	apiKey:string,
}
```

Res:

```
200 OK

401 Unauthorized //api错误
```



#### 新建对话(发送的第一条对话)

URL: POST /v1/chat/new

Req:

```
{
    apiKey: string,
    modelId: string,
    content: {
        personalityId: string,//构造system
        user: string,      // user input
    }
}
```

Res:

```
{
    chatId: int,
    modelId: int,
    content: {
        personalityId: string,
        user: string,      // user input
        assistant:string,
        title:string,
    },
}
```

#### 删除对话

URL：DELETE /v1/chat/chat

Req:

```
{
    apiKey: string,
    chatId: string,
}
```

Res:
200 OK

#### 发送对话

URL: POST /v1/chat/chat

Req:

```
{
    apiKey: string,
    chatId: string,
    content: {
        user: string,
    },
}
```

Res:

```
{
    chatId: string,

    content: {
        user: string,
        assistant: string,
    },
    id:{
        usermsgid:string,
        assmsgid:string,
    }
}
```

#### 根据uid获得chataid

URL: POST /v1/chat/getchat

Req:

```
{
    apiKey: string,
}
```

Res:

```
{
    chat:[
        {
            chatID:string,
            title:string,
        }
    ]
}
```

#### 根据chatid获得message

URL: POST /v1/chat/getmessage

Req:

```
{
     apiKey: string,
    chatId: string,
}
```

Res:

```
{
    chatId:string,

    messages:[
                {
                  type: string,
                  content: string
                },
    				]
}
```

#### 获取 Prompts 商店内容

URL：GET /v1/chat/prompts?tyep=?

Res:

```
{
    prompts: [
        {
            id: string,
            name: string,					//名称
            description: string,	//描述
            icon:string,					//表情
            prompts:string,				//具体的prompt
        },
        ...
    ],
}
```



#### 获取用户自主创建的Prompts

URL：POST/v1/chat/myprompts

Req:

```
{
	apiKey:string,
}
```

Res：

```
{
    prompts: [
        {
            id: string,
            name: string,					//名称
            description: string,	//描述
            prompts:string,				//具体的prompt
        },
        ...
    ],
}
```



#### 创建用户 Prompts

URL：POST /v1/chat/prompts

Req:

```
{
    apiKey: string,
    name: string,					//名称
    description: string,	//描述
    icon:string,            //
    prompts: string,			//具体的prompts
}
```

Res:

```
{
    id: string,
    name: string,					//名称
    description: string,	//描述
    icon:string,					//
    prompts: string,			//具体的prompts
}
```



#### 删除用户 Prompts

URL：DELETE /v1/chat/prompts

Req:

```
{
    apiKey: string,
    promptsId: string,
}
```

Res:
200 OK

#### 获取模型人格

URL：GET /v1/chat/personality

Req:

```
{
	apiKey:string,
}
```

Res:

```
{
    personality: [
        {
            id: string,
            name: string,
            description: string,
            prompots：string,			//*
        },
        ...
    ],
}
```



#### 获取用户创建的人格

URL:POST/v1/chat/mypersonality

Req:

```
{
	apiKey:string,
}
```

Res:

```
{
    personality: [
        {
            id: string,
            name: string,
            description: string,
            prompots：string,			//*
        },
        ...
    ],
}
```



#### 创建用户模型人格

URL：POST /v1/chat/personality

Req:

```
{
    apiKey: string,
    name: string,
    description: string,
    prompts: string,
}
```

Res:

```
{
    id: string,
    name: string,
    description: string,
    prompts: string,
}
```

#### 删除用户模型人格

URL：DELETE /v1/chat/personality

Req:

```
{
    apiKey: string,
    personalityId: string,
}
```

Res:
200 OK



### 7.2 翻译界面接口设计

入口：/v1/translate

#### 发送翻译内容

URL: POST /v1/translate/translate

Req:

```
{
    apiKey: string,
    modelId:int
    content: {
        emotion: string,
        style: string,
        preTranslate: string,
    },
    from: string,
    to: string,
}
```

Res:

```
{
    content: {
        emotion: string,
        style: string,
        preTranslate: string,
        translated: string,
    },
    from: string,
    to: string,
}
```



### 7.3 心理咨询界面接口设计

#### 开始聊天

URL: POST /v1/psychology/chat

Req:

```
{
  selectDirection:int,
  inputContent:string,
}
```

Res:

```
{
    outputContent:string,
}
```

#### 创建方向

URL: POST /v1/psychology/createDirection

Req:

```
{
  directionID:int,
  Direction:string,
}
```

Res:
200 OK

#### 删除方向

URL: DELETE /v1/psychology/deleteDirection

Req:

```
{
  directionID:int,
}
```

Res:
200 OK

### 7.4 游戏界面接口

```
{
    script: [
        {
            id: string,
            name: string,
            description: string,
        },
        ...
    ],
}
```

#### 新建游戏对话

URL：POST /v1/game/new

Req:

```
{

    apiKey: string,
    scriptId: string,

}
```

Res:

```
{
    id: string,
    scriptId: string,
    script: {
        background: string,
        protagonist: string,
        goal: string,
        evnets: string[],
    },
    content: string,
}
```

#### 发送游戏对话

URL：POST /v1/game/chat

Req:

```
{
    apiKey: string,
    gameId: string,
    content: string,
}
```

Res:

```
{
    content: string,
}
```

#### 创建用户剧本

URL：POST /v1/game/script

Req:

```
{
    apiKey: string,
    name: string,
    description: string,
    script: {
        background: string,
        protagonist: string,
        goal: string,
        evnets: string[],
    },
}
```

Res:

```
{
    id: string,
    name: string,
    description: string,
    script: {
        background: string,
        protagonist: string,
        goal: string,
        evnets: string[],
    },
}
```

#### 删除用户剧本

URL：DELETE /v1/game/script

Req:

```
{
    apiKey: string,
    scriptId: string,
}
```

Res:
200 OK

## 八、系统安全与权限设计

### 5.1 系统安全

#### 通讯安全

前后端通讯使用 HTTPS 协议，其安全级别为 TLS 1.2，加密算法为 AES 256。可以有效防止用户的对话内容与 API Key 等敏感信息被窃取。

#### 数据安全

在存储用户数据时，我们不考虑采用账户形式记录用户身份，而是采用由用户提供的 API Key 加盐哈希截断生成的字符串作为用户的唯一标识 UID，服务器不存储用户的敏感信息。这样即使数据库泄露，也无法通过 API Key 反推出用户的真实身份。

#### 服务器安全

考虑在 API 反向代理服务器上部署 WAF，以防止常见的攻击手段。还计划部署并发限制与 IP 黑名单，以防止恶意攻击。

#### 数据库安全

采用 ORM 框架，防止 SQL 注入攻击。

### 5.2 权限控制

用户仅有权限访问与自己相关的数据及服务器公用数据，例如用户创建的角色、Prompts、游戏剧本等，以及公开的 Prompts 商店里的内容。

将通过部署 Gin 中间件，对所有与用户数据存储相关的 API 进行权限控制，根据 API Key 生成的 UID 验证用户身份，防止用户访问他人数据。
