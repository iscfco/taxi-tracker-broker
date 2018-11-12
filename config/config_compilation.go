package config

import (
	"SW-CommonTools/constants"
	"SW-CommonTools/constants/config"
)

/*
const (
	DbEngine              = constants.DbPsqlEngine
	DbConnectionEnv       = constants.DbEnvProduction
	PrintLog              = true
	ImgStorageDestination = config.ImgStorageDestinationDisk
	//ImgStorageDestination = config.ImgStorageDestinationAwsS3
	AwsKeySource          = AwsKeyFromTestVars
	ApiEnv                = constants.ApiEnvProduction // it define the port
)
*/

const (
	DbEngine              = constants.DbPsqlEngine
	DbConnectionEnv       = constants.DbEnvHub
	PrintLog              = true
	ImgStorageDestination = config.ImgStorageDestinationDisk
	AwsKeySource          = AwsKeyFromTestVars
	ApiEnv                = constants.ApiEnvHub // it define the port
)