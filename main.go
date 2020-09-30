package main

import (
    "flag"

    "github.com/faoztas/exrate/cache"
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/db"
)

func main()  {
    var configFilePath string

    flag.StringVar(&configFilePath, "config", "config/config.toml", "configuration file path")
    flag.Parse()

    config.Init(configFilePath)

    db.InitDB()
    db.Migrate(db.DB)

    cache.Load()

}