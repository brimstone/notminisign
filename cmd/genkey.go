// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

// genkeyCmd represents the genkey command
var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		public, secret, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return err
		}

		publicEncoded := base58.CheckEncode(public, 0)
		publicFile, _ := cmd.Flags().GetString("public")
		if publicFile == "" {
			fmt.Println("PUBLIC=" + publicEncoded)
		} else {
			err := ioutil.WriteFile(publicFile, []byte(publicEncoded), 0777)
			if err != nil {
				return err
			}
		}
		secretEncoded := base58.CheckEncode(secret, 0)
		secretFile, _ := cmd.Flags().GetString("secret")
		if secretFile == "" {
			fmt.Println("SECRET=" + secretEncoded)
		} else {
			err := ioutil.WriteFile(secretFile, []byte(secretEncoded), 0777)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genkeyCmd)
	genkeyCmd.Flags().StringP("public", "p", "", "File to store public key")
	genkeyCmd.Flags().StringP("secret", "s", "", "File to store secret key")
}
