package logrus

import "github.com/sirupsen/logrus"

type FieldGenerator[IN interface{}] func(IN) logrus.Fields

func DefaultFields[IN interface{}](fields logrus.Fields) FieldGenerator[IN] {
	return func(i IN) logrus.Fields {
		return fields
	}
}
