package main

/**
根据表生成结构体
*/
import (
	"github.com/qmhball/db2gorm/gen"
)

// mysql 数据生成结构体
func mysqlGen(tblName string) {
	dsn := "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8&parseTime=true&loc=Local"
	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}

func main() {
	mysqlGen("t_order_detail_jst")
}
