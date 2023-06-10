package setup

import (
	"database/sql"
	"fmt"
	"github.com/arvians-id/go-portfolio/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func NewPostgresSQLGormTest(configuration config.Config) (*gorm.DB, error) {
	username := configuration.Get("DB_USERNAME_TEST")
	password := configuration.Get("DB_PASSWORD_TEST")
	host := configuration.Get("DB_HOST_TEST")
	port := configuration.Get("DB_PORT_TEST")
	database := configuration.Get("DB_DATABASE_TEST")
	sslMode := configuration.Get("DB_SSL_MODE_TEST")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, err
	}

	_, err = databasePooling(configuration, dbPool)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TearDownDBTest(db *gorm.DB) error {
	err := db.Exec("DELETE FROM contacts").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM educations").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM certificates").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM users").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM category_skills").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM work_experiences").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM projects").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM skills").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM work_experience_skill").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM project_skill").Error
	if err != nil {
		return err
	}

	err = db.Exec("DELETE FROM project_images").Error
	if err != nil {
		return err
	}

	return nil
}

func databasePooling(configuration config.Config, db *sql.DB) (*sql.DB, error) {
	setMaxIdleConns, err := strconv.Atoi(configuration.Get("DB_POOL_MIN_TEST"))
	if err != nil {
		return nil, err
	}
	setMaxOpenConns, err := strconv.Atoi(configuration.Get("DB_POOL_MAX_TEST"))
	if err != nil {
		return nil, err
	}
	setConnMaxIdleTime, err := strconv.Atoi(configuration.Get("DB_MAX_IDLE_TIME_TEST"))
	if err != nil {
		return nil, err
	}
	setConnMaxLifetime, err := strconv.Atoi(configuration.Get("DB_MAX_LIFE_TIME_TEST"))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(setMaxIdleConns)                                    // minimal connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Second) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Second)

	return db, nil
}
