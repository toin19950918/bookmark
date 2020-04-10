package gormdao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
	"github.com/jinzhu/gorm"
	"github.com/robin019/bookmark/src/utils/config"
	"github.com/robin019/bookmark/src/utils/logger"
	"time"
)

var (
	db       *gorm.DB
	interval = config.Get("db.interval").(int)
	dialect  = config.Get("db.dialect").(string)
	source   = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Get("db.user"),
		config.Get("db.password"),
		config.Get("db.host"),
		config.Get("db.port"),
		config.Get("db.database"),
		config.Get("db.flag"),
	)
)

// DB return database instance
func init() {
	connect(dialect, source)
}

func DB() *gorm.DB {
	if db == nil {
		connect(dialect, source)
	}
	return db
}

// Close close database connection
func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			logger.Debug().Error(err.Error())
		}
	}
}

func connect(dialect string, source string) {
	conn, err := gorm.Open(dialect, source)

	if err != nil {
		logger.Debug().Error(err.Error())
		fmt.Println("MySQL連線失敗，三秒後重新連線")
		time.Sleep(3 * time.Second)
		connect(dialect, source)
	} else {
		conn.DB().SetMaxOpenConns(30)
		conn.DB().SetMaxIdleConns(10)
		conn.DB().SetConnMaxLifetime(90 * time.Second)
		conn.SingularTable(true)
		conn.BlockGlobalUpdate(true)

		db = conn
	}
}
