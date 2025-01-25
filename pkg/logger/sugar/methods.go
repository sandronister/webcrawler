package sugar

func (s *Model) Debug(template string, args ...interface{}) {
	s.sugar.Debugf(template, args)
}

func (s *Model) Error(template string, args ...interface{}) {
	s.sugar.Errorf(template, args)
}

func (s *Model) Fatal(template string, args ...interface{}) {
	s.sugar.Fatalf(template, args)
}

func (s *Model) Info(template string, args ...interface{}) {
	s.sugar.Infof(template, args)
}

func (s *Model) Warn(template string, args ...interface{}) {
	s.sugar.Warnf(template, args)
}
