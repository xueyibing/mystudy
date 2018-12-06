package main

import (
	"fmt"
	"os"
)

const (
	ACC_HOST = iota + 1
	ACC_USER
	ACC_PASS
	ACC_DB_HOST
	GAEA_DB_HOST
	RABBITMQ_HOST
	SUPERVISOR_HOST
	IAM_HOST
	REDIS_HOST
)

func GetValue(key int) string {
	switch key {
	case ACC_HOST:
		return os.Getenv("TEST_ACC_HOST")
	case ACC_USER:
		return os.Getenv("TEST_ACC_USER")
	case ACC_PASS:
		return os.Getenv("TEST_ACC_PASS")
	case ACC_DB_HOST:
		return os.Getenv("TEST_ACC_DB_HOST")
	case GAEA_DB_HOST:
		return os.Getenv("TEST_GAEA_DB_HOST")
	case RABBITMQ_HOST:
		return os.Getenv("TEST_RABBITMQ_HOST")
	case SUPERVISOR_HOST:
		return os.Getenv("TEST_SUPERVISOR_HOST")
	case IAM_HOST:
		return os.Getenv("TEST_IAM_HOST")
	case REDIS_HOST:
		return os.Getenv("TEST_REDIS_HOST")
	default:
		return ""
	}
}



func main()  {

	IamHost := GetValue(IAM_HOST)
	RedisHost := GetValue(REDIS_HOST)

	fmt.Println("IamHost:",IamHost)
	fmt.Println("RedisHost:",RedisHost)

	
}
