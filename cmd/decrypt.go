package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"rayhanadev/gpg-tool/crypto"
	"rayhanadev/gpg-tool/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
    Use:   "decrypt [encrypted message file]",
    Short: "Decrypt a message",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        verbose, _ := cmd.Flags().GetBool("verbose")
        utils.PrintVerbose(verbose, "Decrypt command invoked...")

        encMsgFile := args[0]

        utils.PrintVerbose(verbose, fmt.Sprintf("Reading encrypted message from: %s\n", encMsgFile))
        encryptedMessage, err := ioutil.ReadFile(encMsgFile)
        if err != nil {
            log.Fatalf("Error reading encrypted message: %v", err)
            return
        }

        privateKeys, err := crypto.ListGPGPrivateKeys()
        if err != nil {
            log.Fatalf("Failed to list GPG private keys: %v", err)
        }

        if len(privateKeys) == 0 {
            log.Fatalf("No private keys found")
            return
        }

        var selectedKey string
        keyMap := make(map[string]string)
        options := make([]string, len(privateKeys))
        for i, key := range privateKeys {
            option := fmt.Sprintf("%s (%s)", key.UserID, key.KeyID)
            options[i] = option
            keyMap[option] = key.KeyID
        }

        prompt := &survey.Select{
            Message: "Choose a GPG private key:",
            Options: options,
        }
        err = survey.AskOne(prompt, &selectedKey)
        if err != nil {
            log.Fatalf("Failed to select GPG private key: %v", err)
        }

        selectedKeyID := keyMap[selectedKey]
        utils.PrintVerbose(verbose, fmt.Sprintf("Using private key: %s (%s)\n", selectedKey, selectedKeyID))

        privateKey, err := crypto.ExportPrivateKey(selectedKeyID)
        if err != nil {
            log.Fatalf("Failed to load private key: %v", err)
            return
        }

				var passphrase string
        passphrasePrompt := &survey.Password{
            Message: "Enter the passphrase for the private key:",
        }
        err = survey.AskOne(passphrasePrompt, &passphrase)
        if err != nil {
            log.Fatalf("Failed to get passphrase: %v", err)
        }

        utils.PrintVerbose(verbose, "Decrypting message...")
        decryptedMessage, err := crypto.DecryptMessage(privateKey, passphrase, string(encryptedMessage))
        if err != nil {
            log.Fatalf("Error decrypting message: %v", err)
            return
        }

        utils.PrintVerbose(verbose, "Decrypted Message:")
        fmt.Println(decryptedMessage)
    },
}

func init() {
    decryptCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
    rootCmd.AddCommand(decryptCmd)
}
