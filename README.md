# go-zmxy
## 芝麻信用SDK [查看官方文档](https://b.zmxy.com.cn/technology/openDoc.htm?relInfo=zhima.auth.info.authorize@1.0@1.3)

> 流程

![](http://zmxymerchant-prod.oss-cn-shenzhen.zmxy.com.cn/openApi/openDoc/img_20170509152742412.png)

- 调用接口名：zhima.auth.info.authorize 发起授权，返回授权 URL。
- 访问授权 URL，完成授权，回调到商家页面。
- 解密验签回调地址中所拼接的 params 和 sign 参数，获取 openid。

### DEMO

``` go
package zmxy

import (
	"crypto/rand"
	"errors"
	"fmt"

	"encoding/json"
	"github.com/WeberLong/go-zmxy/zmop"
	"github.com/WeberLong/go-zmxy/zmop/dto"
	"github.com/WeberLong/go-zmxy/zmop/requests"
	"github.com/WeberLong/go-zmxy/zmop/utils"

	"github.com/labstack/echo"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

const (
	gatewayUrl = "https://zmopenapi.zmxy.com.cn/openapi.do"
	appId      = "300***4"
	charset    = "UTF-8"

	pemZhimaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDyFmi6W9ArcUOcGWeVCrXmxuPiYGL0gQXsWVGDMljKnQDhjX68
iJZT9OI/2aTF9rIGy5oPThVaUgNKgBv+CqgqstXnfzj4xyOxzgJt25v7sUoWhwbT
6G97N7sIYH5mNyBfty4DiY+rPqjPKWpo0LyLfBpuR6wmoO0pnut8uQT5OQIDAQAB
****************************************************************
1sm7RC5O3lQOMCI+jwpe+5bPx0lNjkyQdmbh1f+siYJxqB4ZgUNGSnK35J4Kg8ZY
10QI4Fea7Nqo2XlHz5UCQQCzhTj08LCpasxH7oqAw6lQd5wqBhrHobISToVTdlz9
YOILqTwdezdlkGYRXMbZzuCMVoILE9Y3UB7Kd96u7O+4
-----END RSA PRIVATE KEY-----
`

	pemZhimaPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCi8aMzetxBQHAj2VaKlIegGjYr
****************************************************************
Rh32pc4OmgxJwfR7R6kCyTryoTiF4O9gT8JmnYiDPwa79jLCshICsTC4Bf+CJM3l
OlGQ5AYjAYOBZmL1ywIDAQAB
-----END PUBLIC KEY-----
`
)

type respAuth struct {
	Url string `json:"url"`
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

func Auth(c echo.Context) error {
	phone := c.QueryParam("phone")
	name := c.QueryParam("name")
	idcard := c.QueryParam("idcard")
	method := "1"
	if phone != "" {
		method = "1"
	}
	if idcard != "" && name != "" {
		method = "2"
	}

	client, err := zmop.NewZhimaClient(gatewayUrl, appId, charset, pemZhimaPrivateKey, pemZhimaPublicKey)
	if err != nil {
		return err
	}
	req := new(requests.ZhimaAuthInfoAuthorizeRequest)
	req.SetChannel("apppc")
	if method == "1" {
		req.InitBizParams("{\"auth_code\":\"M_H5\"}", "{\"mobileNo\":\""+phone+"\"}", method)
	} else if method == "2" {
		req.InitBizParams("{\"auth_code\":\"M_H5\"}", "{\"certNo\":\""+idcard+"\", \"certType\":\"IDENTITY_CARD\", \"name\":\""+name+"\"}", method)
	}

	res, err := client.Execute(req)
	if err != nil {
		return err
	}

	data := new(respAuth)
	data.Url = res
	return nil
}

func Notify(c echo.Context) error {
	params := c.QueryParam("params")
	sign := c.QueryParam("sign")
	bizPrivateKey, err := utils.GetPemKeyFile(pemZhimaPrivateKey)
	if err != nil {
		return err
	}
	zhimaPublicKey, err := utils.GetPemKeyFile(pemZhimaPublicKey)
	if err != nil {
		return err
	}

	res, err := utils.DecryptRSA(params, bizPrivateKey)
	if err != nil {
		return err
	}
	if utils.VerifySign(res, sign, zhimaPublicKey) == false {
		err = errors.New("签名错误")
		return err
	}
	u, err := url.Parse(res)
	if err != nil {
		return err
	}
	m, err := url.ParseQuery(u.Path)
	if err != nil {
		return err
	}
	if m["success"][0] == "true" {
		openId := m["open_id"][0]
		fmt.Printf("%+v\n", openId)
		return nil
	} else if m["success"][0] == "false" {
		err = errors.New("信息非法")
		return err
	}

	return nil
}

func Score(c echo.Context) error {
	client, err := zmop.NewZhimaClient(gatewayUrl, appId, charset, pemZhimaPrivateKey, pemZhimaPublicKey)
	if err != nil {
		return err
	}
	req := new(requests.ZhimaCreditScoreGetRequest)
	req.SetChannel("apppc")
	req.SetPlatform("zmop")

	unix_time := strconv.FormatInt(time.Now().UnixNano(), 10)

	openID := "123456"
	req.InitBizParams(openID, "w1010100100000000001", unix_time+"123"+strconv.FormatInt(RandInt64(1111111111111, 9999999999999), 10))

	res, err := client.ExecuteScore(req)
	if err != nil {
		return err
	}

	var respData dto.ResponseScore
	err = json.Unmarshal([]byte(res), &respData)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", respData.Score)

	return nil
}
```

### 注意
授权时，填写必要参数格式要正确，例如：
``` go
req.InitBizParams("{\"auth_code\":\"M_H5\"}", "{\"mobileNo\":\""+phone+"\"}", method)
```