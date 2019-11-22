# piCollectData
A raspberry pi collects the data from IoT devices connected via wifi.

1. Select the network adapter to scan for the devices.
2. Check the data in the IoT devices (all the devices have a static IP address)
  * The devices address are listed inside one slice.
  * Each time the collector request data, generates an HTTP get to check the latest device value.
3. Based on the collected data compares the values with the past data and generates another request to another server.

fooServer contains dummy fake server to emulate the data from iot devices.
docker-compose is used to easily multiply the services. 

Create 5 instances of the devices using ips in the range 8080-8084
`docker-compose up -d --scale foo=5` 

Testing the code with iot:
Compile the dowloaded code `go run main.go -o serverCollect`