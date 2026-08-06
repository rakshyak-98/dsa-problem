package main

func isRunKind(s string) bool {
	return s == "core" || s == "reflex"
}

func parseBackendArgs(args []string) (drill, catalog, cram, setup bool, runMode string) {
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--drill":
			drill = true
		case "--catalog":
			catalog = true
		case "--cram":
			cram = true
		case "--setup":
			setup = true
		case "--run":
			if i+1 < len(args) && isRunKind(args[i+1]) {
				i++
				runMode = args[i]
			} else {
				runMode = "all"
			}
		}
	}
	return drill, catalog, cram, setup, runMode
}
