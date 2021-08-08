package engine

var Bots map[string]*Engine

// New starts a new engine
func New() (*Engine, error) {
	EngineMutex.Lock()
	defer EngineMutex.Unlock()
	var b Engine
	return &b, nil
}
