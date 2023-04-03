package logger

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lggr *zap.SugaredLogger

var rawJSON = []byte(`{
	"level": "` + os.Getenv("PREDESTINATION_LOG_LEVEL") + `",
	"encoding": "json",
	  "initialFields": {
		  "application": "amaginow"
		},
	"encoderConfig": {
		"levelKey": "level",
		"messageKey": "message",
		"levelEncoder": "lowercase",
		"nameKey":"name",
		"stacktraceKey":"stack"
	}
  }`)

func NewLogger(moduleName string, initialFields map[string]interface{}, filename string) *zap.SugaredLogger {
	config := setUpConfig(moduleName, initialFields)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	lg, _ := config.Build()
	lggr = lg.Sugar()
	return lggr
}

func setUpConfig(moduleName string, initialFields map[string]interface{}) (config zap.Config) {
	json.Unmarshal(rawJSON, &config)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.InitialFields["module"] = moduleName
	if initialFields != nil {
		for k, v := range initialFields {
			config.InitialFields[k] = v
		}
	}
	return config
}
