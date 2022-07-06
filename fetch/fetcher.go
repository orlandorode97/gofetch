package fetch

type Versioner interface {
	GetOSVersion() string
}

type Namer interface {
	GetName() string
}

type Timer interface {
	GetUptime() string
}

type Packager interface {
	GetNumberPackages() string
}

type Sheller interface {
	GetShellInformation() string
}

type Resolusioner interface {
	GetResolution() string
}

type Environment interface {
	GetDesktopEnvironment() string
}

type Terminal interface {
	GetTerminalInfo() string
}

type CPU interface {
	GetCPU() string
}

type GPU interface {
	GetGPU() string
}

type Usager interface {
	GetMemoryUsage() string
}

type Kernel interface {
	GetKernelVersion() string
}

type Fetcher interface {
	Versioner
	Namer
	Timer
	Packager
	Sheller
	Resolusioner
	Environment
	Terminal
	CPU
	GPU
	Usager
	Kernel
}
