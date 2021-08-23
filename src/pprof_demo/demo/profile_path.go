package main

import "go_demo/src/profile"

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
}
