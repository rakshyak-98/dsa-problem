package main

func parseBackendArgs(args []string) (drill, run, catalog, cram, setup bool) {
	for _, a := range args {
		switch a {
		case "--drill":
			drill = true
		case "--run":
			run = true
		case "--catalog":
			catalog = true
		case "--cram":
			cram = true
		case "--setup":
			setup = true
		}
	}
	return drill, run, catalog, cram, setup
}
