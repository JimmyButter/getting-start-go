package request

import (
	"sync"

	"hertz_demo/middleware/nacos"
	"hertz_demo/middleware/nacos/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var instance *gorm.DB
var once sync.Once

func GetConnection() *gorm.DB {

	if instance == nil {

		once.Do(func() {
			context := nacos.GetConfig("database", "DEFAULT_GROUP")
			var config = &model.DatabaseConfig{}
			nacos.GetDecode(context, config)
			dsn := config.User + ":" + config.Password + "(" + config.Url + ")/" + config.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   "t_", // 表名前缀，`Article` 的表名应该是 `it_articles`
					SingularTable: true, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
				}})

			if err != nil {
				panic(err)
			}

			instance = db
		})
	}

	return instance
}

// db.AutoMigrate(&model.Candle{})
// c1 := model.Candle{
// 	Id:      1,
// 	InstId:  "BTC-USD-SWAP",
// 	Bar:     "1min",
// 	O:       100.00,
// 	H:       100.00,
// 	L:       100.00,
// 	C:       100.00,
// 	Vol:     100.00,
// 	Confirm: "0",
// 	Ts:      11343455,
// }

// db.Create(&c1)
