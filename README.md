# PowerFlex Backend Challenge
## Dependencies
In order to run/test this service properly, you need the following:
1. Docker and docker-compose
2. migrate and mockery (can be installed via `make install`)
3. Something to test the REST API with (such as Postman)
## How to run
Running `make run` will create the docker image for the service, then run it on localhost:8080. Use Postman
or something similar to test each of the endpoints. The database is be pre-seeded, with the "GetById" 
endpoints referring to a serial ID for each table.
