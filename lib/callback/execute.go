package callback

import (
	"msh/lib/config"
	"msh/lib/errco"
	"os/exec"
)

func Execute() {
	if !config.ConfigRuntime.Msh.CallbackEnabled || len(waiters) == 0 {
		return
	}
	// Defines the Slice capacity to match the Map elements count
	waitersArr := make([]string, 0, len(waiters))
	for tx := range waiters {
		waitersArr = append(waitersArr, tx)
	}
	waiters = map[string]bool{}

	go func() {
		var args []string
		if len(config.ConfigRuntime.Commands.CallBack) > 1 {
			args = append(args, config.ConfigRuntime.Commands.CallBack[1:]...)
		}
		args = append(args, waitersArr...)

		_, err := exec.Command(config.ConfigRuntime.Commands.CallBack[0], args...).Output()
		if err != nil {
			errco.LogMshErr(errco.NewErr(-1, errco.LVL_B, "callback", err.Error()))
		}
	}()

}
