# ChatShell

ChatShell is a CLI tool that helps you find shell commands for specific tasks. It uses OpenAI's GPT-3.5 language model to generate commands based on user input.

## Installation

To use ChatShell, you will need to have Golang installed on your machine. Once you have installed Golang, you can install ChatShell using the following command:

go get github.com/yourusername/chatshell

## Configuration

ChatShell reads your OpenAI API key from a JSON configuration file named config.json. By default, it looks for this file in the $HOME/.chatshell/ directory.

Here's an example of what your configuration file should look like:

``` json 
{
  "OPENAI_AUTH_TOKEN": "your_api_key_here"
}
```

## Usage

To use ChatShell, simply run the ask command followed by your query. For example:

`chatshell ask "How do I list all files in a directory?"`


This will return a shell command that you can use to list all files in a directory.

You can also enter chat mode by using the -c flag:


`chatshell ask -c "Hello, how are you?"`


This will enter chat mode and allow you to have a normal conversation with ChatGPT.

## Releases

Releases of ChatShell will be generated periodically and added to the repository. These releases will be available for download in the Releases section of this repository. To use a release, simply download the binary for your operating system and place it in a directory that is in your system's PATH.

## Contributing

If you encounter any bugs or issues with ChatShell, please feel free to create an issue in the repository. Pull requests are also welcome if you would like to contribute to the development of ChatShell.