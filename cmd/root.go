package cmd

import (
	"amin/common/global"
	"errors"
	"fmt"
	"os"

	"amin/core/sdk/pkg"

	"github.com/spf13/cobra"

	"amin/cmd/api"
	"amin/cmd/config"
	"amin/cmd/migrate"
	"amin/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "amin",
	Short:        "amin",
	SilenceUsage: true,
	Long:         `amin是一个后台管理系统`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`amin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
