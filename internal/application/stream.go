package application

import (
	"context"
	"fmt"
	"kafka-final/internal/infra/msg"

	"github.com/manifoldco/promptui"
)

func StartStream(
	p *msg.Processor,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterCh := make(chan string)
	defer close(filterCh)

	go func() {
		p.Run(ctx, filterCh)
		defer p.Stop()
	}()

	prompt := promptui.Prompt{
		Label: "Filter product name contain",
	}

	for {
		p, _ := prompt.Run()
		fmt.Printf("Filtering substring %q\n", p)
		filterCh <- p
	}
}
