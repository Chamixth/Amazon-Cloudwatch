package dto
type Dashboard struct {
	DashboardName string       `json:"DashboardName"`
	DashboardBody DashboardBody `json:"DashboardBody"`
}

type DashboardBody struct {
	Start          string     `json:"start"`
	PeriodOverride string     `json:"periodOverride"`
	Widgets        []Widget   `json:"widgets"`
}

type Widget struct {
	Type       string      `json:"type"`
	X          int         `json:"x"`
	Y          int         `json:"y"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	Properties WidgetProps `json:"properties"`
}

type WidgetProps struct {
	Metrics   [][]string `json:"metrics"`
	Period    int        `json:"period"`
	Stat      string     `json:"stat"`
	Region    string     `json:"region"`
	Title     string     `json:"title"`
	LiveData  bool       `json:"liveData"`
	Legend    Legend     `json:"legend"`
}

type Legend struct {
	Position string `json:"position"`
}