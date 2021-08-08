package engine

import "sync"

// EngineMutex only locks and unlocks on engine creation functions
// as engine modifies global files, this protects the main bot creation
var EngineMutex sync.Mutex

// Engine will have all the Engine based
type Engine struct {
}

// Config will have all the engine config
type Config struct {
}
