# http协议从1到3发生了什么

## http1.0
 - **http1.0的问题：** 	每次请求都要新起TCP连接，浪费太大
 - **[http1.1优化](https://p3-tt.byteimg.com/origin/pgc-image/c93ed5e27e4841caae9d4750b20dab6a?from=pc)**
	**A.** 支持长连接 keepalive: 一个TCP连接，多次请求来回用，每个请求依次处理
	**B.** 支持管道化请求
		一个TCP连接，比长连接的优点在于，可以一次发多个请求，依次等响应
		[链接](https://p6-tt.byteimg.com/origin/pgc-image/9d0ac20b51414a92ab3eba15654c06e6?from=pc)
 - **http1.1的问题**
	**A.** 不管是长连接，还是keepalive，响应都是依次处理（按请求顺序），某一个http卡住流程就全卡住了
	临时解决方案：
		浏览器对于一个域名，可以建多个连接（chrome是5个），减少依次处理的影响
	**B.** 服务器不能主动推送信息下去
	**C.** header浪费流量

## http2.0

 - http引入多路复用（请求与响应都拆成帧,发端与收端收完帧后再组合起来) 		多路复用允许同时通过单一的 HTTP/2
   [连接发起多重的请求](https://p1-tt.byteimg.com/origin/pgc-image/f3e2f60a0be443bf991f5df2027846f6?from=pc)-响应消息不用多建连接了，一个连接就可以干所有的事情了，也不怕一个http请求卡住了
   [比较](https://pic2.zhimg.com/80/b1e608ddb7493608efea3e76912aabe1_720w.jpg?source=1940ef5c)
		
	**http2.0**
		[图](https://pic2.zhimg.com/80/ae5418b7da1c6593fd6addad0310faa5_720w.jpg?source=1940ef5c)
	**A.** 把文本协议改二进制协议
	**B.** 提高了协议的效率

	**header压缩（使用专门的HPACK算法）**
		**A.** http1.1不支持header部分压缩，由于cookie和ua比较长，每次都要传，占用大量流量
		**B.** http2.0支持了，减少了header的传输


 - **移动端现状**
	**A.** ios9起支持http2.0
	**B.** andriod5起完美支持http2.0

 - **http2.0的问题**
 	 ```
	现在http层使用了多路复用的技术，不再会在http层上卡住了
	但由于http底层使用tcp协议，并且现在一个域名上只使用一个tcp连接
	```
	
	**我们知道tcp传输的时候，是有序的（注意快重传算法）**

|慢启动 |
|--|
|慢启动 | 
| 拥塞避免 | 
|快重传 | 
|快恢复|

	
	所以一旦服务器上，有一个tcp卡住了，就会导致整个域名上的所有http请求全部都被卡住
	所以我们需要从底层上抛弃tcp协议

## QUIC协议(http3.0)

	http3.0主要问题是要绕开tcp协议，但如果新加一种协议，则中间的设备都需要升级，带来很多问题
	google在2013年提出了QUIC协议(quick udp internet connections）
	2018年被命名为http3.0协议

## QUIC特点

	基于UDP的传输层协议
	可靠性
		udp本身不可靠，QUIC在udp的基础上做了些改造，提供了tcp类似的可靠性
		QUIC提供了如下功能：
			数据包重传
			拥塞控制
			调整传输节奏
	实现了无序、并发字节流
		quic单个数据流可以保证有序交付，但多个数据流间可能乱序（与http2.0中http帧类似）
	快速握手
		QUIC提供0-RTT和1-RTT的连接建立
	使用TLS3.0传输层安全协议
	
[http3.0图](https://p6-tt.byteimg.com/origin/pgc-image/092b605b771442d0b17b8a3bd19a97af?from=pc)
	

 - QUIC的连接建立
	**A.** TCP要三次握手建立连接，需要1.5RTT
	**B.** QUIC提出一种新的连接建立机制，实现了快速握手功能，一次QUIC连接建立可以用0-RTT或1-RTT来[建立链接](https://p1-tt.byteimg.com/origin/pgc-image/f2c5204e17734da99167346418e9b023?from=pc)
	
**两个步骤：**
	最终（重复）握手
	[初始握手](https://p1-tt.byteimg.com/origin/pgc-image/a6f20f7e846d44c29c5a52fc34ea3ca0?from=pc)


