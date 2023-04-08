### LV1

基于go-zero框架的rpc框架zrpc简单使用consul

##### rpc方配置

在goctl生成的配置文件rpc/internal/config 添加consul的配置

![image-20230408151346663](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408151346663.png)

在goctl生成的 rpc/etc/yaml 文件添加consul的配置，去掉go-zero自带的etcd配置

![image-20230408150800795](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408150800795.png)

在rpc服务的启动文件注册consul

![image-20230408150852331](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408150852331.png)

##### api方配置

在api服务的配置文件api/etc 添加consul配置

![image-20230408162152409](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408162152409.png)

在api服务的配置文件api/internal/config 添加zrpc.RpcClientConf配置

![image-20230408151830355](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408151830355.png)

最后在api服务的启动文件导入一个省略包  *_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"*

![image-20230408152208528](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230408152208528.png)

### LV2

zRPC框架的balancer模块已经内置了P2C负载均衡算法，在客户端进行负载均衡

[企业级RPC框架zRPC | go-zero](https://go-zero.dev/cn/docs/blog/showcase/zrpc#balancer模块)