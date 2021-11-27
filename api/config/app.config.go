package config

import (
	"github.com/techvisionus/golang-source/api/utils"
)

var POSTGRES string = utils.GetEnv("POSTGRES_CONFIG", "host=172.16.16.101 dbname=postgres user=postgres password=postgres port=31090 sslmode=disable")

var PORT string = utils.GetEnv("PORT", "3000")
