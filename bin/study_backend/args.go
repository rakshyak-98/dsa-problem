package main

func parseBackendArgs(args []string) (micro, run, catalog, cram, setup, specialty bool) {
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
		case "--specialty":
			specialty = true
		}
	}
	return micro, run, catalog, cram, setup, specialty
}
