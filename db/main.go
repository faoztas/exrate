package db

import (
    "fmt"

    "github.com/faoztas/exrate/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB()  {
    if config.Env.DB.Postgres.URL == "" {
        config.Env.DB.Postgres.URL = buildPostresqlURL(config.Env)
    }
}

func connectPostgresql(env config.Config) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.New(postgres.Config{DSN: buildPostresqlURL(env)}), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func GetDSN(d config.Config) string {
    username := d.DB.Postgres.Username
    password := d.DB.Postgres.Password
    name := d.DB.Postgres.DBName
    host := fmt.Sprintf("%s:%d", d.DB.Postgres.Host, d.DB.Postgres.Port)
    DSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
        host, username, name, password)
    return DSN
}

func buildPostresqlURL(d config.Config) string {
    format := "postgres://%s:%s@%s:%d/%s?client_encoding=%s&connect_timeout=%d"
    if !d.DB.Postgres.HasSSL {
        format += "&sslmode=disable"
    }
    return fmt.Sprintf(format, d.DB.Postgres.Username,
        d.DB.Postgres.Password, d.DB.Postgres.Host,
        d.DB.Postgres.Port, d.DB.Postgres.DBName,
        d.DB.Postgres.ClientEncoding, d.DB.Postgres.Timeout)
}