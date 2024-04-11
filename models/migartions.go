package models

import "TemplateService/database"

func Migrate() error {
	return database.BD.AutoMigrate()
}
