// Define a function that handles the controller.
func handleController(controller input.Controller) {

    // Define a function to handle pressing the Start Menu Button "ButtonStart" on your controller, logging the start time.
    printStartTime := func(ctx context.Context, event input.Event) {
        logger.Info("Start Menu Button was pressed at this time: %v", event.Time)
    }

    // Define the EventType "ButtonPress" to serve as the trigger for printStartTime.
    triggers := [1]input.EventType{input.ButtonPress}

    // Get the controller's Controls.
    controls, err := controller.Controls(context.Background(), nil)

    // If the "ButtonStart" Control is found, register the function printStartTime to fire when "ButtonStart" has the event "ButtonPress" occur.
    if slices.Contains(controls, Control.ButtonStart) {
        err := controller.RegisterControlCallback(context.Background(), Control: input.ButtonStart, triggers, printStartTime, nil)
    }
     else {
        logger.Fatalf("Oops! Couldn't find the start button control! Is your controller connected?")
    }
}

func main() {
    utils.ContextualMain(mainWithArgs, logging.NewLogger("client"))
}


func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
    // ... < INSERT CONNECTION CODE FROM MACHINE'S CODE SAMPLE TAB >

    // Get the controller from the machine.
    myController, err := input.FromRobot(myRobotWithController, "my_controller")

    // Run the handleController function.
    err := HandleController(myController)

    // Delay closing your connection to your machine.
    err = myRobotWithController.Start(ctx)
    defer myRobotWithController.Close(ctx)

    // Wait to exit mainWithArgs() until Context is Done.
    <-ctx.Done()

    // ... < INSERT ANY OTHER CODE FOR MAIN FUNCTION >

    return nil
}