package main

import "fmt"

var visualizerLinks = map[string]string{
	"01_arrays_reflex":        "reference/reference/visualizer/index.html#reverse-string",
	"02_hashing_reflex":       "reference/reference/visualizer/index.html#two-sum",
	"03_two_pointers_reflex":  "reference/reference/visualizer/index.html#two-pointers",
	"04_binary_search_reflex": "reference/reference/visualizer/index.html#binary-search",
	"05_trees_stacks_reflex":  "reference/reference/visualizer/index.html#valid-parentheses",
	"06_dp_reflex":            "reference/reference/visualizer/index.html#climbing-stairs",
	"07_graphs_reflex":        "reference/reference/visualizer/index.html#num-islands",
	"08_heap_reflex":          "reference/reference/visualizer/index.html#kth-largest",
	"09_backtrack_reflex":     "reference/reference/visualizer/index.html#subsets",
	"core5":                   "reference/reference/visualizer/index.html#two-sum",
}

func printVisualizerLink(drillFile string) {
	link, ok := visualizerLinks[drillFile]
	if !ok {
		return
	}
	fmt.Printf("\n── VISUALIZER ─────────────────────────────────────────\n  %s\n", link)
}
