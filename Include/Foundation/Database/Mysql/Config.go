package Mysql

type MysqlConfig struct {
	DatabasePath    string
	Port            int
	DatabaseName    string
	UserName        string
	Password        string
	MaxOpenConns    int
	MaxConnLifetime int
	MaxIdleConns    int
	MaxIdleLifetime int
}
