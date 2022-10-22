package controller

func Reconcile () {
    if currentState != desiredState {
        doALittleWork()
        setStatus()
    }
}
