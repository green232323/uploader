# uploader
Project for uploading streams of data and store them

This project contains 2 services: 

 - **clientAPI**
 - **portDomainService**

### clientAPI
This service contains REST API with 2 endpoints:

 - **BASE_URL/ports/ [POST]**
	This one consumes json object and store it in another service
 - **BASE_URL/ports/<id> [GET]**
	This one returns object from previously saved one by id

This service runs on 8090 port, that can be checked in docker-compose.yml. It is the only entry to the application.

### domainService
This service is a gRPC server that also has 2 operations:

 - **LoadPorts**: get stream of ports and save them in DB
 - **Get**: returns port by id from DB

### DB
As a database I have used mongoDB but DB here doesn't really matter, it just should implement repository interface. Data from db stores locally in `./uploader/data` folder.
### Usage
To run the application, run in terminal next command:
```./run.sh```.

To run tests, run in terminal next command:
```./tests.sh```




