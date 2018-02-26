package interfaces

type ZhimaRequest interface {
	GetApiMethodName() string
	GetApiParams() *map[string]string
	GetFileParams() *map[string]string
	SetApiVersion(string)
	GetApiVersion() string
	SetScene(string)
	GetScene() string
	SetChannel(string)
	GetChannel() string
	SetPlatform(string)
	GetPlatform() string
	SetExtParams(string)
	GetExtParams() string
}

type ZhimaRequestParams struct {
	ApiParams  *map[string]string
	FileParams *map[string]string
	ApiVersion string
	Scene      string
	Channel    string
	Platform   string
	ExtParams  string
}
