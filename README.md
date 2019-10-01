# Parking lot CLI program.

## Building
`./bin/setup` script can be used to test, compile and prepare the CLI binary. The script detects necessary dependencies
and error in case of missing dependencies such as compilers, go version, go path setup.  

## Running Functional tests.
Command `./bin/setup && bin/run_functional_tests` can be used to run both unit and functional tests.

## Usage
Run command `./bin/parking_lot` shell script. This script execs and runs the acutal binary.
Program can be either run in interactive or non-interactive mode.

### Interactive mode
`./bin/parking_lot`

### Non-Interactive mode
`./bin/parking_lot <path-to-input-file>`

## Features
- Streaming input parser that works without pre-allocating memory for the entire input.
- Supports color separated with space (Eg: "Light Coral").
- Unix exit codes in case of errors.
- Interactive and Non-Interactive modes.
- Example usage of mutexes to handle concurrency. This is not necessarily required for a CLI driven program from Stdin
or file as there would be single reader/writer. However in real world application, a parking lot program needs to handle
concurrency as there might be multiple entry/exit points or terminals.   