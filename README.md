# Hostname & DNS Configuration Service
This project includes the gRPC service that configures hostname & DNS server list of the machine it runs on and the CLI client.

## Features
- **gRPC Service:** Listens for gRPC requests.
- **.proto File:** Responsible for generation of gRPC service code in Go and swagger documentation.
- **REST API:** Enabled by the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) library.
- **CLI Client:** Command line interface written using [Cobra](https://github.com/spf13/cobra) framework that allows user to interact with the service.

## Installation
1. Clone the repository:
    ```shell
    git clone https://github.com/stepan-tikunov/hostname-dns-configurer.git
    cd hostname-dns-configurer
    ```
2. Install dependencies & build service binary:
   ```shell
   cd service && go build && cd ..
   ```
3. Install dependencies & build client binary:
   ```shell
   cd client && go build && cd ..
   ```

## Usage

### Service
The steps to start the service vary depending on your operating system.
#### Linux
If you're running Linux, starting the server is as simple as running
```shell
sudo ./service/service
```
Sudo is required in order for some features to work. To specify custom ports for gRPC endpoint and REST API gateway, use flags
`--api` and `--gateway`:
```shell
sudo ./service/service --api 1234 --gateway 5678
```
#### Other OS
However, if you're using other OS, some functionalities may be not supported.
To test out all the features, you will need to run the app in a Docker container. Install
Docker Compose and run this command:
```shell
docker-compose -f ./service/compose-dev.yaml up
```
This will automatically build and start the service.

### Client
The client binary is located in `./client/client` and supports various actions.

#### Get hostname
```shell
./client/client hostname get
```

#### Set hostname
```shell
./client/client hostname set new-hostname
```

#### List all DNS servers used by service
```shell
./client/client dns list
```

#### Add new DNS server
This command adds a server to the end of the list. 
```shell
./client/client dns add 8.8.8.8
```
To specify where the server must be inserted, use `--index n` flag
and the command will insert the server before n-th one in the list.
```shell
./client/client dns add --index 0 8.8.8.8
```
#### Delete the DNS server
```shell
./client/client dns delete --index 0
```

## License
Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Author
Stepan Tikunov, 2024
