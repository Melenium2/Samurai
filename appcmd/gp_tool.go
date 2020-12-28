package appcmd

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var gpTool = &cobra.Command{
	Use:   "google",
	Short: "Tool for tracking google play apps",
	Long:  "Tool for tracking google play apps",
	Run: func(cmd *cobra.Command, args []string) {
		c := loadConfig(config.GooglePlay)
		ex := loadTracker(
			c,
			api.New(mobilerpc.New(mobilerpc.FromConfig(c)), inhuman.NewApiPlay(inhuman.FromConfig(c))),
		)

		if err := ex.Work(); err != nil {
			log.Fatal(err)
		}

		log.Print("Off")
	},
}

func init() {
	rootCmd.AddCommand(gpTool)
	gpTool.PersistentFlags().StringVarP(&bundle, "target", "t", "", "Target bundle for tracking")
	gpTool.MarkFlagRequired("target")
	gpTool.PersistentFlags().StringVarP(&locale, "locale", "l", "ru_RU", "Locale for tracking")
	gpTool.PersistentFlags().IntVarP(&period, "period", "p", 30, "Period of tracking")
	gpTool.PersistentFlags().DurationVarP(&intensity, "intensity", "i", time.Hour*24, "tracking frequency")
	gpTool.PersistentFlags().IntVarP(&itemsCount, "count", "c", 250, "set count of apps for tracking")
	gpTool.Flags().StringVarP(&email, "email", "e", "", "email for the device user account")
	gpTool.MarkFlagRequired("email")
	gpTool.Flags().StringVar(&password, "password", "", "password for the device user account")
	gpTool.MarkFlagRequired("password")
	gpTool.Flags().StringVar(&token, "token", "", "token instead of user password (must be paired with gsfid)")
	gpTool.Flags().IntVar(&gsfid, "gsfig", 0, "gsfid instead of user email (must be paired with token)")
	gpTool.PersistentFlags().StringVarP(&device, "device", "d", "whyred", "name of user device")
	gpTool.PersistentFlags().StringVar(&proxy, "proxy", "", "proxy for external requests from the device")
	gpTool.PersistentFlags().StringVarP(&keywords, "keywords", "k", "", "keywords for tracking separated by commas")
	gpTool.PersistentFlags().StringVarP(&keyFile, "file", "f", "", "file with keywords separated by '\n'")
	gpTool.PersistentFlags().BoolVar(&force, "force", false, "force a new tracking instance")
}
