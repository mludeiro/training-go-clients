package database

import (
	"log"
	"time"
	"training-go-clients/entity"
	"training-go-clients/tools"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.Client{})
	return db
}

func (db *Database) CreateSampleData() *Database {
	carlos := entity.Client{Name: "Carlos"}
	laura := entity.Client{Name: "Laura"}
	pedro := entity.Client{Name: "Pedro"}

	db.GetDB().Create(&carlos).Create(&laura).Create(&pedro)

	return db
}

func (db *Database) InitializePostgress() *Database {
	dsn := "host=localhost user=postgres password=postgres dbname=go_micro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func (db *Database) InitializeMySQL() *Database {
	dsn := "root:@tcp(127.0.0.1:3306)/go_micro?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func (db *Database) InitializeSqlite() *Database {

	DB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func createGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			tools.GetLogger(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			}),
	}
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) GetQueryDB(query entity.Query) *QueryExecutor {
	tx := db.GetDB()

	for _, fetch := range query.Fetchs {
		tx = tx.Preload(fetch)
	}

	for _, cond := range query.Conditions {
		switch cond.Comparator {
		case "eq":
			tx = tx.Where(cond.Field, cond.Value)
		case "lk":
			tx = tx.Where(cond.Field, cond.Value)
		}
	}

	for _, order := range query.OrderBy {
		tx = tx.Order(order)
	}

	return &QueryExecutor{trx: tx, pageSize: query.GetPageSize(), pageNumber: query.PageNumber}
}
