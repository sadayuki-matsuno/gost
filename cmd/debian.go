package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/knqyf263/gost/db"
	"github.com/knqyf263/gost/fetcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debianCmd represents the debian command
var debianCmd = &cobra.Command{
	Use:   "debian",
	Short: "Fetch the CVE information from Debian",
	Long:  `Fetch the CVE information from Debian`,
	RunE:  fetchDebian,
}

func init() {
	fetchCmd.AddCommand(debianCmd)
}

func fetchDebian(cmd *cobra.Command, args []string) (err error) {
	log15.Info("Fetched all CVEs from Debian")
	cves, err := fetcher.RetrieveDebianCveDetails()

	log15.Info("Initialize Database")
	driver, err := db.InitDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		return err
	}

	log15.Info("Insert Debian CVEs into DB", "db", driver.Name())
	if err := driver.InsertDebian(cves); err != nil {
		log15.Error("Failed to insert.", "dbpath",
			viper.GetString("dbpath"), "err", err)
		return err
	}

	return nil
}
