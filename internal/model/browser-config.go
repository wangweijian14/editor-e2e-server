package model

type BrowserConfig struct {
	//若true不应用下面内容
	LocalBrowser       bool   `json:"localBrowser"`
	LauncherManager    string `json:"launcherManager"`
	DisableGpu         bool   `json:"disableGpu"`
	LaunchHeadfulMode  bool   `json:"launchHeadfulMode"`
	ServeMonitor       string `json:"serveMonitor"`
	StartMonitorServer bool   `json:"startMonitorServer"`
}
