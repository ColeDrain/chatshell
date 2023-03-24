/*
Copyright © 2023 Ugochukwu Onyebuchi <pyvinci@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	openai "github.com/sashabaranov/go-openai"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask [text]",
	Short: "ask ChatGPT",
	Long:  `ask is used to ask for shell commands for a particular task`,
	Args:  cobra.ExactArgs(1),
	Run:   runAsk,
}
var authToken string

func init() {
	rootCmd.AddCommand(askCmd)

	askCmd.PersistentFlags().BoolP("chat", "c", false, "Chat with ChatGPT")

	// Read the auth token from a configuration file
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.chatshell")
	viper.AddConfigPath(`%USERPROFILE%/.chatshell`)
	viper.SetEnvPrefix("OPENAI")
	viper.AutomaticEnv()


	if err := viper.ReadInConfig(); err == nil {
		authToken = viper.GetString("OPENAI_AUTH_TOKEN")
	} else {
		fmt.Printf("Error reading configuration file: %v\n", err)
	}

	if authToken == "" {
		fmt.Println("Error: OPENAI_AUTH_TOKEN environment variable not set")
		os.Exit(1)
	}
}

func runAsk(cmd *cobra.Command, args []string) {
	client := openai.NewClient(authToken)
	osType := runtime.GOOS
	shell := os.Getenv("SHELL")
	content := ""

	chatMode, _ := cmd.Flags().GetBool("chat")

	if !chatMode {
        content = fmt.Sprintf(`You are a very helpful shell assistant that gives users only shell commands to achieve a task, just give out only the shell command(s), and nothing else, no preamble, greetings or explanation please, just the shell command. When you can't find a command for a query/prompt/greeting respond strictly with "Sorry I can't seem to find a command for that". Start now: "%v in %v os using %v shell"`, args[0], osType, shell)
    } else {
        content = args[0]
    }

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
