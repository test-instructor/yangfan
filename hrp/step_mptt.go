package hrp

import (
	"encoding/json"
	"fmt"
)

type MqttMessageType string

const (
	MTJson   MqttMessageType = "json"
	MtString MqttMessageType = "string"
)

type MqttStep struct {
	step *TStep
}

func (m *MqttStep) Name() string {
	return m.step.Name
}

func (m *MqttStep) Type() StepType {
	return stepTypeGRPC
}

func (m *MqttStep) Struct() *TStep {
	return m.step
}

func (m *MqttStep) Run(r *SessionRunner) (*StepResult, error) {
	return runStepMQTT(r, m.step)
}

func runStepMQTT(r *SessionRunner, step *TStep) (*StepResult, error) {
	return nil, nil
}

type MQTT struct {
	Host         string          `json:"host" yaml:"host"`
	Client       string          `json:"client" yaml:"client"`
	Username     string          `json:"username" yaml:"username"`
	Password     string          `json:"password" yaml:"password"`
	Topic        string          `json:"topic" yaml:"topic"`
	Body         string          `json:"Body" yaml:"Body"`
	BodyType     MqttMessageType `json:"body_type" yaml:"bodyType"`
	ResponseType MqttMessageType `json:"response_type" yaml:"responstType"`
}

func convertMqttMessage(msgType MqttMessageType, msg string) (interface{}, error) {
	switch msgType {
	case MTJson:
		var jsonObj interface{}
		err := json.Unmarshal([]byte(msg), &jsonObj)
		if err != nil {
			return nil, err
		}
		msgBytes, err := json.Marshal(msg)
		return msgBytes, err
	case MtString:
		return msg, nil
	default:
		return nil, fmt.Errorf("Unsupported body type: %s", msg)
	}
}
