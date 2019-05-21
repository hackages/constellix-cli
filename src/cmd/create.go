package cmd

import (
	"github.com/spf13/cobra"
	F "../functions"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a dns records of types CNAME or A",
	Run: func(cmd *cobra.Command, args []string) {},
}

var (
	createCnameCmd = &cobra.Command{
        	Use:   "cname",
		Short: "Create a dns record of type CNAME",
		Args: cobra.MinimumNArgs(1),
		Run: createCname,
	}

	createCNAMECmd = &cobra.Command{
		Use: "CNAME",
		Short: "Create a dns record of type CNAME",
		Hidden: true,
		Args: cobra.MinimumNArgs(1),
		Run: createCname,
	}

	createaCmd = &cobra.Command{
        	Use:   "A",
        	Short: "Create a dns record of type A",
		Hidden: true,
		Args: cobra.MinimumNArgs(1),
		Run: createA,
	}

	createACmd = &cobra.Command{
                Use:   "a",
                Short: "Create a dns record of type A",
		Args: cobra.MinimumNArgs(1),
                Run: createA,
        }
)

func createCname(cmd *cobra.Command, args []string) {
	file, _:= cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)
        
	if len(args) == 1 {
        	F.CreateCNAME(args[0],c)
        } else if len(args) > 1 {
        	for i, _ := range args {
                	F.CreateCNAME(args[i],c)
        	}
	}
}

func createA(cmd *cobra.Command, args []string) {
        file, _:= cmd.Flags().GetString("command-config")
	name, _:= cmd.Flags().GetString("name")

	c := F.InitConfig(file)
	F.CreateA(args[0],name,c)
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createCnameCmd)
	createCmd.AddCommand(createaCmd)

	// Aliases
	createCmd.AddCommand(createCNAMECmd)
	createCmd.AddCommand(createACmd)

	// Local flags
	createACmd.Flags().StringP("name", "n", "", "Add a name to your A record")
	createaCmd.Flags().StringP("name", "n", "", "Add a name to your A record")
}
