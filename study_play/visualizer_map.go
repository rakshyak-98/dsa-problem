package main

import "fmt"

var visualizerLinks = map[string]string{
	"01_arrays_reflex":       "visualizer/index.html#reverse-string",
	"02_hashing_reflex":      "visualizer/index.html#two-sum",
	"03_two_pointers_reflex": "visualizer/index.html#two-pointers",
	"04_binary_search_reflex": "visualizer/index.html#binary-search",
	"05_trees_stacks_reflex": "visualizer/index.html#valid-parentheses",
	"06_dp_reflex":           "visualizer/index.html#climbing-stairs",
	"07_graphs_reflex":       "visualizer/index.html#num-islands",
	"08_heap_reflex":         "visualizer/index.html#kth-largest",
	"09_backtrack_reflex":    "visualizer/index.html#subsets",
	"core5":                  "visualizer/index.html#two-sum",
}

func printVisualizerLink(drillFile string) {
	link, ok := visualizerLinks[drillFile]
	if !ok {
		return
	}
	fmt.Printf("\n── VISUALIZER ─────────────────────────────────────────\n  %s\n", link)
}
