package db

import (
	"gorm.io/gorm"
)

type migrator struct {
	db *gorm.DB
}

func (m *migrator) CreateTables(dst ...interface{}) error {
	for _, d := range dst {
		if m.db.Migrator().HasTable(d) {
			continue
		}

		if err := m.db.Migrator().CreateTable(d); err != nil {
			return err
		}
	}

	return nil
}

func newMigrator(db *gorm.DB) *migrator {
	return &migrator{db}
}
