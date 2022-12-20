package web

type webTemplateData struct {
	Containers     []webContainerInfo
	GlobalWarnings []string
}

type webContainerInfo struct {
	Index         int
	Hostname      string
	DisplayLabels map[string]string
	Warnings      []string
}

type reportData struct {
	Index    int
	Graphs   map[string]map[string]string
	Hostname string
}
