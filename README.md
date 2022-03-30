# Oh-heck CLI

[Oh-heck](https://oh-heck.dev) is a CLI tool that takes natural language input and outputs a terminal command using GPT-3.

![](https://i.imgur.com/taLeNZ8.gif)

## Installation

To install go [https://oh-heck.dev](https://oh-heck.dev) and download the correct binary for your device. Once extracted move `oh-heck` to `/usr/local/bin`.

It's possible that you get a permission error when trying to execute the application. If this happens, run:

```bash
chmod +x oh-heck
```

## Usage

Using oh-heck is simple. All you have to do is call it and send your question as the only parameter (in quotes).

```bash
oh-heck "How do I install vim using brew?"
```

The AI will then return its best guess and you can accept this by typing in `y` or `n`. If you accept it, the command will be copied to your clipboard for you to paste into Terminal. If you reject it, you can edit your question to try and get a different output. Press `Ctrl + c` at any time to exit the application.

```bash
? Output: $ brew install vim? [y/N] █
```
