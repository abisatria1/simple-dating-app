package gorm

type DbConfig struct {
	Dsn           string
	RetryInterval int
	MaxIdleCon    int
	MaxCon        int
}
