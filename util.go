package measuregc

type errorConsumer struct {
	err error
}

func (e *errorConsumer) consume(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *errorConsumer) panic() {
	if e.err != nil {
		panic(e.err)
	}
}

func (e *errorConsumer) clear() {
	e.err = nil
}
