import * as VIAM from '@viamrobotics/sdk';

async function main() {
  const host = '<HOST_ADDRESS>';

  const robot = await VIAM.createRobotClient({
    host,
    credential: {
      type: 'api-key',
	 // Replace "<API-KEY>" (including brackets) with your machine's api key
      payload: '<API-KEY>',
    }, 
	 // Replace "<API-KEY-ID>" (including brackets) with your machine's api key id
    authEntity: '<API-KEY-ID>',
    signalingAddress: 'https://app.viam.com:443',
  });
  
  console.log('Resources:');
  console.log(await robot.resourceNames());
}

main().catch((error) => {
  console.error('encountered an error:', error)
});
