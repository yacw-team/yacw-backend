# 系统设计说明书

[toc]

### 7.1 对话界面接口设计

入口：/v1/chat

#### 用户输入apiKey

URL：POST/v1/chat/apiKey

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
    chatId:string,//由前端生成
    content: {
        personalityId: string,//构造system
        user: string,      // user input
    }
}
```

Res:

```
{
    chatId: string,
    modelId: string,
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
    Prompts: [
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
    Prompts: [
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
    prompts: string,			//具体的prompts
}
```

Res:

```
{
    id: string,
    name: string,					//名称
    description: string,	//描述
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

### 7.4 游戏界面接口

#### 剧本展示

URL：GET /v1/game/story

Res:

```
{
		story:[
        {
						gameId:string,
						name:string,
						description:string,
        },
        ...
		],
}
```



#### 选择剧本

URL：POST /v1/game/new

Req:

```
{
    apiKey: string,
    gameId: string,
}
```

Res:

```
{
    story:string,//故事的发展剧情
    choice:[
    	string,
    	...
    ],
    round:int,
}
```

#### 发送游戏对话

URL：POST /v1/game/chat

Req:

```
{
    apiKey: string,
    choiceID: string,//A,B,C,D
}
```



Res:

```
{
	story:string,//故事的发展剧情
    choice:[
    	string,
    	...
    ],//(A,B,C,D四个选项)
    round:int,
}
```



