/*
Copyright © 2023 Mert Yiğittop <yigittopm@hotmail.com>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var (
	all        bool
	numbers    bool
	characters bool
	special    bool
	lenght     int
)

var (
	numbersList    = "0123456789"
	charactersList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialList    = "!@#$%^&*()-_+=<>?/[]{}|"
	charset        = numbersList + charactersList + specialList
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate password",
	Long:  "Generate password from simple to difficult",
	Run:   generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	if numbers || characters || special {
		all = false
		charset = ""
	}

	if numbers {
		charset += numbersList
	}
	if characters {
		charset += charactersList
	}
	if special {
		charset += specialList
	}

	password := make([]byte, lenght)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println("Generating password:")
	fmt.Println(string(password))
}

func init() {

	passwordCmd.Flags().BoolVarP(&all, "all", "a", true, "Password contain all chars")
	passwordCmd.Flags().BoolVarP(&numbers, "numbers", "n", false, "Password contain numbers")
	passwordCmd.Flags().BoolVarP(&characters, "characters", "c", false, "Password contain characters")
	passwordCmd.Flags().BoolVarP(&special, "special", "s", false, "Password contain special chars")
	passwordCmd.Flags().IntVarP(&lenght, "lenght", "l", 8, "Password length")

	rootCmd.AddCommand(passwordCmd)

}
