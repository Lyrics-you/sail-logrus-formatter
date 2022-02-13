package hooks

import (
	"github.com/sirupsen/logrus"
)

// ContextHook for log the call context
type positionHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// PositionHook use to make an hook
// Set the recursion depth to 5
func PositionHook(levels ...logrus.Level) logrus.Hook {
	hook := positionHook{
		Field:  "source",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook positionHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implement fire
func (hook positionHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = FindCaller(hook.Skip)
	return nil
}
