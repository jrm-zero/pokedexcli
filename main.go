package main

func main() {
	myconfig := config{
		commands: getcommands(),
	}
	myConfigPtr := &myconfig
	startRepl(myConfigPtr)
}