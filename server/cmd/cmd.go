package cmd

func StartClient() {

}

func SignUp(args string) {
	// parse args into a message type
	msg := parseArgs(args)
	return handler.NewUser(msg)
}
