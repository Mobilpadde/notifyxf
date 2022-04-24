package notifyxf

type Option func(*Notifier)

func WithHandle(handle string) Option {
	return func(n *Notifier) {
		n.handle = handle
	}
}

func WithParseMode(mode ParseMode) Option {
	return func(n *Notifier) {
		if mode == ParseModeHTML {
			n.parseMode = "HTML"
		} else {
			n.parseMode = "MarkdownV2"
		}
	}
}

func WithRedact() Option {
	return func(n *Notifier) {
		n.redact = true
	}
}

func WithDisabledPreview() Option {
	return func(n *Notifier) {
		n.preview = false
	}
}
