go+vue：前后端一体化
===
##网络转发和网络代理
网络转发是指通过一个路由器对数据进行转发，期间会更改请求的信息，将请求参数发送给目标服务，然后返回也一样<br>
网络代理:proxy_ip服务器通过端口接收客户端请求，并通过socket请求真实目标服务器，客户端不知道真实服务器ip多少<br>
正向代理：客户端不直接连服务器，通过代理服务器访问，例如访问google，由于防火墙，需要通过vpn、浏览器代理等
反向代理：