package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/oikomi/FishChatServer/monitor/models"
)

type UserController struct {
	beego.Controller
}

/**
 * @api {post} /user/ 用户注册
 * @apiName userStore
 * @apiGroup User
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} password 用户密码
 *
 * @apiParamExample {String} Request-Example:
 * username=13590210000&password=123456
 *
 * @apiSuccess {String} username 用户名
 * @apiSuccess {String} ticket 用户接口调用凭据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "注册成功",
 *         "data": {
 *             "username": "13590210000",
 *             "ticket": "abcdefg"
 *         }
 *     }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 20002,
 *         "errmsg": "用户名已经存在",
 *         "data": {
 *         }
 *     }
 */
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

/**
 * @api {get} /user 查看用户信息
 * @apiName userDetail
 * @apiGroup User
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 *
 * @apiSuccess {String} username 用户名
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "操作成功",
 *         "data": {
 *             "username": "13590210000"
 *         }
 *     }
 */

func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

/**
 * @api {put} /user 更新用户信息
 * @apiName userUpdate
 * @apiGroup User
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 * @apiParam {String} [oldPassword] 用户旧密码
 * @apiParam {String} [newPassword] 用户新密码
 *
 * @apiParamExample {String} Request-Example:
 * oldPassword=123456&newPassword=111111
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "操作成功",
 *         "data": {
 *         }
 *     }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 20003,
 *         "errmsg": "原密码正确",
 *         "data": {
 *         }
 *     }
 */

func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

/**
 * @api {post} /user/login 用户登录
 * @apiName userLogin
 * @apiGroup User
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} password 用户密码
 *
 * @apiParamExample {String} Request-Example:
 * username=13590210000&password=111111
 *
 * @apiSuccess {String} username 用户名
 * @apiSuccess {String} ticket 用户接口调用凭据
 *
 * @apiSuccessExample 正常回复
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "操作成功",
 *         "data": {
 *             "username": "13590210000",
 *             "ticket": "abcdefg"
 *         }
 *     }
 *
 * @apiErrorExample 用户名不存在回复
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 20003,
 *         "errmsg": "用户名不存在",
 *         "data": {
 *         }
 *     }
 *
 * @apiErrorExample 用户密码错误回复
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 20004,
 *         "errmsg": "用户密码错误",
 *         "data": {
 *         }
 *     }
 */
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = map[string]string{
			"abc": "user not exist",
		}
	}
	u.ServeJSON()
}

/**
 * @api {post} /user/logout 用户退出登录
 * @apiName userLogout
 * @apiGroup User
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "操作成功",
 *         "data": {
 *         }
 *     }
 */
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
