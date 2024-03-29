package web

// webTemplateData is the data passed to the web template
type webTemplateData struct {
	Containers     []webContainerInfo
	GlobalWarnings []string
}

// webContainerInfo is the data for a single container
type webContainerInfo struct {
	Index         int
	Hostname      string
	DisplayLabels map[string]string
	Warnings      []string
}
