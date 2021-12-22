# Apollo

Apollo is a simple deep health check system.  A typical health check service usually checks whether a given service is up or not by using a simple HTTP ping on a pre-configured service endpoint.  Apollo goes one step further by asking the service to report the health of its dependencies that are necessary for the application to work optimally.

## How does this work?

The system has two parts.  The apollo server that will actually trigger periodic health checks to the service over HTTP or TCP.  apollo SDKs will be integrated into the source code of the application to be monitored.  The SDK should have helper checks for common infrastructure components.  For example, there could be a `apollo-go-sql` extention that can perform a check from the application to the SQL server ideally using the same connection that is used by the application.  The SDK will inform the apollo-server of the health status of services. apollo-server will use its algorithms to detect the status of the service.
