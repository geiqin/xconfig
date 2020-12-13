package client

import (
	"github.com/geiqin/xconfig/model"
	"github.com/micro/go-micro/v2/util/log"
	"os"
)

var confAddress string

func init() {
	confAddress = os.Getenv("SAAS_CONFIG_ADDRESS")
}

func GetAppConfig() *model.AppConfig {
	conf, _ := LoadSource(confAddress, "app")
	c := &model.AppConfig{}
	if err := conf.Scan(c); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func GetDatabaseConfig() *model.DatabaseConfig {
	conf, _ := LoadSource(confAddress, "database")
	c := &model.DatabaseConfig{}
	if err := conf.Scan(c); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func GetFilesystemConfig() *model.FilesystemConfig {
	conf, _ := LoadSource(confAddress, "filesystem")
	c := &model.FilesystemConfig{}
	if err := conf.Scan(c); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func GetPaymentConfig() *model.PaymentConfig {
	conf, _ := LoadSource(confAddress, "payment")
	c := &model.PaymentConfig{}
	if err := conf.Scan(c); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
