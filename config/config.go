package config

import (
    "io/ioutil"
    "fmt"
    "os"
    "encoding/json"
)

type DBConf struct {
    Host string `json:"host"`
    Username string `json:"username"`
    Password string `json:"password"`
    Database string `json:"database"`
}

type Config struct {
    DBConf            `json:"db"`
    Host       string `json:"host"`
    Port       string `json:"port"`
    LogConfig  string `json:"log"`
    StaticPath string `json:"static"`
    APIPrefix  string `json:"api_prefix"`
}

var Conf Config

func init() {
    file, err := ioutil.ReadFile("config/config.json")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // unmarshal json into a config struct
    ErrJSONUnmarshal := json.Unmarshal(file, &Conf)
    if ErrJSONUnmarshal != nil {
        fmt.Println(ErrJSONUnmarshal)
        os.Exit(1)
    }
}
