package cmd

import (
	F "../functions"
	"github.com/spf13/cobra"
	"fmt"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get one or all dns records/domains",
	Long: `constellix-cli get [all] return all records in a domain `,
	Run: func(cmd *cobra.Command, args []string) {},
}

var ( 
	getDomainCmd = &cobra.Command{
        	Use:   "domain",
	        Short: "Get all domains",
		Run: getDomain,
	}

	getDOMAINCmd = &cobra.Command{
                Use:   "DOMAIN",
		Hidden: true,
                Short: "Get all domains",
        	Run: getDomain,
	}

	getAllCmd = &cobra.Command{
        	Use:   "all",
        	Short: "Get all records in a domain",
		Run: getAll,
	}
	
	getALLCmd = &cobra.Command{
                Use:   "ALL",
		Hidden: true,
                Short: "Get all records in a domain",
                Run: getAll,
        }

	getCnameIdCmd = &cobra.Command{
                Use:   "cname-id",
                Short: "Get a CNAME record id from a domain",
                Run: getId,
        }

	getCNAMEIdCmd = &cobra.Command{
                Use:   "CNAME-ID",
                Short: "Get a CNAME record id from a domain",
                Hidden: true,
                Run: getId,
        }
	
	getaIdCmd = &cobra.Command{
                Use:   "a-id",
                Short: "Get an A record id from a domain",
                Run: getId,
        }

	getAIdCmd = &cobra.Command{
                Use:   "A-ID",
                Short: "Get an A record id from a domain",
                Hidden: true,
		Run: getId,
        }
)

func getDomain(cmd *cobra.Command, args []string) {
	file, _:= cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	F.GetAllDomain(c)
}

func getAll(cmd *cobra.Command, args []string) {
	file, _:= cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

        F.GetAll(c)
}

func getId(cmd *cobra.Command, args []string) {
	file, _:= cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

        nameId := F.GetId(args[0],c)
	fmt.Println(nameId)
}


func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getDomainCmd)
	getCmd.AddCommand(getAllCmd)
	getCmd.AddCommand(getCnameIdCmd)
	getCmd.AddCommand(getaIdCmd)
	
	// Aliases
	getCmd.AddCommand(getALLCmd)
	getCmd.AddCommand(getDOMAINCmd)
	getCmd.AddCommand(getCNAMEIdCmd)
	getCmd.AddCommand(getAIdCmd)
}
