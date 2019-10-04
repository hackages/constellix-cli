package cmd

import (
	F "../functions"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a dns records of types CNAME, A or NS",
}

var (
	createCnameCmd = &cobra.Command{
		Use:   "cname",
		Short: "Create a dns record of type CNAME",
		Args:  cobra.MinimumNArgs(1),
		Run:   createCname,
	}

	createCNAMECmd = &cobra.Command{
		Use:    "CNAME",
		Short:  "Create a dns record of type CNAME",
		Hidden: true,
		Args:   cobra.MinimumNArgs(1),
		Run:    createCname,
	}

	createACmd = &cobra.Command{
		Use:    "A",
		Short:  "Create a dns record of type A",
		Hidden: true,
		Args:   cobra.MinimumNArgs(1),
		Run:    createA,
	}

	createaCmd = &cobra.Command{
		Use:   "a",
		Short: "Create a dns record of type A",
		Args:  cobra.MinimumNArgs(1),
		Run:   createA,
	}

	creatensCmd = &cobra.Command{
		Use:   "ns",
		Short: "Create a dns record of type NS",
		Args:  cobra.MinimumNArgs(1),
		Run:   createNS,
	}

	createNSCmd = &cobra.Command{
		Use:    "NS",
		Short:  "Create a dns record of type NS",
		Args:   cobra.MinimumNArgs(1),
		Hidden: true,
		Run:    createNS,
	}

	createDomainCmd = &cobra.Command{
		Use:   "domain",
		Short: "Create a domain with one or multiple names",
		Args:  cobra.MinimumNArgs(1),
		Run:   createDomain,
	}
)

func createCname(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	if len(args) == 1 {
		F.CreateCNAME(args[0], c)
	} else if len(args) > 1 {
		for i, _ := range args {
			F.CreateCNAME(args[i], c)
		}
	}
}

func createA(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	name, _ := cmd.Flags().GetString("name")

	c := F.InitConfig(file)
	F.CreateA(args[0], name, c)
}

func createNS(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	name, _ := cmd.Flags().GetString("name")

	c := F.InitConfig(file)
	F.CreateNS(args[0], name, c)
}

func createDomain(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")

	c := F.InitConfig(file)

	F.CreateDomain(args, c)
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createCnameCmd)
	createCmd.AddCommand(createaCmd)
	createCmd.AddCommand(creatensCmd)
	createCmd.AddCommand(createDomainCmd)

	// Aliases
	createCmd.AddCommand(createCNAMECmd)
	createCmd.AddCommand(createACmd)
	createCmd.AddCommand(createNSCmd)

	// Local flags
	createACmd.Flags().StringP("name", "n", "", "Add a name value to your record")
	createaCmd.Flags().StringP("name", "n", "", "Add a name value to your record")

	createNSCmd.Flags().StringP("name", "n", "", "Add a name value to your record")
	creatensCmd.Flags().StringP("name", "n", "", "Add a name value to your record")
}
