define({ "api": [
  {
    "type": "post",
    "url": "/device/:IMEI/sendVoice",
    "title": "发送聊天语音",
    "name": "DeivceSendVoice",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "mp3Url",
            "description": "<p>mp3地址（等待商定）</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "put",
    "url": "/device/:IMEI/action/shutdown",
    "title": "设备关机",
    "name": "DeivceShutdown",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "put",
    "url": "/device/:IMEI",
    "title": "更新设备信息",
    "name": "DeivceUpdate",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "work_model",
            "description": "<p>工作模式</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "emeregncyPhone",
            "description": "<p>设备紧急号码</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "volume",
            "description": "<p>设备音量</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "nick",
            "description": "<p>设备昵称</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "work_model=1234567891011&nick=abc&volume=6&emeregncyPhone=13590210000",
          "type": "String"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "put",
    "url": "/device/:IMEI/action/location",
    "title": "设备实时定位",
    "name": "DeivceUpdateLocation",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "post",
    "url": "/deivce",
    "title": "用户绑定设备",
    "name": "deviceBinding",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "IMEI",
            "description": "<p>设备IMEI</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "nick",
            "description": "<p>设备昵称</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "IMEI=1234567891011&nick=abc",
          "type": "String"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "delete",
    "url": "/deivce/:IMEI",
    "title": "用户删除绑定设备",
    "name": "deviceDestory",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "get",
    "url": "/device/:IMEI",
    "title": "查看用户设备详情",
    "name": "deviceDetail",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n        \"IMEI\": \"123456789101112\",\n        \"nick\": \"123\",\n        \"status\": 1\n        \"work_model\": 1,\n        \"volume\": 6,\n        \"electricity\": 100,\n        \"emeregncyPhone\": \"13590210000\",\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "type": "get",
    "url": "/device",
    "title": "查看用户设备列表",
    "name": "deviceList",
    "group": "Device",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": [{\n        \"IMEI\": \"123456789101112\",\n        \"nick\": \"123\",\n        \"status\": 1\n        \"work_model\": 1,\n        \"volume\": 6,\n        \"electricity\": 100,\n        \"emeregncyPhone\": \"13590210000\",\n    }]\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/device.go",
    "groupTitle": "Device"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "F__git_client_RDAWatchServer_src_github_com_oikomi_FishChatServer_monitor_doc_main_js",
    "groupTitle": "F__git_client_RDAWatchServer_src_github_com_oikomi_FishChatServer_monitor_doc_main_js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/user",
    "title": "查看用户信息",
    "name": "userDetail",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n        \"username\": \"13590210000\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/user.go",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/user/login",
    "title": "用户登录",
    "name": "userLogin",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>用户密码</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "username=13590210000&password=111111",
          "type": "String"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正常回复",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n        \"username\": \"13590210000\",\n        \"ticket\": \"abcdefg\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "用户名不存在回复",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 20003,\n    \"errmsg\": \"用户名不存在\",\n    \"data\": {\n    }\n}",
          "type": "json"
        },
        {
          "title": "用户密码错误回复",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 20004,\n    \"errmsg\": \"用户密码错误\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/user.go",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/user/logout",
    "title": "用户退出登录",
    "name": "userLogout",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/user.go",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/user/",
    "title": "用户注册",
    "name": "userStore",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>用户密码</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "username=13590210000&password=123456",
          "type": "String"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"注册成功\",\n    \"data\": {\n        \"username\": \"13590210000\",\n        \"ticket\": \"abcdefg\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 20002,\n    \"errmsg\": \"用户名已经存在\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/user.go",
    "groupTitle": "User"
  },
  {
    "type": "put",
    "url": "/user",
    "title": "更新用户信息",
    "name": "userUpdate",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket",
            "description": "<p>用户接口调用凭据</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "oldPassword",
            "description": "<p>用户旧密码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "newPassword",
            "description": "<p>用户新密码</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "oldPassword=123456&newPassword=111111",
          "type": "String"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 0,\n    \"errmsg\": \"操作成功\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"errcode\": 20003,\n    \"errmsg\": \"原密码正确\",\n    \"data\": {\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/user.go",
    "groupTitle": "User"
  }
] });
