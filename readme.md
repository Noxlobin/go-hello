# Go Hello

Run these commands right now in your terminal (linux/debian only):

```bash
sudo apt install hello
```

```bash
hello
```

I'll spare the results for you. Today, I woke up and remade this app which just prints a greeting onto the screen. If you have a linux (or any os that has go installed) server that hasn't been used it a long time, this program might be useful to see that everything works well.

I dont know what this program solves, especially cause there is a GNU version out there. This uses go, a fairly good language. It uses cobra for the cli part. Its _very_ idiotproof. Honestly.

## Usage

I already told you: Just prints a altered version of hello world. Lets see some of the commands packaged with it.

- Traditional: A version without punctuations.
- Version: Prints the version.

There was one greeting command in the og version. It only supports one word, and can be filtered out by echo. In other words, its stupid.

## Installation

I'll probably make a flatpak or appimage of it someday. For now, clone the repo, build it using `go build`, add it to your path, and run `go-hello` or make a alias of it.

For choosing where to put the app, I reccomend `usr/bin`.
