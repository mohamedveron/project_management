package domains

type EnumProjectState string

const (
	EnumProjectStatePlanned EnumProjectState = "planned"
	EnumProjectStateActive  EnumProjectState = "active"
	EnumProjectStateDone    EnumProjectState = "done"
	EnumProjectStateFailed  EnumProjectState = "failed"
)

type Project struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Owner        Employee         `json:"owner"`
	Progress     float32          `json:"progress"`
	State        EnumProjectState `json:"state" validate:"oneof=planned active done failed"`
	Participants []Employee       `json:"participants"`
}
