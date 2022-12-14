package panel

type Panel struct {
	config Configuration
}

var panelInstance *Panel = nil

func NewPanel(config Configuration) *Panel {
	panelInstance = &Panel{config: config}
	return panelInstance
}
