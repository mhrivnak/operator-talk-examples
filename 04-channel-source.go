var ctrl controller.Controller

func main() {
	events := make(chan event.GenericEvent)

    go monitorExternalStuff(events)

	_ := ctrl.Watch(
		&source.Channel{Source: events},
		&handler.EnqueueRequestForObject{},
	)
}
