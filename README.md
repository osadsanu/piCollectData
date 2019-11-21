# piCollectData
A raspberry pi collects the data from IoT devices connected via wifi.

1. Select the network adapter to scan for the devices.
2. Check the data in the IoT devices (all the devices have a static IP address)
  * The devices address are listed inside one slice.
  * Each time the collector request data, generates an HTTP get to check the latest device value.
3. Based on the collected data compares the values with the past data and generates another request to another server.

