package service

type SI interface {
	Run() error
	Stop() error
	Call() int
}

func NewRPCService() *Ss {
	return &Ss{}
}

type Ss struct {
}

func (s *Ss) Run() error {
	return nil
}

func (s *Ss) Stop() error {
	return nil
}

func (s *Ss) Call() int {
	return 100
}
