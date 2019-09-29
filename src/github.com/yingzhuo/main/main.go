package main

func main() {

	if flags.version {
		printVersion()
	} else {
		startHttpServer()
	}
}
