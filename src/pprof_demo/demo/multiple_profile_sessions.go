package main

import "go_demo/src/profile"

func main() {
	profile.Start(profile.CPUProfile).Stop()
	profile.Start(profile.MemProfile).Stop()
	profile.Start(profile.BlockProfile).Stop()
	profile.Start(profile.CPUProfile).Stop()
	profile.Start(profile.MutexProfile).Stop()
}
