package cmd

import (
	F "../functions"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a dns record by its name",
}

var (
	deleteCnameCmd = &cobra.Command{
		Use:   "cname",
		Short: "Delete a dns record of type CNAME",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteCname,
	}

	deleteCNAMECmd = &cobra.Command{
		Use:    "CNAME",
		Short:  "Delete a dns record of type CNAME",
		Run:    deleteCname,
		Args:   cobra.MinimumNArgs(1),
		Hidden: true,
	}

	deleteaCmd = &cobra.Command{
		Use:   "a",
		Short: "Delete a dns record of type A",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteA,
	}

	deleteACmd = &cobra.Command{
		Use:    "A",
		Short:  "Delete a dns record of type A",
		Run:    deleteA,
		Args:   cobra.MinimumNArgs(1),
		Hidden: true,
	}

	deletensCmd = &cobra.Command{
		Use:   "ns",
		Short: "Delete a dns record of type NS",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteNS,
	}

	deleteNSCmd = &cobra.Command{
		Use:    "NS",
		Short:  "Delete a dns record of type NS",
		Run:    deleteNS,
		Args:   cobra.MinimumNArgs(1),
		Hidden: true,
	}
)

func deleteCname(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	if len(args) == 1 {
		F.DeleteCNAME(args[0], c)
	} else if len(args) > 1 {
		for i, _ := range args {
			F.DeleteCNAME(args[i], c)
		}
	}
}

func deleteA(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	if len(args) >= 1 {
		F.DeleteA(args[0], c)
	}
}

func deleteNS(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("command-config")
	c := F.InitConfig(file)

	if len(args) >= 1 {
		F.DeleteNS(args[0], c)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteaCmd)
	deleteCmd.AddCommand(deleteCnameCmd)
	deleteCmd.AddCommand(deletensCmd)

	// Aliases
	deleteCmd.AddCommand(deleteACmd)
	deleteCmd.AddCommand(deleteCNAMECmd)
	deleteCmd.AddCommand(deleteNSCmd)
}
