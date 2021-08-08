package _default

// BaseBroker stores the individual broker information
type BaseBroker struct {
	Name                 string
	IsWebSocketSupported bool
	Enabled              bool
	//Verbose                       bool
	//LoadedByConfig                bool
	//SkipAuthCheck                 bool
	//HTTPTimeout                   time.Duration
	//HTTPUserAgent                 string
	//HTTPRecording                 bool
	//HTTPDebugging                 bool
	//WebsocketResponseCheckTimeout time.Duration
	//WebsocketResponseMaxLimit     time.Duration
	//WebsocketOrderBookBufferLimit int64
	//SettingsMutex                 sync.RWMutex
}
