package appcmd

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var ErrNotConfigured = func(value interface{}) error {
	return fmt.Errorf("error app not configured %v", value)
}

var (
	PRODUCTION bool
	bundle     string
	locale     string
	period     int
	intensity  time.Duration
	email      string
	password   string
	proxy      string
	token      string
	gsfid      int
	device     string
	keywords   string
	keyFile    string
	itemsCount int
	force      bool
	onlyMeta   bool
)

// Root cobra command. With use pattern 'samurai'
var rootCmd = &cobra.Command{
	Use:   "samurai",
	Short: "Start the samurai tool",
	Long:  "/|\n                          /'||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||\n                         |  ||         __.--._\n                         |  ||      /~~   __.-~\\ _\n                         |  ||  _.-~ / _---._ ~-\\/~\\\n                         |  || // /  /~/  .-  \\  /~-\\\n                         |  ||((( /(/_(.-(-~~~~~-)_/ |\n                         |  || ) (( |_.----~~~~~-._\\ /\n                         |  ||    ) |              \\_|\n                         |  ||     (| =-_   _.-=-  |~)        ,\n                         |  ||      | `~~ |   ~~'  |/~-._-'/'/_,\n                         |  ||       \\    |        /~-.__---~ , ,\n                         |  ||       |   ~-''     || `\\_~~~----~\n                         |  ||_.ssSS$$\\ -====-   / )\\_  ~~--~\n                 ___.----|~~~|%$$$$$$/ \\_    _.-~ /' )$s._\n        __---~-~~        |   |%%$$$$/ /  ~~~~   /'  /$$$$$$$s__\n      /~       ~\\    ============$$/ /        /'  /$$$$$$$$$$$SS-.\n    /'      ./\\\\\\\\\\\\_( ~---._(_))$/ /       /'  /$$$$%$$$$$~      \\\n    (      //////////(~-(..___)/$/ /      /'  /$$%$$%$$$$'         \\\n     \\    |||||||||||(~-(..___)$/ /  /  /'  /$$$%$$$%$$$            |\n      `-__ \\\\\\\\\\\\\\\\\\\\\\(-.(_____) /  / /'  /$$$$%$$$$$%$             |\n          ~~\"\"\"\"\"\"\"\"\"\"-\\.(____) /   /'  /$$$$$%%$$$$$$\\_            /\n                        $|===|||  /'  /$$$$$$$%%%$$$$$( ~         ,'|\n                    __  $|===|%\\/'  /$$$$$$$$$$$%%%%$$|        ,''  |\n                   ///\\ $|===|/'  /$$$$$$%$$$$$$$%%%%$(            /'\n                    \\///\\|===|  /$$$$$$$$$%%$$$$$$%%%%$\\_-._       |\n                     `\\//|===| /$$$$$$$$$$$%%%$$$$$$-~~~    ~      /\n                       `\\|-~~(~~-`$$$$$$$$$%%%///////._       ._  |\n                       (__--~(     ~\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\        \\ \\\n                       (__--~~(       \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\|        \\/\n                        (__--~(       ||||||||||||||||||/       _/\n                         (__.--._____//////////////////__..---~~\n                         |   \"\"\"\"'''''           ___,,,,ss$$$%\n                        ,%\\__      __,,,\\sssSS$$$$$$$$$$$$$$%%\n                      ,%%%%$$$$$$$$$$\\;;;;\\$$$$$$$$$$$$$$$$%%%$.\n                     ,%%%%%%$$$$$$$$$$%\\;;;;\\$$$$$$$$$$$$%%%$$$$\n                   ,%%%%%%%%$$$$$$$$$%$$$\\;;;;\\$$$$$$$$$%%$$$$$$,\n                  ,%%%%%%%%%$$$$$$$$%$$$$$$\\;;;;\\$$$$$$%%$$$$$$$$\n                 ,%%%%%%%%%%%$$$$$$%$$$$$$$$$\\;;;;\\$$$%$$$$$$$$$$$\n                 %%%%%%%%%%%%$$$$$$$$$$$$$$$$$$\\;;;$$$$$$$$$$$$$$$\n                   \"\"==%%%%%%%$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$SV\"\n                               $$$$$$$$$$$$$$$$$$$$====\"\"\"\"\n                                 \"\"\"\"\"\"\"\"\"~~~~\n",
}

// Method for starting cobra cli.
// Method also check production build or not
// Production flag must contains in sysenvs
func Execute() {
	production := os.Getenv("production")
	if production == "true" {
		PRODUCTION = true
	}
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
