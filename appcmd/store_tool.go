package appcmd

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/inhuman"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var storeTool = &cobra.Command{
	Use:   "store",
	Short: "Tool for tracking app store apps",
	Long:  "Tool for tracking app store apps",
	Run: func(cmd *cobra.Command, args []string) {
		c := loadConfig(config.AppStore)
		ex := loadTracker(
			c,
			api.NewRequester(inhuman.NewApiStore(inhuman.FromConfig(c))),
		)

		if err := ex.Work(); err != nil {
			log.Fatal(err)
		}

		log.Print("Off")
	},
}

func init() {
	rootCmd.AddCommand(storeTool)
	storeTool.PersistentFlags().StringVarP(&bundle, "target", "t", "", "Target bundle for tracking")
	storeTool.MarkFlagRequired("target")
	storeTool.PersistentFlags().StringVarP(&locale, "locale", "l", "ru_RU", "Locale for tracking")
	storeTool.PersistentFlags().IntVarP(&period, "period", "p", 30, "Period of tracking")
	storeTool.PersistentFlags().DurationVarP(&intensity, "intensity", "i", time.Hour*24, "tracking frequency")
	storeTool.PersistentFlags().IntVarP(&itemsCount, "count", "c", 200, "set count of apps for tracking")
	storeTool.PersistentFlags().StringVarP(&keywords, "keywords", "k", "", "keywords for tracking separated by commas")
	storeTool.PersistentFlags().StringVarP(&keyFile, "file", "f", "", "file with keywords separated by '\n'")
	storeTool.PersistentFlags().BoolVar(&force, "force", false, "force a new tracking instance")
}
