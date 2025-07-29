package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDns                string
	ProductTopic               string
	ProductWithFullImgSetTopic string
	ProductFiltered            string
	ProductFind                string
	SchemaRegistryUrl          string
	Brokers                    []string
}

func NewConfig() Config {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return Config{
		DatabaseDns:                getEnvStr("DATABASE_DSN", ""),
		ProductTopic:               getEnvStr("PRODUCT_TOPIC", "shop_products"),
		ProductWithFullImgSetTopic: getEnvStr("PRODUCT_WITH_FULL_IMG_SET_TOPIC", "analytic_products_with_full_img_set"),
		ProductFiltered:            getEnvStr("PRODUCT_FILTERED_TOPIC", "analytic_products_filtered"),
		ProductFind:                getEnvStr("PRODUCT_FIND_TOPIC", "analytic_products_find"),
		SchemaRegistryUrl:          getEnvStr("SCHEMA_REGISTRY_URL", "http://127.0.0.1:8081"),
		Brokers:                    strings.Split(getEnvStr("KAFKA_BROKERS", "127.0.0.1:9093"), ","),
	}
}

func getEnvInt(key string, def int) int {
	v, e := strconv.Atoi(getEnvStr(key, strconv.Itoa(def)))
	if e != nil {
		return def
	} else {
		return v
	}
}

func getEnvStr(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}
