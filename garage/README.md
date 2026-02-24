# GARAGE HQ Object Storage Service

For this simple project, in order to simulate an s3 bucket (without actually having to register and create one in aws), I use [garagehq object storage service](https://garagehq.deuxfleurs.fr).
There's a full documentation page, with a [Quick Start guide](https://garagehq.deuxfleurs.fr/documentation/quick-start), however, I will enumerate here the simple steps to get this up and running (at least the basic things that this project needs)

Build the docker image and start the docker container with docker-compose

Inside the container, there's a "garage" binary that allows the service administration. You can access it with the following command:
```bash
docker exec -it md-drive-garage-1 /garage
```
Note: I was getting this error at first:
```
OCI runtime exec failed: exec failed: unable to start container process: exec: "C:/Program Files/Git/garage": stat C:/Program Files/Git/garage: no such file or directory
```
This happens due to some address translations when using bash in windows. If this happens to you, you can solve it by using the powershell or running the command like this:
```bash
MSYS_NO_PATHCONV=1 docker exec -it md-drive-garage-1 /garage
```
## Init script

I created an initialization script in the garage folder. Execute it passing the container name as a parameter, like this:
```bash
./garage/init.sh md-drive-garage-1
```
This script will run the configuration commands, create a layout, a bucket and the api keys. You have to copy the keys ID and secret and use to connect to the service. For example, copy them into the .env file in the files-service

It should only be necessary the first time you run start the service. This means you may want to start it separately from the rest of the services (because you will then have to load the keys in the other services' environments, before you can start them up).

Finally, as long as you keep the volumes mounted correctly, the data and configuration will stay persistent.