package main

func parseBackendArgs(args []string) (micro, run, catalog, cram, setup bool) {
	for _, a := range args {
		switch a {
		case "--micro":
			micro = true
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
	return micro, run, catalog, cram, setup
}
