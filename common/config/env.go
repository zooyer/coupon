package config

func IsDebug() bool {
	return env == "debug"
}

func IsTest() bool {
	return env == "test"
}

func IsSim() bool {
	return env == "sim"
}

func IsProd() bool {
	return env == "prod"
}
