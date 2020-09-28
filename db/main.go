package db

import (
    "fmt"

    "github.com/faoztas/exrate/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB()  {
    var err error
    if DB, err = connectPostgresql(config.Env.DB.Postgres); err != nil {
        panic(err)
    }
    if config.Env.DB.Postgres.Debug {
        DB.Logger.LogMode(logger.Info)
    }
}

func connectPostgresql(psql config.Postgres) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(GetPsqlDSN(psql)), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func GetPsqlDSN(psql config.Postgres) string {
    sslmode := ""
    if !psql.HasSSL {
        sslmode = "sslmode=disable"
    }
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d host=%s client_encoding=%s connect_timeout=%d time_zone=%s %s",
        psql.Username, psql.Password, psql.DBName, psql.Port, psql.Host, psql.ClientEncoding, psql.Timeout, psql.Timezone, sslmode)
    return dsn
}

func Migrate() {
    // DB.AutoMigrate(&model.Todo{})
}