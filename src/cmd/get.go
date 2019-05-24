package cmd

import (
	"fmt"

	F "../functions"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get one or all dns records/domains",
}

var (
	getDomainCmd = &cobra.Command{
		Use:   "domain",
		Short: "Get all domains",
		Run:   getDomain,
	}

	getDOMAINCmd = &cobra.Command{
		Use:    "DOMAIN",
		Hidden: true,
		Short:  "Get all domains",
		Run:    getDomain,
	}

	getAllCmd = &cobra.Command{
		Use:   "all",
		Short: "Get all records in a domain",
		Run:   getAll,
	}

	getALLCmd = &cobra.Command{
		Use:    "ALL",
		Hidden: true,
		Short:  "Get all records in a domain",
		Run:    getAll,
	}

	getCnameIdCmd = &cobra.Command{
		Use:   "cname-id",
		Short: "Get a CNAME record id from a domain",
		Args:  cobra.MinimumNArgs(1),
		Run:   getId,
	}

	getCNAMEIdCmd = &cobra.Command{
		Use:    "CNAME-ID",
		Short:  "Get a CNAME record id from a domain",
		Hidden: true,
		Args:   cobra.MinimumNArgs(1),
		Run:    getId,
	}

	getaIdCmd = &cobra.Command{
		Use:   "a-id",
		Args:  cobra.MinimumNArgs(1),
		Short: "Get an A record id from a domain",
		Run:   getId,
	}

	getAIdCmd = &cobra.Command{
		Use:    "A-ID",
		Args:   cobra.MinimumNArgs(1),
		Short:  "Get an A record id from a domain",
		Hidden: true,
		Run:    getId,
	}

	getnsIdCmd = &cobra.Command{
		Use:   "ns-id",
		Args:  cobra.MinimumNArgs(1),
		Short: "Get an NS record id from a domain",
		Run:   getId,
	}

	getNSIdCmd = &cobra.Command{
		Use:    "NS-ID",
		Args:   cobra.MinimumNArgs(1),
		Short:  "Get an NS record id from a domain",
		Hidden: true,
		Run:    getId,
	}

	getDomainIdCmd = &cobra.Command{
		Use:   "domain-id",
		Args:  cobra.MinimumNArgs(1),
		Short: "Get a domain id from its name",
		Run:   getDomainId,
	}
)

func getDomain(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	domains := F.GetAllDomain(c)

	for _, i := range domains {
		fmt.Printf("%v\t %v \n", i.Id, i.Name)
	}
}

func getAll(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	F.GetAll(c)
}

func getId(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	nameId := F.GetId(args[0], c)
	fmt.Println(nameId)
}

func getDomainId(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	domainId := F.GetDomainId(args[0], c)
	fmt.Println(domainId)
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getDomainCmd)
	getCmd.AddCommand(getAllCmd)
	getCmd.AddCommand(getCnameIdCmd)
	getCmd.AddCommand(getaIdCmd)
	getCmd.AddCommand(getnsIdCmd)
	getCmd.AddCommand(getDomainIdCmd)

	// Aliases
	getCmd.AddCommand(getALLCmd)
	getCmd.AddCommand(getDOMAINCmd)
	getCmd.AddCommand(getCNAMEIdCmd)
	getCmd.AddCommand(getAIdCmd)
	getCmd.AddCommand(getNSIdCmd)
}
