package main

import (
    "flag"

    "github.com/faoztas/exrate/config"
)

func main()  {
    var configFilePath string

    flag.StringVar(&configFilePath, "config", "config/config.toml", "configuration file path")
    flag.Parse()

    config.Init(configFilePath)
}