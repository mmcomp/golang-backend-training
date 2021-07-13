package arg

import "flag"

var (
	Values []string
)

var (
	Query string
)

var (
	Verbose                     bool
	VeryVerbose                 bool
	VeryVeryVerbose             bool
	VeryVeryVeryVerbose         bool
	VeryVeryVeryVeryVerbose     bool
	VeryVeryVeryVeryVeryVerbose bool
)

func init() {
	flag.StringVar(&Query, "q", "", "search query")

	flag.BoolVar(&Verbose, "v", false, "verbose logs outputted")
	flag.BoolVar(&VeryVerbose, "vv", false, "very verbose logs outputted")
	flag.BoolVar(&VeryVeryVerbose, "vvv", false, "very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVerbose, "vvvv", false, "very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVerbose, "vvvvv", false, "very very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVeryVerbose, "vvvvvv", false, "very very very very very verbose logs outputted")

	flag.Parse()

	Values = flag.Args()
}