package main

func isRunKind(s string) bool {
	return s == "core" || s == "reflex"
}

func parsePlayArgs(args []string) (drill, brief, runMath bool, runMode string) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--drill":
			drill = true
		case "--brief":
			brief = true
		case "--run-math":
			runMath = true
		case "--run-core5":
			runMode = "core"
		case "--run":
			if i+1 < len(args) && isRunKind(args[i+1]) {
				i++
				runMode = args[i]
			} else {
				runMode = "all"
			}
		}
	}
	return drill, brief, runMath, runMode
}
