package main

func parseDailyArgs(args []string) (passArgs []string, run bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for _, a := range args {
		switch a {
		case "--micro", "--run":
			passArgs = append(passArgs, a)
			if a == "--run" {
				run = true
			}
		}
	}
	return passArgs, run
}
