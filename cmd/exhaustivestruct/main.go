package main

import (
	"flag"

	"github.com/syumai/exhaustivestruct/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	flag.Bool("unsafeptr", false, "")
	singlechecker.Main(analyzer.Analyzer)
}
