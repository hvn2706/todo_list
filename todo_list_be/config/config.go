package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"todo_list/pkg/constants"

	"github.com/spf13/viper"
	"todo_list/logger"
	"todo_list/pkg/common"
)

const (
	tmpConfigFileName    = "config.tmp.yml"
	actualConfigFileName = "config.yml"
)

// GlobalCfg global variable to access app configuration without passing it around
var GlobalCfg Config

// ===== Init structs =====

// ServerListen for specifying host & port
type ServerListen struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

// ServerConfig for configure HTTP host & port
type ServerConfig struct {
	HTTP ServerListen `mapstructure:"http"`
}

type Database struct {
	MySQLConfig     MySQLConfig `mapstructure:"mysql"`
	MySQLTestConfig MySQLConfig `mapstructure:"mysql_test"`
}

type MySQLConfig struct {
	Host string `mapstructure:"db_host"`
	Port string `mapstructure:"db_port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
	Name string `mapstructure:"db_name"`

	MaxOpenCons        int   `mapstructure:"max_open_cons"`
	MaxIdleCons        int   `mapstructure:"max_idle_cons"`
	ConnMaxIdleTimeSec int64 `mapstructure:"conn_max_idle_time_sec"`
	ConnMaxLifetimeSec int64 `mapstructure:"conn_max_life_time_sec"`
}

// Config for app configuration
type Config struct {
	Server   ServerConfig        `mapstructure:"server"`
	Logger   logger.LoggerConfig `mapstructure:"logger"`
	Database Database            `mapstructure:"database"`
}

// ===== Util func =====

func (s ServerListen) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// ListenString for listen to host and port
func (s ServerListen) ListenString() string {
	return fmt.Sprintf(":%d", s.Port)
}

func Load() {
	if !common.CheckIfFileExist(actualConfigFileName) {
		err := os.Rename(tmpConfigFileName, actualConfigFileName)
		if err != nil {
			fmt.Println("Error renaming file:", err)
		} else {
			fmt.Println("File renamed successfully.")
			defer os.Rename(actualConfigFileName, tmpConfigFileName)
		}
	}

	vip := viper.New()
	vip.SetConfigName(constants.ConstConfig)
	vip.SetConfigType(constants.Yml)
	vip.AddConfigPath(constants.RootPath) // ROOT

	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}

	for _, key := range vip.AllKeys() {
		var (
			js     interface{}
			val    = vip.Get(key)
			valStr = fmt.Sprintf("%v", val)
		)

		err := json.Unmarshal([]byte(valStr), &js)

		if err != nil {
			vip.Set(key, val)
		} else {
			vip.Set(key, js)
		}
	}

	fmt.Printf("===== Config file used: %+v \n", vip.ConfigFileUsed())

	GlobalCfg = Config{}
	err = vip.Unmarshal(&GlobalCfg)
	if err != nil {
		panic(err)
	}
	return
}
