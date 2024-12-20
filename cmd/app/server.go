package app

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/darianJmy/fortis-services/api/server/router"
	"github.com/darianJmy/fortis-services/cmd/app/options"
)

func NewFortisServerCommand() *cobra.Command {
	o := options.NewServerRunOptions()

	cmd := &cobra.Command{
		Use: "fortis-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}

			if err := o.Registry(); err != nil {
				return err
			}

			return Run(o)
		},
	}

	o.BindFlags(cmd)

	return cmd
}

func Run(o *options.ServerRunOptions) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", o.Config.Default.Listen),
		Handler: o.HttpEngine,
	}

	router.InstallRouters(o)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
