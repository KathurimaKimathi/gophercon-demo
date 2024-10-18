package main

import (
	"context"
	"strconv"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/presentation"
)

func main() {
	ctx := context.Background()

	port, err := strconv.Atoi(helpers.MustGetEnvVar("PORT"))
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	presentation.PrepareAndStartServer(ctx, port)
}
