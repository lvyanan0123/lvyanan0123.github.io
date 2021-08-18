# activity

## 1. 充值送花币

```shell
进入redis：redis-cli -h 10.100.42.7 -p 19000
查看所有的兑换码：smembers redeem_code_act_1284_all
查看剩余的兑换码：smembers redeem_code_act_1284


强制发布活动：
http://192.168.7.38:8080/act/recharge_package/publish?json_string={"phone_models":"","update_user":"lvyanan@zhangyue.com","type_id":125,"app_id":"","create_user":"xinhai@zhangyue.com","create_time":1582648245,"s_name":"\u534e\u4e3a\u5145\u503c\u9001\u82b1\u5e01","id":2323,"is_excluded":0,"icon_url":"","ext":{"gift_infos":[{"timeTypeValue":"30","expire_info":{"days":"30","limit_type":1},"useFlag":"1","timeType":"day","info1":102399,"gift_name":"\u9605\u8bfb1000\u4ee3\u91d1\u5238-\u7070\u5ea6","info2":0,"desc":"","count":"0","gift_id":1284,"amount":"10","gift_img":"","ext":"","gift_type":"redeem_code","op_note":"OPPO\u65b0\u673a\u798f\u522910\u5143\u4ee3\u91d1\u5238\u7070\u5ea6","ambass":"None"}],"recharge_amount":"1","huawei_act_name":"","buy_type":"","recharge_times":"10000","huawei_act_id":""},"version_ids":"","record_id":"","status":1,"update_time":1627502714,"publish_time":1606938843,"description":"","deleted":0,"channel_ids":"","start_time":"2021-07-27+00:00:00","name":"\u534e\u4e3a\u5145\u503c\u9001\u82b1\u5e01lynn","amount":0,"end_time":"2021-08-31+00:00:00","icon_link":""}

[队列](http://192.168.7.38:15672/#/queues/%2F/q_act_recharge_gray)
发送数据：
{"status": "1", "ip": "36.110.103.118", "date": "2021-07-21 17:12:29", "phoneModel": "vivo+Y67A", "userAgent": "okhttp/3.11.0", "basicParams": {"vaid": "",
"p1": "Wa+JJkkCIPMDALE251Kb0Fuy", "oaid": "", "imei": "__622889015473258", "ip": "36.110.103.118", "version_code": "74302", "app_id": "zy4a7949", "version_id": "17430003", "platform_id": "501603", "channel_id": "107105", "idfa": "", "log_src": "", "fid": "41", "device": "vivo+Y67A", "pkg_name": "com.chaozh.iReaderFree", "android_id": "25d682c7343a3580"}, "oaid": "", "version": "17430003", "usr": "i2117184293", "channel": "107105", "orderId": "222T21072121111114", "giftCoupon": "0", "pkgName":"com.chaozh.iReaderFree", "adid": "__45d264c9181a1560", "orderOrigin": "yb_active_1_0","imei":"__622889015473258", "giftCouponFlag": "", "p1": "Wa+JJkkCIPMDALE251Kb0Fuy", "p7": "__622889015473258", "amount": 1.0, "giftAmount": "0", "rechargingType": "1", "originDesc": "", "actId":"2323"}

去礼品中心查i2113184293礼品是否赠送成功


VIP队列
{"usr":"i2117184293","vipLevel":1,"expireTime":"2021-09-04 11:25:50","isVipRenew":false,"subId":"383","productId":"100662","productDays":30,"channel":"107157","version":"17180003","phoneType":"OPPO+R9s","orderId":"999922337040872082559300001","serialGiftCoupon":0,"thirdGift":0,"thirdGiftExt":"","basicParams":{"p1":"WK7VTNbTQhoDAPgcKd0Xj6t7","channel_id":"107157","version_id":"17180003","platform_id":"501603","imei":"__628093016088019","device":"OPPO+R9s","version_code":"71400","fid":"41","oaid":"","ip":"36.110.103.118","app_id":"freeoppo81195","idfa":"","android_id":"","vaid":"","pkg_name":"","log_src":""}}

```

## 2. 消费提现活动

[ES查询地址](http://10.100.20.116:5602/app/kibana#/discover?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:now-30d,mode:quick,to:now))&_a=(columns:!(_source),filters:!(),index:'067341a0-e9c4-11ea-b26a-ef2d48a61824',interval:auto,query:(language:kuery,query:i136778954),sort:!('@timestamp',desc)))

## 3. 发布大转盘奖品替换掉docker里的gift_center_client.py

```shell
docker exec -ti www_activity_1 bash
pip list
Python3 -m site
进入zylib这个目录再进入zylib找到gift_center_client.py
vim /root/src/zylib/zylib/gift_center_client.py

1.删除所有内容
命令为：ggdG
其中，gg为跳转到文件首行；dG为删除光标所在行以及其下所有行的内容；
再细讲，d为删除，G为跳转到文件末尾行；
输入:set paste 命令 再复制

大转盘抽取奖品：
http://192.168.7.67:8080/act/wheel/draw/confirm?user_name=libaisan&version_id=1926
```
