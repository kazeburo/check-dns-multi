package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	dnsmulti "github.com/kazeburo/check-dns-multi"
)

func main() {
	opt := &dnsmulti.Opt{}
	psr := flags.NewParser(opt, flags.HelpFlag|flags.PassDoubleDash)
	_, err := psr.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	ckr := opt.Resolve()
	ckr.Name = "DNS"
	ckr.Exit()
}
