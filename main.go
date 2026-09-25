package main

func main() {
	myConfigPtr := &config{
		commands: getcommands(),
		locationsOffset: 0,
	}
	startRepl(myConfigPtr)
}