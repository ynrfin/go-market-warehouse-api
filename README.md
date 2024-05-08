This is a toy project for portofolio and exploration.

To run this, in linux / mac :
```
chmod 755 dbuild.sh
./dbuild.sh
```

For windows, just build the docker image using Dockerfile and then run docker compose. As stated in `dbuild.sh`

# Docker Compose: sometimes the application not wait for postgre

Keep rerun the compose file, I still could not find the right way to wait for postgres service to be started and available.

Things that I tried:
- use `depends_on` and `condition = service_healthy`  which return error `dependency failed to start: container local-pg-16 has no healthcheck configured`

Solution:
- use healthcheck from above, which create new section on the postgres service to define `healthcheck`
