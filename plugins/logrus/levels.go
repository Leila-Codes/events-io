package logrus

import "github.com/sirupsen/logrus"

type LevelPicker[IN interface{}] func(IN) logrus.Level

func AllInfo[IN interface{}](IN) logrus.Level {
	return logrus.InfoLevel
}

func AllWarn[IN interface{}](IN) logrus.Level {
	return logrus.WarnLevel
}

func AllError[IN interface{}](IN) logrus.Level {
	return logrus.ErrorLevel
}

func AllDebug[IN interface{}](IN) logrus.Level {
	return logrus.DebugLevel
}

func AllFatal[IN interface{}](IN) logrus.Level {
	return logrus.FatalLevel
}
