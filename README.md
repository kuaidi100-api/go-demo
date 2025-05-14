<h1 align="center"><a href="https://github.com/kuaidi100-api/kuadi100-api/" target="_blank">go-demo Project</a></h1>

## Introduce

go-demo 是由[快递100](https://api.kuaidi100.com/home)官方提供的golang sdk，方便测试使用。

go-demo 集成了快递查询（base）、电子面单与云打印（elec_print）、物流服务（border、work_order、corder、bsamecity、monitor）、增值服务（value_add）、跨境服务（international）等接口

## Features

- 提供测试类调试。

## Getting started

go-demo使用和测试可参考各级目录下的`*_test.go`文件。

```
# git clone https://github.com/kuaidi100-api/go-demo.git
```

## Add Config

使用前先配置[account.go](https://github.com/kuaidi100-api/go-demo/blob/main/config/account.go)，账号信息可以登录快递100获取https://poll.kuaidi100.com/manager/page/myinfo/enterprise （注意不要泄露快递100的账号密码以及授权key等敏感信息，以防被他人盗用！！！）


## Use Junit Test

以下是各个接口的测试示例，完整代码可在GitHub仓库中查看：
### 快递查询(base)接口
- [实时快递查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/base/query.go)
- [快递信息推送服务-订阅接口](https://github.com/kuaidi100-api/go-demo/blob/main/base/poll.go)
- [快递查询地图轨迹](https://github.com/kuaidi100-api/go-demo/blob/main/base/maptrack.go)
- [地图轨迹推送接口](https://github.com/kuaidi100-api/go-demo/blob/main/base/pollMap.go)

### 电子面单与云打印(elec_print)接口
- [电子面单下单接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/elec_order.go)
- [第三方平台账号授权](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/authThird.go)
- [第三方平台网点&面单余额接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/third_info.go)
- [电子面单取消接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/elec_cancel.go)

- [获取店铺授权超链接接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/shop_authorize.go)
- [提交销售订单获取任务接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/order_task.go)
- [提交售后（退货）订单获取任务接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/refundOrder_task.go)
- [物流信息发送接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/logistic_send.go)

- [自定义模板打印接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/elec_custom.go)
- [自定义模板打印复打接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/elec_printOld.go)
- [硬件状态查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/elec_print/print_task.go)

### 商家寄件(border)接口
- [寄件下单接口](https://github.com/kuaidi100-api/go-demo/blob/main/border/border_create.go)
- [运费预估接口](https://github.com/kuaidi100-api/go-demo/blob/main/border/border_price.go)
- [订单详情查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/border/border_detail.go)
- [取消寄件接口](https://github.com/kuaidi100-api/go-demo/blob/main/border/border_cancel.go)

### 寄件工单服务(work_order)接口
- [创建工单接口](https://github.com/kuaidi100-api/go-demo/blob/main/work_order/work_order_create.go)
- [查询工单接口](https://github.com/kuaidi100-api/go-demo/blob/main/work_order/work_order_query.go)
- [工单新增/查询留言接口](https://github.com/kuaidi100-api/go-demo/blob/main/work_order/work_order_reply.go)
- [工单附件上传接口](https://github.com/kuaidi100-api/go-demo/blob/main/work_order/work_order_upload.go)

### C端寄件(corder)接口
- [C端寄件 下单接口](https://github.com/kuaidi100-api/go-demo/blob/main/corder/corder_create.go)
- [C端寄件 价格预估接口](https://github.com/kuaidi100-api/go-demo/blob/main/corder/corder_price.go)
- [C端寄件 订单详情查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/corder/corder_detail.go)
- [C端寄件 订单取消接口](https://github.com/kuaidi100-api/go-demo/blob/main/corder/corder_cancel.go)

### 同城配送(bsamecity)接口
- [同城配送 下单接口](https://github.com/kuaidi100-api/go-demo/blob/main/bsamecity/bsamecity_order.go)
- [同城配送 价格预估接口](https://github.com/kuaidi100-api/go-demo/blob/main/bsamecity/bsamecity_price.go)
- [同城配送 订单取消接口](https://github.com/kuaidi100-api/go-demo/blob/main/bsamecity/bsamecity_cancel.go)
- [同城配送 订单预取消接口](https://github.com/kuaidi100-api/go-demo/blob/main/bsamecity/bsamecity_precancel.go)
- [同城配送 订单加小费接口](https://github.com/kuaidi100-api/go-demo/blob/main/bsamecity/bsamecity_addfee.go)

### 物流全链路监控服务(monitor)接口
- [订单导出接口](https://github.com/kuaidi100-api/go-demo/blob/main/monitor/monitor_orderExport.go)
- [订单发货接口](https://github.com/kuaidi100-api/go-demo/blob/main/monitor/monitor_sendOut.go)

### 增值服务(value_add)接口
- [智能地址解析接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/address_resoluton.go)
- [快递预估时效查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/estimate_time.go)
- [拦截改址接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/intercept_order.go)
- [运单附件查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/back_order.go)
- [快递100短信发送接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/sms_send.go)
- [快递智能识别单号](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/auto_number.go)
- [快递可用性接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/reachable.go)
- [快递面单OCR识别接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/det_ocr.go)
- [快递预估价格查询接口](https://github.com/kuaidi100-api/go-demo/blob/main/value_add/estimate_price.go)

### 跨境服务(international)接口
- [国际电子面单下单API](https://github.com/kuaidi100-api/go-demo/blob/main/international/api_call.go)
- [预约取件API](https://github.com/kuaidi100-api/go-demo/blob/main/international/pick_up.go)
- [取消预约API](https://github.com/kuaidi100-api/go-demo/blob/main/international/cancel_pick_up.go)
- [国际地址解析接口](https://github.com/kuaidi100-api/go-demo/blob/main/international/international_address_resolution.go)

## Tips
如需获取账号信息（如 key、customer、secret），或免费试用100单，请访问[API开放平台](https://api.kuaidi100.com/register/diff/)进行注册
