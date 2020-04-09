// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-load-test-ng/deployment"
	"github.com/mattermost/mattermost-load-test-ng/deployment/terraform"
	"github.com/mattermost/mattermost-load-test-ng/logger"

	"github.com/spf13/cobra"
)

func RunCreateCmdF(cmd *cobra.Command, args []string) error {
	config, err := getConfig(cmd)
	if err != nil {
		return err
	}

	t := terraform.New(config)
	defer t.Cleanup()
	return t.Create()
}

func RunDestroyCmdF(cmd *cobra.Command, args []string) error {
	config, err := getConfig(cmd)
	if err != nil {
		return err
	}

	t := terraform.New(config)
	defer t.Cleanup()
	return t.Destroy()
}

func RunStartCmdF(cmd *cobra.Command, args []string) error {
	config, err := getConfig(cmd)
	if err != nil {
		return err
	}

	t := terraform.New(config)
	defer t.Cleanup()
	return t.StartCoordinator()
}

func RunInfoCmdF(cmd *cobra.Command, args []string) error {
	config, err := getConfig(cmd)
	if err != nil {
		return err
	}

	t := terraform.New(config)
	return t.Info()
}

func RunStopCmdF(cmd *cobra.Command, args []string) error {
	config, err := getConfig(cmd)
	if err != nil {
		return err
	}

	t := terraform.New(config)
	defer t.Cleanup()
	return t.StopCoordinator()
}

func getConfig(cmd *cobra.Command) (*deployment.Config, error) {
	configFilePath, _ := cmd.Flags().GetString("config")
	cfg, err := deployment.ReadConfig(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if err := cfg.IsValid(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	logger.Init(&cfg.LogSettings)
	return cfg, nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:          "ltctl",
		SilenceUsage: true,
		Short:        "Manage and control load-test deployments",
	}
	rootCmd.PersistentFlags().StringP("config", "c", "", "path to the deployer configuration file to use")

	deploymentCmd := &cobra.Command{
		Use:   "deployment",
		Short: "Manage a load-test deployment",
	}

	deploymentCommands := []*cobra.Command{
		{
			Use:   "create",
			Short: "Create a new load-test deployment",
			RunE:  RunCreateCmdF,
		},
		{
			Use:   "destroy",
			Short: "Destroy the current load-test deployment",
			RunE:  RunDestroyCmdF,
		},
		{
			Use:   "info",
			Short: "Display information about the current load-test deployment",
			RunE:  RunInfoCmdF,
		},
	}

	deploymentCmd.AddCommand(deploymentCommands...)
	rootCmd.AddCommand(deploymentCmd)

	loadtestCmd := &cobra.Command{
		Use:   "loadtest",
		Short: "Manage the load-test",
	}

	loadtestComands := []*cobra.Command{
		{
			Use:   "start",
			Short: "Start the coordinator in the current load-test deployment",
			RunE:  RunStartCmdF,
		},
		{
			Use:   "stop",
			Short: "Stop the coordinator in the current load-test deployment",
			RunE:  RunStopCmdF,
		},
	}

	loadtestCmd.AddCommand(loadtestComands...)
	rootCmd.AddCommand(loadtestCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
