package panel

type Panel struct {
	Config Configuration
}

var panelInstance *Panel = nil

func NewPanel(config Configuration) *Panel {
	panelInstance = &Panel{Config: config}
	return panelInstance
}

func GetPanelInstance() *Panel {
	//TODO
	return panelInstance
}
