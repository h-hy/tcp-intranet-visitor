package controllers

import (
	//	"encoding/json"

	"github.com/astaxie/beego"
	//	"github.com/oikomi/FishChatServer/monitor/models"
)

type DeviceController struct {
	beego.Controller
}

/**
 * @api {get} /device 查看用户设备列表
 * @apiName deviceList
 * @apiGroup Device
 *
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "errcode": 0,
 *         "errmsg": "操作成功",
 *         "data": [{
 *             "IMEI": "123456789101112",
 *             "nick": "123",
 *             "status": 1
 *             "work_model": 1,
 *             "volume": 6,
 *             "electricity": 100,
 *             "emeregncyPhone": "13590210000",
 *         }]
 *     }
 */

/**
 * @api {get} /device/:IMEI 查看用户设备详情
 * @apiName deviceDetail
 * @apiGroup Device
 *
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
 *             "IMEI": "123456789101112",
 *             "nick": "123",
 *             "status": 1
 *             "work_model": 1,
 *             "volume": 6,
 *             "electricity": 100,
 *             "emeregncyPhone": "13590210000",
 *         }
 *     }
 */

/**
* @api {post} /deivce 用户绑定设备
* @apiName deviceBinding
* @apiGroup Device
*
* @apiParam {String} username 用户名
* @apiParam {String} ticket 用户接口调用凭据
* @apiParam {String} IMEI 设备IMEI
* @apiParam {String} nick 设备昵称
*
* @apiParamExample {String} Request-Example:
* IMEI=1234567891011&nick=abc
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

/**
 * @api {delete} /deivce/:IMEI 用户删除绑定设备
 * @apiName deviceDestory
 * @apiGroup Device
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

/**
 * @api {put} /device/:IMEI 更新设备信息
 * @apiName DeivceUpdate
 * @apiGroup Device
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 * @apiParam {Number} work_model 工作模式
 * @apiParam {String} emeregncyPhone 设备紧急号码
 * @apiParam {Number} volume 设备音量
 * @apiParam {String} nick 设备昵称
 *
 * @apiParamExample {String} Request-Example:
 * work_model=1234567891011&nick=abc&volume=6&emeregncyPhone=13590210000
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

/**
 * @api {put} /device/:IMEI/action/location 设备实时定位
 * @apiName DeivceUpdateLocation
 * @apiGroup Device
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

/**
 * @api {put} /device/:IMEI/action/shutdown 设备关机
 * @apiName DeivceShutdown
 * @apiGroup Device
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

/**
 * @api {post} /device/:IMEI/sendVoice 发送聊天语音
 * @apiName DeivceSendVoice
 * @apiGroup Device
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} ticket 用户接口调用凭据
 * @apiParam {String} mp3Url mp3地址（等待商定）
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
