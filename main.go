package main

// short introduction:
// Run this command right now in your terminal (linux/debian only): sudo apt install hello
// Now run 'hello' in ur terminal. What do you see? 'Hello, World!'
// Its somewhat of a safety maintainace program. And I'm here to recreate it.

import "github.com/Noxlobin/go-hello/cmd"

func main() {
	cmd.Execute()
}
