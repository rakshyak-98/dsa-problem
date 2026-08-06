package main

import "fmt"

var visualizerLinks = map[string]string{
	"01_arrays_reflex":        "reference/visualizer/index.html#reverse-string",
	"02_hashing_reflex":       "reference/visualizer/index.html#two-sum",
	"03_two_pointers_reflex":  "reference/visualizer/index.html#two-pointers",
	"04_binary_search_reflex": "reference/visualizer/index.html#binary-search",
	"05_trees_stacks_reflex":  "reference/visualizer/index.html#valid-parentheses",
	"06_dp_reflex":            "reference/visualizer/index.html#climbing-stairs",
	"07_graphs_reflex":        "reference/visualizer/index.html#num-islands",
	"08_heap_reflex":          "reference/visualizer/index.html#kth-largest",
	"09_backtrack_reflex":     "reference/visualizer/index.html#subsets",
	"10_math_reflex":          "doc/write/MATH_CONCEPTS.md",
	"core5":                   "reference/visualizer/index.html#two-sum",
}

func printVisualizerLink(drillFile string) {
	link, ok := visualizerLinks[drillFile]
	if !ok {
		return
	}
	fmt.Printf("\n── VISUALIZER ─────────────────────────────────────────\n  %s\n", link)
}
