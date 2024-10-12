package TencentCos

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gsync/conf"
)

var (
	secretId, secretKey, defaultBucketUrl, appId string
)

func setDefaultGsyConfig() {
	viper.SetDefault("gsy.secretId", "")
	viper.SetDefault("gsy.secretKey", "")
	viper.SetDefault("gsy.defaultBucketUrl", "")
}

func loadGsyConfig() {
	secretId = viper.GetString("gsy.secretId")
	secretKey = viper.GetString("gsy.secretKey")
	defaultBucketUrl = viper.GetString("gsy.defaultBucketUrl")

	//fmt.Println(secretId, secretKey, defaultBucketUrl)
}

func init() {
	conf.InitConfig()
	setDefaultGsyConfig()
	loadGsyConfig()
	LoadAppId()
	OnWatchConfigFileForGsy()
}

func OnWatchConfigFileForGsy() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		loadGsyConfig()
	})
	go viper.WatchConfig()
}

func LoadAppId() {
	_, appId = GetBucketName(defaultBucketUrl)
}
