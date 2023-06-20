from viam.robot.client import RobotClient
from viam.rpc.dial import Credentials, DialOptions


async def connect(host, payload, channel=False):
    creds = Credentials(
        type='robot-location-secret',
        payload=payload
    )
    opts = RobotClient.Options(
        refresh_interval=0,
        dial_options=DialOptions(credentials=creds)
    )
    if channel:
        await RobotClient.with_channel(host, opts)
    return await RobotClient.at_address(host, opts)
