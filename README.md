This is a simple test on using go-reflex to provide a local environment with docker-compose for multiple go services.

it relies on go-reflex, which provides a docker image that rebuilds any go project upon change within the container, therfor allowing us to have a stable service environment but still be able to work on any service without overhead.
