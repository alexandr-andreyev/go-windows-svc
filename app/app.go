package app

import "time"

// The wrapper of your app
func proskladSenderApp(s server) {

	s.winlog.Info(1, "In app.yourApp") //TODO исправить логи

	// This is just some sample code to do something
	time.Sleep(1 * time.Second)
	s.winlog.Info(1, "Still running")

	time.Sleep(2 * time.Second)
	s.winlog.Info(1, "And running")

	time.Sleep(3 * time.Second)
	s.winlog.Info(1, "But the service will keep running")

	// Notice that if this exits, the service continues to run
	// You can launch web servers, etc.
}
