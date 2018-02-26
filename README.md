# go-zmxy
## go-zmxy [芝麻信用SDK](https://b.zmxy.com.cn/technology/openDoc.htm?relInfo=zhima.auth.info.authorize@1.0@1.3)

> 流程

![](http://zmxymerchant-prod.oss-cn-shenzhen.zmxy.com.cn/openApi/openDoc/img_20170509152742412.png)

- 调用接口名：zhima.auth.info.authorize 发起授权，返回授权 URL。
- 访问授权 URL，完成授权，回调到商家页面。
- 解密验签回调地址中所拼接的 params 和 sign 参数，获取 openid。

### 注意
    授权时，填写必要参数格式要正确，例如：
    ```
    req.InitBizParams("{\"auth_code\":\"M_H5\"}", "{\"mobileNo\":\""+phone+"\"}", method)
    ```