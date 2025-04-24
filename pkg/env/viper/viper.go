package envViper

import (
	"bytes"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/novando/go-ska/pkg/logger"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// InitViper
// Initialize Viper to use the config file as env variable
func InitViper(path string, l ...*logger.Logger) (err error) {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	var configName string
	splitPaths := strings.Split(path, "/")
	if len(splitPaths) > 0 {
		for i := 0; i < len(splitPaths); i++ {
			configName = splitPaths[i]
		}
	}
	splitNames := strings.Split(configName, ".")
	if len(splitNames) < 2 {
		err = errors.New("failed to parse config name")
		log.Fatalf(err.Error())
		return
	}
	formatName := splitNames[len(splitNames)-1]
	viper.SetConfigName(strings.TrimRight(configName, "."+formatName))
	viper.SetConfigType(formatName)
	viper.AddConfigPath(strings.TrimRight(path, configName))
	if err = viper.ReadInConfig(); err != nil {
		logger.Call().Warnf("Configs file: %v", err)
		return
	}

	if err = viper.Unmarshal(&Env); err != nil {
		logger.Call().Warnf("Configs unmarshar error: %v", err)
	}
	return
}

// InitRemoteViper
// Initialize Viper using remote config
func InitRemoteViper(user, pass, url string, l ...*logger.Logger) (err error) {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	client := resty.New()
	res, err := client.R().
		SetBasicAuth(user, pass).
		Get(url)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		if res.StatusCode() == http.StatusUnauthorized {
			err = errors.New("wrong config's credential")
		}
		log.Errorf(res.String())
		return errors.New(res.String())
	}
	viper.SetConfigType("json")
	if err = viper.ReadConfig(bytes.NewReader(res.Body())); err != nil {
		logger.Call().Warnf("Remote configs: %v", err)
		return
	}

	if err = viper.Unmarshal(&Env); err != nil {
		logger.Call().Warnf("Configs unmarshar error: %v", err)
	}
	return
}
