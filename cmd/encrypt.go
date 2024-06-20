package cmd

import (
	"fmt"
	"log"
	"strings"

	"rayhanadev/gpg-tool/crypto"
	"rayhanadev/gpg-tool/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
    Use:   "encrypt [message]",
    Short: "Encrypt a message",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        verbose, _ := cmd.Flags().GetBool("verbose")
        utils.PrintVerbose(verbose, "Encrypt command invoked...")

        message := strings.Join(args, " ")

        keyID, _ := cmd.Flags().GetString("key-id")

        if keyID == "" {
            keys, err := crypto.ListGPGKeys()
            if err != nil {
                log.Fatalf("Failed to list GPG keys: %v", err)
            }

            var selectedKey string
            keyMap := make(map[string]string)
            options := make([]string, len(keys))
            for i, key := range keys {
                option := fmt.Sprintf("%s (%s)", key.UserID, key.KeyID)
                options[i] = option
                keyMap[option] = key.KeyID
            }

            prompt := &survey.Select{
                Message: "Choose a GPG key:",
                Options: options,
            }
            err = survey.AskOne(prompt, &selectedKey)
            if err != nil {
                log.Fatalf("Failed to select GPG key: %v", err)
            }

            keyID = keyMap[selectedKey]
        }

        publicKey, err := crypto.ExportPublicKey(keyID)
        if err != nil {
            log.Fatalf("Failed to export public key: %v", err)
        }

        utils.PrintVerbose(verbose, "Encrypting message...")
        encryptedMessage, err := crypto.EncryptMessage(publicKey, message)
        if err != nil {
            log.Fatalf("Error encrypting message: %v", err)
        }

        utils.PrintVerbose(verbose, "Encrypted Message:")
        fmt.Println(encryptedMessage)
    },
}

func init() {
    encryptCmd.Flags().String("key-id", "", "GPG key ID to use for encryption")
    encryptCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
    rootCmd.AddCommand(encryptCmd)
}
