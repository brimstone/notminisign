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
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secret, _ := cmd.Flags().GetString("secret")
		if secret == "" {
			return errors.New("Must specify a secret")
		}
		secretActual, version, err := base58.CheckDecode(secret)
		if err != nil {
			secretContents, err := ioutil.ReadFile(secret)
			if err != nil {
				return err
			}
			secretActual, version, err = base58.CheckDecode(string(secretContents))
			if err != nil {
				return err
			}
		}

		filename, _ := cmd.Flags().GetString("input")
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		if version != 0 {
			return errors.New("Unsupported Version")
		}
		sig := ed25519.Sign(ed25519.PrivateKey(secretActual), file)
		sigEncoded := base58.CheckEncode(sig, 0)
		output, _ := cmd.Flags().GetString("output")
		if output == "" {
			fmt.Println("SIGNATURE=" + sigEncoded)
		} else {
			err := ioutil.WriteFile(output, []byte(sigEncoded), 0777)
			if err != nil {
				return err
			}

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringP("input", "i", "", "File to sign")
	signCmd.Flags().StringP("output", "o", "", "Filename to contain signature")
	signCmd.Flags().StringP("secret", "s", "", "Secret key")

}
