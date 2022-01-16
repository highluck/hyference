package cmd

import (
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/container"
	"github.com/hyference/internal/filesystem"
	"github.com/hyference/internal/server"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Inference struct {
	Port                   int                     `json:"port"`
	DebugLabel             string                  `json:"debug_label"`
	MlLibType              string                  `json:"ml_lib_type"`
	ModelName              string                  `json:"model_name"`
	ModelPath              string                  `json:"model_path"`
	FileSystemClientType   string                  `json:"file_system_type"`
	FileSystemClientDetail filesystem.ClientDetail `json:"file_system_client_detail"`
	ParameterType          string                  `json:"parameter_type"`
	Uri                    string                  `json:"uri"`
}

func Start() error {
	c := &cobra.Command{
		Use:   "ss",
		Short: "search-flow application",
		Run: func(c *cobra.Command, _ []string) {
			_ = c.Help()
		},
	}

	var cfgJson string
	var fileType string
	var cfgPath string

	options := func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&cfgJson, "config", "c", "", "config (s)")
		cmd.Flags().StringVarP(&fileType, "fileType", "f", "", "fileType (s)")
		cmd.Flags().StringVarP(&cfgPath, "configPath", "p", "", "config path(s)")
	}

	c.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "start application",
		RunE: func(c *cobra.Command, _ []string) error {
			log.Info().Msgf("start command ")
			cfg := config.Config{}
			if cfgJson != "" {
				err := jsoniter.UnmarshalFromString(cfgJson, &cfg)
				if err != nil {
					log.Fatal().Err(err).Msgf("server load fail error config")
				}
			} else if cfgPath != "" {
				cfg = readingYml(cfgPath, fileType)
			}

			server.Server(container.New(cfg))
			return nil
		},
	})
	options(c)
	return nil
}

func readingYml(configPath string, fileType string) config.Config {
	if fileType == "" {
		fileType = "yaml"
	}

	cfg := config.Config{}
	viper.SetConfigType(fileType)
	viper.AddConfigPath(".")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msgf("failed to read config")
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msgf("failed to unmarshal config")
	}
	return cfg
}
