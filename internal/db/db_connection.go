package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	dbConfig "github.com/KrizzMU/coolback-alkol/internal/config/dbConf"
	"github.com/KrizzMU/coolback-alkol/internal/core"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&core.Sessions{})

	db.AutoMigrate(&core.Course{})

	db.AutoMigrate(&core.Module{})
	db.Model(&core.Module{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Lesson{})
	db.Model(&core.Lesson{}).AddForeignKey("module_id", "modules(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Email{})

	return db
}
