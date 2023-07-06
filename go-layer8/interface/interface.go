package serverInterface

type ServerImpl interface {
	Serve() error
	Shutdown()
}
