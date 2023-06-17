package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

type processor struct { // 全体の処理の構成
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

const maxInt = 20
const timeLimit = 50

func gatherAndProcess(ctx context.Context, data Input) (Cout, error) {
	log.SetFlags(log.Lmicroseconds)
	log.Println("starting: timelimit:", timeLimit*time.Millisecond)
	ctx, cancel := context.WithTimeout(ctx, timeLimit*time.Millisecond)
	defer cancel()
	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan Cin, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2),
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}
