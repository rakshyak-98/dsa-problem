package main

func parseReadArgs(args []string) (micro, run, catalog bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for _, a := range args {
		switch a {
		case "--micro":
			micro = true
		case "--run":
			run = true
		case "--catalog":
			catalog = true
		}
	}
	return micro, run, catalog
}
