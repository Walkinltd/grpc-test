package logger

import (
	"fmt"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

func Setup(locally bool, name string) *zap.Logger {
	setupFunc := zapdriver.NewProduction
	if locally {
		setupFunc = zap.NewDevelopment
	}

	l, err := setupFunc()
	if err != nil {
		panic(fmt.Errorf("failed to initialise logger: %w", err))
	}
	return l.Named(name)
}
