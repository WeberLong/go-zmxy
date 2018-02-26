package zmop

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
	"github.com/WeberLong/go-zmxy/zmop/utils"

	"encoding/base64"
	"encoding/json"
	"errors"
)

type ZhimaClient struct {
	AppId string

	bizPrivateKey  []byte
	zhimaPublicKey []byte

	GatewayUrl string
	Format     string
	ApiVersion string
	Charset    string
	signType   string
	interfaces.ZhimaRequestParams
}

func (m *ZhimaClient) InitPemKey(pemZhimaPrivateKey, pemZhimaPublicKey string) error {
	var err error
	m.bizPrivateKey, err = utils.GetPemKeyFile(pemZhimaPrivateKey)
	if err != nil {
		return err
	}
	m.zhimaPublicKey, err = utils.GetPemKeyFile(pemZhimaPublicKey)
	// m.zhimaPublicKey = "MIG0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCi8aMzetxBQHAj2VaKlIegGjYrRIV52QQVBlHrG1BQvZHXYTTg2Y9FRft1ZLlXdOprlkygl90KCYkZgt4iiyc/MmrYRh32pc4OmgxJwfR7R6kCyTryoTiF4O9gT8JmnYiDPwa79jLCshICsTC4Bf+CJM3lOlGQ5AYjAYOBZmL1ywIDAQAB"
	if err != nil {
		return err
	}
	return nil
}

func (m *ZhimaClient) Execute(request interfaces.ZhimaRequest) (string, error) {
	bizParamStr := m.GetBizParamStr(&request)
	httpParams := m.GetSystemParams(&request)

	enBizParam, err := utils.EncryptRSA([]byte(bizParamStr), m.zhimaPublicKey)
	if err != nil {
		return "", err
	}
	(*httpParams)["params"] = base64.StdEncoding.EncodeToString(enBizParam)
	sign, err := utils.Sign(bizParamStr, m.bizPrivateKey)

	if err != nil {
		return "", err
	}
	(*httpParams)["sign"] = sign

	respData, err := utils.PostRequest(m.GatewayUrl, &request, httpParams)
	if err != nil {
		return "", nil
	}

	return respData, nil
}

func (m *ZhimaClient) ExecuteScore(request interfaces.ZhimaRequest) (string, error) {
	bizParamStr := m.GetBizParamStr(&request)
	httpParams := m.GetSystemParams(&request)

	enBizParam, err := utils.EncryptRSA([]byte(bizParamStr), m.zhimaPublicKey)
	if err != nil {
		return "", err
	}
	(*httpParams)["params"] = base64.StdEncoding.EncodeToString(enBizParam)
	sign, err := utils.Sign(bizParamStr, m.bizPrivateKey)

	if err != nil {
		return "", err
	}
	(*httpParams)["sign"] = sign
	respData, err := utils.PostRequestScore(m.GatewayUrl, &request, httpParams)
	if err != nil {
		return "", nil
	}

	// Decrypt if needed
	if respData.Encrypted {
		respData.Data, err = utils.DecryptRSA(respData.Data, m.bizPrivateKey)
		if err != nil {
			return "", err
		}
	}

	// Check format (must be json)
	var temp map[string]interface{}
	err = json.Unmarshal([]byte(respData.Data), &temp)
	if err != nil {
		return "", err
	}

	// Verify signature
	if !utils.VerifySign(respData.Data, respData.Sign, m.zhimaPublicKey) {
		return "", errors.New("verify signature failed")
	}

	return respData.Data, nil
}

func (m *ZhimaClient) GetBizParamStr(request *interfaces.ZhimaRequest) string {
	return utils.BuildQueryString((*request).GetApiParams(), true)
}

func (m *ZhimaClient) GetSystemParams(request *interfaces.ZhimaRequest) *map[string]string {
	if m.Charset == "" {
		m.Charset = "UTF-8"
	}
	version := (*request).GetApiVersion()
	if version == "" {
		version = m.ApiVersion
	}
	sysParams := make(map[string]string)

	sysParams["app_id"] = m.AppId
	sysParams["version"] = version
	sysParams["method"] = (*request).GetApiMethodName()
	sysParams["charset"] = m.Charset
	sysParams["scene"] = (*request).GetScene()
	sysParams["channel"] = (*request).GetChannel()
	sysParams["platform"] = (*request).GetPlatform()
	sysParams["ext_params"] = (*request).GetExtParams()
	return &sysParams
}

func NewZhimaClient(gatewayUrl, appId, charset, pemZhimaPrivateKey, pemZhimaPublicKey string) (*ZhimaClient, error) {
	client := &ZhimaClient{
		AppId:      appId,
		GatewayUrl: gatewayUrl,
		Format:     "json",
		ApiVersion: "1.0",
		Charset:    charset,
		signType:   "RSA",
	}

	err := client.InitPemKey(pemZhimaPrivateKey, pemZhimaPublicKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}
