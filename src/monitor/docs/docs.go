package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"0.1.0","swaggerVersion":"1.2","apis":[{"path":"/object","description":"Operations about object\n"},{"path":"/user","description":"Operations about Users\n"}],"info":{"title":"智能硬件接口文档","description":"提供智能硬件的操作接口"}}`
    Subapi string = `{"/object":{"apiVersion":"0.1.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/object","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"create","type":"","summary":"创建项目111","parameters":[{"paramType":"body","name":"body","description":"\"The object content\"","dataType":"Object","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"GET","nickname":"查看项目","type":"","summary":"通过objectid查看项目","parameters":[{"paramType":"path","name":"objectId","description":"项目id","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"GetAll","type":"","summary":"get all objects","responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"PUT","nickname":"update","type":"","summary":"update the object","parameters":[{"paramType":"path","name":"objectId","description":"\"The objectid you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"The body\"","dataType":"Object","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the object","parameters":[{"paramType":"path","name":"objectId","description":"\"The objectId you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"delete success!","responseModel":""},{"code":403,"message":"objectId is empty","responseModel":""}]}]}],"models":{"Object":{"id":"Object","properties":{"ObjectId":{"type":"string","description":"","format":""},"PlayerName":{"type":"string","description":"","format":""},"Score":{"type":"int64","description":"","format":""}}}}},"/user":{"apiVersion":"0.1.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
