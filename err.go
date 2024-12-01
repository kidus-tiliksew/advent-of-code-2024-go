package main

import "fmt"

type ErrorContext struct {
	Day  string
	Part string
	Msg  string
}

func handleError(err error, ctx ErrorContext) {
	if err != nil {
		panic(fmt.Errorf("day %s: part %s: %s: %w", ctx.Day, ctx.Part, ctx.Msg, err))
	}
}
