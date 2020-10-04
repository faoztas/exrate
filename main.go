package main

import (
    "flag"

    "github.com/faoztas/exrate/apps/logging"
    "github.com/faoztas/exrate/cache"
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/db"
    job "github.com/faoztas/exrate/jobs"
)

func main()  {
    var configFilePath string

    flag.StringVar(&configFilePath, "config", "config/config.toml", "configuration file path")
    flag.Parse()

    config.Init(configFilePath)

    logging.InitLogging()

    db.InitDB()
    db.Migrate(db.DB)

    cache.Load()

    job.Jobs()

}