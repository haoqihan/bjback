package config

type Server struct {
}

type System struct {
}

type Mysql struct {
	Username     string `json:"username"`       // 设置用户名
	Password     string `json:"password"`       // 设置密码
	Path         string `json:"path"`           // 设置路径
	Dbname       string `json:"dbname"`         // 设置数据库名称
	Config       string `json:"config"`         // 设置配置
	MaxIdleConns int    `json:"max_idle_conns"` // 设置执行完闲置的链接
	MaxOpenConns int    `json:"max_open_conns"` // 设置最大连接数
	LogMode      bool   `json:"log_mode"`
}

type Sqlite struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
	Config   string `json:"config"`
	LogMode  bool   `json:"log_mode"`
}
