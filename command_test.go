package cluster

func ExampleRunSerial() {
	cmd := Command([]string{"localhost"}, "echo 'test'")
	cmd.RunSerial()
	// Output: test
}
