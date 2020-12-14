package l9format


func (event *L9Event) RemoveTransport(transportCheck string) {
	transports := event.Transports
	event.Transports = []string{}
	for _, transport := range transports {
		if transport != transportCheck {
			event.Transports = append(event.Transports, transport)
		}
	}
}

func (event *L9Event) HasTransport(transport string) bool {
	for _, check := range event.Transports {
		if check == transport {
			return true
		}
	}
	return false
}

func (event *L9Event) HasSource(source string) bool {
	for _, check := range event.EventPipeline {
		if check == source {
			return true
		}
	}
	return false
}

func (event *L9Event) AddSource(source string) {
	event.EventPipeline = append(event.EventPipeline, source)
	event.EventSource = source
}