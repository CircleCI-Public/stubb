package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "stubb",
	Short: "A Docker utility tool",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/stubb.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		cfgDir, err := os.UserConfigDir()
		if err != nil {
			cfgDir = "~/.config"
		}

		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("stubb")
	}

	viper.SetEnvPrefix("DOCKER")
	viper.AutomaticEnv() // read in envars that match

	// set config defaults
	viper.SetDefault("version", "0.1")
	viper.SetDefault("defaultRepository", "hub")

	viper.ReadInConfig()
}
