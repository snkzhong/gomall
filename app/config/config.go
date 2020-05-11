package config

type Config struct {
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type Log struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Sql Sql `mapstructure:"sql" json:"sql" yaml:"sql"`
}

type App struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	Formatter string `mapstructure:"formatter" json:"formatter" yaml:"formatter"`
	File string `mapstructure:"file" json:"file" yaml:"file"`
}

type Sql struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	File string `mapstructure:"file" json:"file" yaml:"file"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Config string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	SingularTable bool `mapstructure:"singular-table" json:"singular-table" yaml:"singular-table"`
	LogMode bool `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
}

type Redis struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db int `mapstructure:"db" json:"db" yaml:"db"`
}
