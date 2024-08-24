package logrus

import "github.com/sirupsen/logrus"

func loggerThread[IN interface{}](
	input chan IN,
	logger *logrus.Logger,
	level LevelPicker[IN],
	fields ...FieldGenerator[IN],
) {
	for event := range input {
		var e *logrus.Entry
		for _, fieldGen := range fields {
			if e == nil {
				e = logger.WithFields(fieldGen(event))
			} else {
				e = e.WithFields(fieldGen(event))
			}
		}

		e.Log(
			level(event),
			event,
		)
	}
}

func Logger[IN interface{}](
	input chan IN,
	logger *logrus.Logger,
	level LevelPicker[IN],
	fields FieldGenerator[IN],
) {
	go loggerThread(input, logger, level, fields)
}
