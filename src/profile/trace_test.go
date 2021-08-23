package profile_test

import "go_demo/src/profile"

func ExampleTraceProfile() {
	// use execution tracing, rather than the default cpu profiling.
	defer profile.Start(profile.TraceProfile).Stop()
}