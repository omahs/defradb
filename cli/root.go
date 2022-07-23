// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sourcenetwork/defradb/config"
	"github.com/sourcenetwork/defradb/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootDirParam string

var rootCmd = &cobra.Command{
	Use:   "defradb",
	Short: "DefraDB Edge Database",
	Long: `DefraDB is the edge database to power the user-centric future.

Start a database node, query a local or remote node, and much more.

DefraDB is released under the BSL license, (c) 2022 Democratized Data Foundation.
See https://docs.source.network/BSLv0.2.txt for more information.
`,
	// Runs on subcommands before their Run function, to handle configuration and top-level flags.
	// Loads the rootDir containing the configuration file, otherwise warn about it and load a default configuration.
	// This allows some subcommands (`init`, `start`) to override the PreRun to create a rootDir by default.
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		rootDir, exists, err := config.GetRootDir(rootDirParam)
		if err != nil {
			return fmt.Errorf("failed to get root dir: %w", err)
		}
		defaultConfig := false
		if exists {
			err := cfg.Load(rootDir)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}
<<<<<<< HEAD
			loggingConfig, err := cfg.GetLoggingConfig()
			if err != nil {
				return fmt.Errorf("failed to get logging config: %w", err)
			}
			logging.SetConfig(loggingConfig)
			log.Debug(cmd.Context(), fmt.Sprintf("Configuration loaded from DefraDB directory %v", rootDir))
||||||| parent of 5538451 (Support named logger config overrides)
			loggingConfig, err := cfg.GetLoggingConfig()
			if err != nil {
				log.FatalE(ctx, "Could not get logging config", err)
			}
			logging.SetConfig(loggingConfig)
			log.Debug(ctx, fmt.Sprintf("Configuration loaded from DefraDB directory %v", rootDir))
=======
>>>>>>> 5538451 (Support named logger config overrides)
		} else {
			err := cfg.LoadWithoutRootDir()
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}
<<<<<<< HEAD
			loggingConfig, err := cfg.GetLoggingConfig()
			if err != nil {
				return fmt.Errorf("failed to get logging config: %w", err)
			}
			logging.SetConfig(loggingConfig)
||||||| parent of 5538451 (Support named logger config overrides)
			loggingConfig, err := cfg.GetLoggingConfig()
			if err != nil {
				log.FatalE(ctx, "Could not get logging config", err)
			}
			logging.SetConfig(loggingConfig)
			log.Info(ctx, "Using default configuration. To create DefraDB's directory, use defradb init.")
=======
			defaultConfig = true
		}

		// parse loglevel overrides
		// we use `cfg.Logging.Level` as an argument since the viper.Bind already handles
		// binding the flags / EnvVars to the struct
		parseAndConfigLog(ctx, cfg.Logging, cmd)

		if defaultConfig {
			log.Info(ctx, "Using default configuration")
		} else {
			log.Debug(ctx, fmt.Sprintf("Configuration loaded from DefraDB directory %v", rootDir))

>>>>>>> 5538451 (Support named logger config overrides)
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&rootDirParam, "rootdir", "",
		"Directory for data and configuration to use (default \"$HOME/.defradb\")",
	)

	rootCmd.PersistentFlags().String(
		"loglevel", cfg.Log.Level,
		"Log level to use. Options are debug, info, error, fatal",
	)
	err := viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("loglevel"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind log.loglevel", err)
	}

	rootCmd.PersistentFlags().String(
<<<<<<< HEAD
		"logoutput", cfg.Log.OutputPath,
		"Log output path",
||||||| parent of 5538451 (Support named logger config overrides)
		"logoutput", cfg.Logging.OutputPath,
		"log output path",
=======
		"logger", "",
		"named logger parameter override. usage: --logger <name>,level=<level>,output=<output>,etc...",
	)

	rootCmd.PersistentFlags().String(
		"logoutput", cfg.Logging.OutputPath,
		"log output path",
>>>>>>> 5538451 (Support named logger config overrides)
	)
	err = viper.BindPFlag("log.outputpath", rootCmd.PersistentFlags().Lookup("logoutput"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind log.outputpath", err)
	}

	rootCmd.PersistentFlags().String(
		"logformat", cfg.Log.Format,
		"Log format to use. Options are text, json",
	)
	err = viper.BindPFlag("log.format", rootCmd.PersistentFlags().Lookup("logformat"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind log.format", err)
	}

	rootCmd.PersistentFlags().Bool(
		"logtrace", cfg.Log.Stacktrace,
		"Include stacktrace in error and fatal logs",
	)
	err = viper.BindPFlag("log.stacktrace", rootCmd.PersistentFlags().Lookup("logtrace"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind log.stacktrace", err)
	}

	rootCmd.PersistentFlags().Bool(
		"logcolor", cfg.Log.Color,
		"Enable colored output",
	)
	err = viper.BindPFlag("log.color", rootCmd.PersistentFlags().Lookup("logcolor"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind log.color", err)
	}

	rootCmd.PersistentFlags().String(
		"url", cfg.API.Address,
		"URL of the target database's HTTP endpoint",
	)
	err = viper.BindPFlag("api.address", rootCmd.PersistentFlags().Lookup("url"))
	if err != nil {
		log.FatalE(context.Background(), "Could not bind api.address", err)
	}
}

// parses and then configures the given config.Config logging subconfig.
// we use log.Fatal instead of returning an error because we can't gurantee
// atomic updates, its either everything is properly set, or we Fatal()
func parseAndConfigLog(ctx context.Context, cfg *config.LoggingConfig, cmd *cobra.Command) {
	// apply logger configuration at the end
	// once everything has been processed.
	defer func() {
		loggingConfig, err := cfg.ToLoggerConfig()
		if err != nil {
			log.FatalE(ctx, "Could not get logging config", err)
		}
		logging.SetConfig(loggingConfig)
	}()

	// handle --loglevels <default>,<name>=<value>,...
	parseAndConfigLogStringParam(ctx, cfg, cfg.Level, func(l *config.LoggingConfig, v string) {
		l.Level = v
	})

	// handle --logger <name>,<field>=<value>,...
	loggerKVs, err := cmd.Flags().GetString("logger")
	if err != nil {
		log.FatalE(ctx, "can't get logger flag", err)
	}

	if loggerKVs != "" {
		parseAndConfigLogAllParams(ctx, cfg, loggerKVs)
	}
}

func parseAndConfigLogAllParams(ctx context.Context, cfg *config.LoggingConfig, kvs string) {
	if kvs == "" {
		return //nothing todo
	}

	// check if a CSV is provided
	parsed := strings.Split(kvs, ",")
	if len(parsed) <= 1 {
		log.Fatal(ctx, "invalid --logger format, must be a csv")
	}
	name := parsed[0]

	// verify KV format (<default>,<field>=<value>,...)
	// skip the first as that will be set above
	for _, kv := range parsed[1:] {
		parsedKV := strings.Split(kv, "=")
		if len(parsedKV) != 2 {
			log.Fatal(ctx, "level was not provided as <key>=<value> pair", logging.NewKV("pair", kv))
		}

		logcfg, err := cfg.GetOrCreateNamedLogger(name)
		if err != nil {
			log.FatalE(ctx, "could not get named logger config", err)
		}

		// handle field
		switch strings.ToLower(parsedKV[0]) {
		case "level": // string
			logcfg.Level = parsedKV[1]
		case "format": // string
			logcfg.Format = parsedKV[1]
		case "output": // string
			logcfg.OutputPath = parsedKV[1]
		case "stacktrace": // bool
			boolValue, err := strconv.ParseBool(parsedKV[1])
			if err != nil {
				log.FatalE(ctx, "couldn't parse kv bool", err)
			}
			logcfg.Stacktrace = boolValue
		case "color": // bool
			boolValue, err := strconv.ParseBool(parsedKV[1])
			if err != nil {
				log.FatalE(ctx, "couldn't parse kv bool", err)
			}
			logcfg.Color = boolValue
		}
	}
}

func parseAndConfigLogStringParam(ctx context.Context, cfg *config.LoggingConfig, kvs string, paramSetterFn logParamSetterStringFn) {
	if kvs == "" {
		return //nothing todo
	}

	// check if a CSV is provided
	// if its not a CSV, then just do the regular binding to the config
	parsed := strings.Split(kvs, ",")
	paramSetterFn(cfg, parsed[0])
	if len(parsed) == 1 {
		return //nothing more todo
	}

	// verify KV format (<default>,<name>=<value>,...)
	// skip the first as that will be set above
	for _, kv := range parsed[1:] {
		parsedKV := strings.Split(kv, "=")
		if len(parsedKV) != 2 {
			log.Fatal(ctx, "level was not provided as <key>=<value> pair", logging.NewKV("pair", kv))
		}

		logcfg, err := cfg.GetOrCreateNamedLogger(parsedKV[0])
		if err != nil {
			log.FatalE(ctx, "could not get named logger config", err)
		}

		paramSetterFn(&logcfg.LoggingConfig, parsedKV[1])
	}
}

//
// LEAVE FOR NOW - IMPLEMENTING SOON - PLEASE IGNORE FOR NOW
//
// func parseAndConfigLogBoolParam(ctx context.Context, cfg *config.LoggingConfig, kvs string, paramFn logParamSetterBoolFn) {
// 	if kvs == "" {
// 		return //nothing todo
// 	}

// 	// check if a CSV is provided
// 	// if its not a CSV, then just do the regular binding to the config
// 	parsed := strings.Split(kvs, ",")
// 	boolValue, err := strconv.ParseBool(parsed[0])
// 	if err != nil {
// 		log.FatalE(ctx, "couldn't parse kv bool", err)
// 	}
// 	paramFn(cfg, boolValue)
// 	if len(parsed) == 1 {
// 		return //nothing more todo
// 	}

// 	// verify KV format (<default>,<name>=<level>,...)
// 	// skip the first as that will be set above
// 	for _, kv := range parsed[1:] {
// 		parsedKV := strings.Split(kv, "=")
// 		if len(parsedKV) != 2 {
// 			log.Fatal(ctx, "field was not provided as <key>=<value> pair", logging.NewKV("pair", kv))
// 		}

// 		logcfg, err := cfg.GetOrCreateNamedLogger(parsedKV[0])
// 		if err != nil {
// 			log.FatalE(ctx, "could not get named logger config", err)
// 		}

// 		boolValue, err := strconv.ParseBool(parsedKV[1])
// 		if err != nil {
// 			log.FatalE(ctx, "couldn't parse kv bool", err)
// 		}
// 		paramFn(&logcfg.LoggingConfig, boolValue)
// 	}
// }

type logParamSetterStringFn func(*config.LoggingConfig, string)

// type logParamSetterBoolFn func(*config.LoggingConfig, bool)
