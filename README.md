---
![Project Status](https://img.shields.io/badge/Project%20Status-Work%20in%20Progress-orange)

# OpenAPI Mock Server Generator

This application generates and stubs a server and handlers based on an OpenAPI 3.0 specification. The generated server will be dockerized and served on a specified port.

## Project Status
Please be aware that this project is currently in early development stage and is a work in progress. While it can already be used to generate a server, handlers, and Docker files from an OpenAPI 3.0 spec, there are still many areas that need improvement and it may not cover all use cases or edge cases. There are likely to be bugs, incomplete features, incorrect documentation, or other discrepancies.

We welcome any contributions, suggestions, and feedback. If you encounter any problems or have any ideas for improvement, please feel free to open an issue or submit a pull request.

## Program Args/Flags

The application uses the following program args/flags:

- `openAPISpecPath` - Path to the OpenAPI spec file
- `handlerFuncTemplatePath` - Path to the template for the handler functions
- `handlersTemplatePath` - Path to the template for the handlers
- `handlersOutputPath` - Path to output the generated handler functions
- `serverTemplatePath` - Path to the template for the server
- `serverOutputPath` - Path to output the generated server
- `staticResponses` - Path to static responses (named the same as the endpoints) to be used for specified endpoints
- `structsTemplatePath` - Path to the template for the structs
- `structsOutputPath` - Path to output the generated structs
- `dockerfileTemplatePath` - Path to the template for the Dockerfile
- `dockerfileOutputPath` - Path to output the generated Dockerfile
- `dockerComposeTemplatePath` - Path to the template for Docker Compose file
- `dockerComposeOutputPath` - Path to output the generated Docker Compose file
- `mockAPIPort` - The port on which to expose the generated mock API
- `dirPath` - The directory in which to generate the server
- `modName` - The name of the Go module being generated

## Usage

You can specify the values of these flags directly in the command line like so:

```bash
go run main.go --openAPISpecPath=./your-spec-path.yaml
```

Replace `your-spec-path.yaml` with the path to your OpenAPI 3.0 specification.

## Functionality

The application works as follows:

1. Generates structs based on the OpenAPI 3.0 specification
2. Generates handlers based on the OpenAPI 3.0 specification
3. Generates a server that uses these handlers
4. Generates a Dockerfile and a Docker Compose file
5. Spins up the mock api server on the specified port
---