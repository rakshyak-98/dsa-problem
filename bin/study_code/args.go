package main

func parseReadArgs(args []string) (drill, run, runMath, catalog bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for _, a := range args {
		switch a {
		case "--drill":
			drill = true
		case "--run":
			run = true
		case "--run-math":
			runMath = true
		case "--catalog":
			catalog = true
		}
	}
	return drill, run, runMath, catalog
}
