package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// --fn=NAME[,NAME...] restores just those functions to their blank stubs,
// leaving everything else in the drill file as you wrote it. That is what lets
// a session mix single functions from different drill files.

// parseFnList splits and checks a --fn value against the catalog.
func parseFnList(list string) ([]string, error) {
	var fns []string
	seen := map[string]bool{}
	for _, fn := range strings.Split(list, ",") {
		fn = strings.TrimSpace(fn)
		if fn == "" || seen[fn] {
			continue
		}
		if _, ok := drillFileOf(fn); !ok {
			return nil, fmt.Errorf("unknown function %q%s (see: go run . -- --catalog)", fn, suggestFn(fn))
		}
		seen[fn] = true
		fns = append(fns, fn)
	}
	if len(fns) == 0 {
		return nil, fmt.Errorf("option '--fn' needs at least one function name")
	}
	return fns, nil
}

// suggestFn offers catalog names that contain fn, ignoring case.
func suggestFn(fn string) string {
	var hits []string
	low := strings.ToLower(fn)
	for _, d := range drills {
		for _, f := range d.functions {
			if strings.Contains(strings.ToLower(f), low) {
				hits = append(hits, f)
			}
		}
	}
	if len(hits) == 0 {
		return ""
	}
	return "; did you mean " + strings.Join(hits, ", ") + "?"
}

// blankFunctions rewrites each function in fns back to its blank stub and
// stamps its start time in the log, so the next --run can time it.
func blankFunctions(repoRoot string, fns []string, now time.Time) error {
	for _, g := range groupByFile(fns) {
		path := filepath.Join(writeReflexDir(repoRoot, g.file), "main.go")
		if err := blankInFile(path, g.file, g.fns); err != nil {
			return err
		}
		fmt.Printf("blanked %s in drills/write/reflex/%s/main.go\n", strings.Join(g.fns, ", "), g.file)
	}
	stampStarted(repoRoot, fns, now)
	return nil
}

func blankInFile(path, file string, fns []string) error {
	current, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	blank, err := blankContent(file)
	if err != nil {
		return err
	}
	blankDecls, err := funcSources(blank, "blank "+file)
	if err != nil {
		return err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, current, 0)
	if err != nil {
		return fmt.Errorf("%s does not parse, so one function cannot be swapped out (fix it or use --reset): %w", path, err)
	}

	type span struct {
		start, end int
		repl       string
	}
	var spans []span
	for _, fn := range fns {
		repl, ok := blankDecls[fn]
		if !ok {
			return fmt.Errorf("no blank stub for %s in %s", fn, file)
		}
		decl := findFunc(f, fn)
		if decl == nil {
			return fmt.Errorf("%s: func %s not found", path, fn)
		}
		spans = append(spans, span{fset.Position(decl.Pos()).Offset, fset.Position(decl.End()).Offset, repl})
	}
	// Replace from the end so earlier offsets stay valid.
	sort.Slice(spans, func(i, j int) bool { return spans[i].start > spans[j].start })
	out := current
	for _, s := range spans {
		out = append(append(append([]byte{}, out[:s.start]...), s.repl...), out[s.end:]...)
	}
	if formatted, err := format.Source(out); err == nil {
		out = formatted
	}
	return os.WriteFile(path, out, 0o644)
}

// funcSources maps each top-level function name in src to its source text.
func funcSources(src []byte, name string) (map[string]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, name, src, 0)
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok && fd.Recv == nil {
			out[fd.Name.Name] = string(src[fset.Position(fd.Pos()).Offset:fset.Position(fd.End()).Offset])
		}
	}
	return out, nil
}

func findFunc(f *ast.File, name string) *ast.FuncDecl {
	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok && fd.Recv == nil && fd.Name.Name == name {
			return fd
		}
	}
	return nil
}

// stampStarted records now as the blank time of fns, so the next graded run
// can report how long writing them took.
func stampStarted(repoRoot string, fns []string, now time.Time) {
	log := loadLog(repoRoot)
	for _, fn := range fns {
		rec := log.Functions[fn]
		rec.Started = now.Format(time.RFC3339)
		log.Functions[fn] = rec
	}
	_ = saveLog(repoRoot, log)
}
