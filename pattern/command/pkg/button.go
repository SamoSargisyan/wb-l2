package pkg

type Button struct {
	Command command
}

func (b *Button) Press() {
	b.Command.execute()
}
