package callback

import "msh/lib/config"

var waiters = map[string]bool{}

func Collect(playerName string) {
	if config.ConfigRuntime.Msh.CallbackEnabled {
		waiters[playerName] = true
	}
}
