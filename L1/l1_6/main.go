package main

import (
	"flag"
	"l6/after"
	"l6/condition"
	"l6/ctx"
	"l6/goexit"
)

type Flags struct {
	Exit      bool
	Ctx       bool
	After     bool
	Condition bool
	Time      bool
}

func main() {
	var flags Flags
	flags.Parse()

	if flags.Ctx {
		if flags.Time {
			ctx.CtxWithTimeout(3)
		} else {
			ctx.CtxWithCancel()
		}
	}
	if flags.After {
		after.After()
	}

	if flags.Condition {
		condition.Condition()
	}

	if flags.Exit {
		goexit.GoExit()
	}
}

func (f *Flags) Parse() {
	flag.BoolVar(&f.Exit, "exit", false, "")
	flag.BoolVar(&f.Ctx, "ctx", false, "")
	flag.BoolVar(&f.After, "after", false, "")
	flag.BoolVar(&f.Condition, "cond", false, "")
	flag.BoolVar(&f.Time, "t", false, "")
	flag.Parse()
}
