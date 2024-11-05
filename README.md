# Simple Listener

This is a simple listener that listens to a port and prints the data it receives. For example, if you require an application to send data, but you don't have the receiving application, you can use this listener to see the data being sent and work as a "mock" of the receiving application.

## Usage

The configuration is a simple yaml file named `config.yaml` with the following content:

```yaml
appPort: 3424 # The port the listener will listen to
endpoint: "/hello" # The endpoint where data should be sent
timeZone: "America/New_York" # The timezone to use for the logs
```

Then you simply need to run the listener with the following command:

```bash
go run main.go
```

> **Note:** This assumes that you have Go (`>=1.22.1`) installed in your machine.


## Modification

If you want to modify the listener, you can change the `main.go` file to your needs. The repository is [MIT licensed](LICENSE), so you can do whatever you want with it.