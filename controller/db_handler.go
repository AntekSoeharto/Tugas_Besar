package controller

import (
	. "github.com/Tugas_Besar/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect to database
func connect() *gorm.DB {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // Slow SQL threshold
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      false,       // Disable color
	//	},
	//)

	dsn := "root:@tcp(127.0.0.1:3306)/db_tubes_pbp?charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}

// Migrate for automation database tabel creation
func Migrate() {
	db := connect()

	//Close() method, unsupperted since GORM v2
	//defer db.Close

	_ = db.AutoMigrate(&User{}, &Film{}, &RiwayatUser{})

}
