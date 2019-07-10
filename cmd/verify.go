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

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		public, _ := cmd.Flags().GetString("public")
		if public == "" {
			return errors.New("Must specify a public key")
		}
		publicActual, version, err := base58.CheckDecode(public)
		if err != nil {
			publicContents, err := ioutil.ReadFile(public)
			if err != nil {
				return err
			}
			publicActual, version, err = base58.CheckDecode(string(publicContents))
			if err != nil {
				return err
			}
		}
		if version != 0 {
			return errors.New("Unsupported Version in public key")
		}

		sig, _ := cmd.Flags().GetString("sig")
		if sig == "" {
			return errors.New("Must specify a signature")
		}
		sigActual, version, err := base58.CheckDecode(sig)
		if err != nil {
			sigContents, err := ioutil.ReadFile(sig)
			if err != nil {
				return err
			}
			sigActual, version, err = base58.CheckDecode(string(sigContents))
			if err != nil {
				return err
			}
		}
		if version != 0 {
			return errors.New("Unsupported Version in signature")
		}

		filename, _ := cmd.Flags().GetString("input")
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		if !ed25519.Verify(ed25519.PublicKey(publicActual), file, sigActual) {
			return errors.New("Signature verification failed")
		}

		fmt.Println("Signature verification passed")
		return nil

	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringP("input", "i", "", "File to sign")
	verifyCmd.Flags().StringP("public", "p", "", "Public key")
	verifyCmd.Flags().StringP("sig", "s", "", "Signature")
}
