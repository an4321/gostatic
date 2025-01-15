# GoStatic
A simple static site generator

## Installation
* Download the latest binary from the [releases page](https://github.com/an4321/gostatic/releases).
* Move the binary to your local bin directory:
  ```sh
  mv ./gostatic ~/.local/bin
  ```
* To check usage, run:
  ```sh
  gostatic --help
  ```

## Build Instructions
To build the project from source, follow these steps:

```sh
git clone https://github.com/an4321/gostatic
cd gostatic
go mod tidy && go mod vendor
go build .
```

## Commands

### `init`
Initializes the file structure for the static site generator. It creates the required directories and prepares the project for use.

### `build`
Builds all markdown files in the source directory (`./src`). It processes the markdown files and generates the static site output in the output directory (`./out`).

## Example Usage

To initialize the project directory structure:
```sh
gostatic init
```

To build the static site from markdown files:
```sh
gostatic build
```
