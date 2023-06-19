from viam.components.input import Control, Controller, EventType


async def handle_controller(controller):
    controls = await controller.get_controls()
    events = await controller.get_events()

my_controller = Controller.from_robot(robot=robot, name="my_controller")
await handleController(my_controller)
