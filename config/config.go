package config

import (
	"college-learning-asynchronous-gamification/accessor"
	"college-learning-asynchronous-gamification/gcs"
)

type ConfigMap struct {
	HostUrl     string
	MySQLConfig accessor.SqlConfig
	GCSConfig   gcs.GCSConfig
}
