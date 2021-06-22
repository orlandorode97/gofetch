package main

type OSInformer interface {
	GetOSVersion() (string, error)
	GetHostname() (string, error)
	GetUptime() (string, error)
	GetNumberPackages() (string, error)
	GetShellInformation() (string, error)
	GetResolution() (string, error)
	GetDesktopEnvironment() (string, error)
	GetTerminalInfo() (string, error)
	GetCPU() (string, error)
	GetGPU() (string, error)
	GetMemoryUsage() (string, error)
}

type OSInformation struct {
	Name               string
	OS                 string
	Host               string
	Uptime             string
	Packages           string
	Shell              string
	Resolution         string
	DesktopEnvironment string
	Terminal           string
	CPU                string
	GPU                string
	Memory             string
}
