package config

import (
    "fmt"
    "os"

    "github.com/BurntSushi/toml"
)

var Env Config

type Config struct {
    EnvType      string // production, staging, development
    Adapters     map[string]Credential
    FeatureFlags struct {
        Flag bool
    }
    Logging struct {
        Level  string
        Format string
    }
    DB struct {
        Postgres Postgres
    }
}

type Postgres struct {
    Debug          bool
    Host           string
    Port           int
    Username       string
    Password       string
    DBName         string
    HasSSL         bool
    ClientEncoding string
    Timeout        int
    Timezone       string
}

type Credential struct {
    Auth struct {
        Username string
        Password string
    }
    BaseURL string
    Timeout int
}

func Init(path string) {
    var envPath string
    if path == "" {
        if envPath = os.Getenv("EXRATE_CONFIG_FILE_PATH"); envPath == "" {
            panic("config file path must be gived with flag or EXRATE_CONFIG_FILE_PATH env var")
        }
    }

    if _, err := toml.DecodeFile(path, &Env); err != nil {
        panic(fmt.Errorf("config file read error : %e", err))
    }
}

func (c *Config) IsDevelopment() bool {
    return c.EnvType == "development"
}
