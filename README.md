# Fabric SDK Go Server

This repository contains source code for HTTP server which is able to call the Hyperledger Fabric Chaincode. Both query and invoke operations are supported.

## Prerequisites
- Fully functional Hyperledger Fabric network
- Prepared Fabric SDK configuration file (see [sample][confsdksample])
- Private key of the user which has access to the Hyperledger Fabric network
- Go
- Git
- Make
- Docker Compose

## Getting started
- Clone this repository and `cd` there
    ```
    git clone https://github.com/KompiTech/fabric-sdk-server-go.git && cd fabric-sdk-server-go
    ```
- Copy user's private key into `keystore` directory
- Insert Fabric SDK configuration file into `config` directory and name it `sdk_config.yaml` ([sample][confsdksample] can be found in examples)
- Run the Docker Compose
    ```
    make compose-up
    ```
- Try test query (replace fields in `[]` with your configuration)
    ```
    curl "http://localhost:8080/api/v1/chaincode/query?chaincode=[CHAINCODE_NAME]&user=[USER_NAME]&channel=[CHANNEL_NAME]&fcn=[CHAINCODE_FUNCTION]&args=[CHAINCODE_ARGUMENTS]
    ```

## Configuration
The server itself offers environment variables listed below to set the configuration.
This configuration is not needed to override if you're using Docker Compose.
| Variable | Default Value | Description |
| ------ | ------ |------ |
| FABSRV_SERVER_ADDRESS | 127.0.0.0 | Address the server will listen on |
| FABSRV_SERVER_PORT | 8080 | Port the server will listen on |
| FABSRV_FABRIC_CONFIGPATH | ./sdk_config.yaml | Path to Fabric SDK configuration file |
| FABSRV_FABRIC_USER | admin | Default username used when ommited in HTTP requests |
The `keystore` directory containing user private keys must be placed in CWD from server binary has been started

### Docker Compose configuration
There are another set of variables used only for running the server in Docker Compose.
Those variables don't need to be overriden if you're using default `keystore` dir and `config/sdk_config.yaml` as a config file.
| Variable | Default Value | Description |
| ------ | ------ |------ |
| FSD_KEYSTORE_PATH | ./keystore | Keystore dir can be placed anywhere on the file system when using Docker Compose |
| FSD_CONFIG_PATH | ./config | Path to directory containing Fabric SDK configuration file named ${FSD_CONFIG_FILENAME} |
| FSD_CONFIG_FILENAME | sdk_config.yaml | Fabric SDK configuration file name |

## Using Postman for querying

This repository contains a Postman collection of request which can be imported to the app for making the requests easier. It contains query and invoke requests with predefined default parameters.

If you're not familiar with Postman, you can give it a try [here][postmandl]

[confsdksample]: <https://github.com/KompiTech/fabric-sdk-server-go/blob/master/example/sdk_config.yaml>
[postmandl]: <https://www.getpostman.com/downloads/>

## Using Swagger

This repository contains a file named swagger.json which contains Open API specification of server's API.