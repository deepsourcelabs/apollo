<p align="center">
  <img src="logo.svg" width="600"/>
</p>

---

# Apollo

Apollo is a distributed deep health check system.

A typical health check service usually checks whether a given service is up or not by using a simple HTTP ping on a pre-configured service endpoint.

Apollo goes one step further by asking the service to report the health of its dependencies that are necessary for the application to work optimally.

## How does it work?

Apollo has two parts:
- The Apollo server
- Apollo SDKs

Apollo server is a standalone server which helps in:
- Registering services
- Triggering health checks

Apollo SDKs will be integrated into the source code of the application to be monitored.

## Development

Apollo server can be built using `Makefile`:

The `Makefile` assumes that you have a configuration file (`config.yaml` ) set up in the current directory. (check [how to set up configuration file](https://burntcarrot.github.io/apollo/docs/config)):

```sh
make dev
```

Other commands include:
- `local_build`: Build Apollo server (stored at `/tmp/apollo`)
- `local_run`: Runs Apollo server (requires configuration)
- `swagger`: Generates Swagger documentation
