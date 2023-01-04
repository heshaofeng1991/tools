package config

type NaCos struct {
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	Namespace string `json:"namespace"`
	User      string `json:"user"`
	Password  string `json:"password"` //密码
	DataId    string `json:"dataid"`   //配置名
	Group     string `json:"group"`    //分组
	Cluster   string `json:"cluster"`  //tag 集群名
	Open      bool   `json:"open"`     //开关 true开启 false关闭（使用本地）
}

type Service struct {
	Name      string            `json:"name"`      //服务名
	Metadata  map[string]string `json:"metadata"`  //附加参数
	Host      string            `json:"host"`      //ip
	Port      uint64            `json:"port"`      //端口
	Group     string            `json:"group"`     //组名
	Cluster   string            `json:"cluster"`   //集群名 tag
	CheckAddr string            `json:"checkAddr"` //心跳地址
	Namespace string            `json:"namespace"` //命名空间
}

type Oss struct {
	Endpoint string `json:"endpoint"` //地址
	Bucket   string `json:"bucket"`   //桶名
	Id       string `json:"id"`       //加密id
	Secret   string `json:"secret"`   //密码key
}

type DB struct {
	Host     string `json:"host"`     //地址
	Port     int    `json:"port"`     //端口
	Name     string `json:"name"`     //库名
	User     string `json:"user"`     //用户名
	Password string `json:"password"` //密码
	Prefix   string `json:"prefix"`   //表前缀
}

type Redis struct {
	Host         string `json:"host"`         //地址
	Port         int    `json:"port"`         //端口
	DB           int    `json:"db"`           //库名
	User         string `json:"user"`         //用户名
	Password     string `json:"password"`     //密码
	PoolSize     int    `json:"poolSize"`     //连接池最大socket连接数
	MinIdleConns int    `json:"minIdleConns"` //闲置连接数量
}

type Jaeger struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Kafka struct {
	Address    []string `json:"address"`     //地址列表
	Username   string   `json:"username"`    //用户名
	Password   string   `json:"password"`    //密码
	SaslEnable bool     `json:"sasl_enable"` //是否开启鉴权
	Consume    Consume  `json:"consume"`
}

type Consume struct {
	Group  string   `json:"group"`
	Topics []string `json:"topics"`
}

type Logs struct {
	Filename string `json:"filename"` // 日志文件地址
	Level    string `json:"level"`    // 日志等级
}

type Jwt struct {
	PrivateKey string `json:"private_key"` //私钥
	PublicKey  string `json:"public_key"`  //公钥
	Expires    int64  `json:"expires"`     //过期时间
}

type Consul struct {
	Id   string `json:"id"`   //不用配置，代码里md5生成了
	Host string `json:"host"` //地址
}

type Elasticsearch struct {
	Address  []string `json:"address"`  //地址列表
	Username string   `json:"username"` // Username for HTTP Basic Authentication.
	Password string   `json:"password"` // Password for HTTP Basic Authentication.
}

// Sls 阿里云sls日志配置
type Sls struct {
	Name      string `mapstructure:"name"   json:"name"`          //项目名
	Id        string `mapstructure:"id"     json:"id"`            //key id
	Secret    string `mapstructure:"secret" json:"secret"`        //密钥
	Host      string `mapstructure:"host"   json:"host"`          //地址
	StoreName string `mapstructure:"storeName" json:"store_name"` //文件名（仓库名）
}

// Sms 注册登录，收取验证码
type Sms struct {
	Url      string `mapstructure:"url" json:"url"`           // url
	Account  string `mapstructure:"account" json:"account"`   // 账号
	Password string `mapstructure:"password" json:"password"` // 密钥
}
