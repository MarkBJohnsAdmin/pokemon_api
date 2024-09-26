package upload

import (
	"fmt"
	"server/database/private"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//	===============================================================================================

type Metadata struct {
	Version  string
	Author   string
	Passcode string
}

const dbVersion = "1"

//	===============================================================================================

func InitDatabase() (*gorm.DB, error) {

	dsn := private.GetDSN("")
	dbName := private.GetDbName()

	//	Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	//	Check for database
	exists, err := databaseExists(db, dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to check for database: %w", err)
	}

	if !exists {
		err = createDataBase(dsn, dbName)
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
	} else {
		dsnWithDB := private.GetDSN(dbName)
		db, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to existing database: %w", err)
		}

		valid, err := validMetadata(db)
		if err != nil {
			return nil, fmt.Errorf("failed to validate metadata: %w", err)
		}

		if !valid {
			err = createDataBase(dsn, dbName)
			if err != nil {
				return nil, fmt.Errorf("failed to recreate database: %w", err)
			}
		}
	}

	dsnWithDB := private.GetDSN(dbName)
	db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to final database: %w", err)
	}

	return db, nil
}

//	===============================================================================================

func createDataBase(dsn, dbName string) error {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL for database creation: %w", err)
	}

	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)).Error; err != nil {
		return fmt.Errorf("failed to drop existing database: %w", err)
	}

	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error; err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	dsnWithDB := private.GetDSN(dbName)
	db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return fmt.Errorf("failed to reconnect to new database: %w", err)
	}

	if err := db.AutoMigrate(&Metadata{}); err != nil {
		return fmt.Errorf("failed to migrate metadata table: %w", err)
	}
	author, passcode := private.GetMetadata()
	metadata := Metadata{Version: dbVersion, Author: author, Passcode: passcode}
	if err := db.Create(&metadata).Error; err != nil {
		return fmt.Errorf("failed to insert initial metadata: %w", err)
	}

	tableData := GetTableData()
	err = UploadAllTables(db, tableData)
	if err != nil {
		return fmt.Errorf("failed to upload Pokemon data: %w", err)
	}

	return nil
}

//	-----------------------------------------------------------------------------------------------

func databaseExists(db *gorm.DB, dbName string) (bool, error) {

	var count int

	query := "SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"

	err := db.Raw(query, dbName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//	-----------------------------------------------------------------------------------------------

func validMetadata(db *gorm.DB) (bool, error) {

	var count int

	query := "SELECT COUNT(*) " +
		"FROM information_schema.tables " +
		"WHERE table_schema = ? " +
		"AND table_name = 'metadata'"

	err := db.Raw(query, private.GetDbName()).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	var metadata Metadata
	err = db.Table("metadata").First(&metadata).Error
	if err != nil {
		return false, fmt.Errorf("failed to read metadata: %w", err)
	}

	author, passcode := private.GetMetadata()

	if metadata.Version != dbVersion || metadata.Author != author || metadata.Passcode != passcode {
		return false, nil
	}

	return true, nil
}
