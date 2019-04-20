package dbTools

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var Eloquent *gorm.DB

func init() {
	var err error
	//				gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	Eloquent, err = gorm.Open("mysql", "Hitchhiker:Hitch974412@tcp(139.196.72.249:9090)/datacentre?charset=utf8&parseTime=True&loc=Local&timeout=2000ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}