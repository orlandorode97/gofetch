# gofetch
gofetch is an alternative neofetch, screenfetch, and pfetch CLI tool that retrieves the current operative system information.

## Note ⚠️
Before getting started is important that your current terminal supports `UTF-8`, otherwise, the `gopher` of the `gofetch` tool will not be encoded causing some inconsistencies.

## Getting started
To have `gofetch` in your machine follow the next steps:
1. Clone the project under the `$GOPATH`.
```sh
$ git clone https://github.com/orlandorode97/gofetch.git
```
2. The project contains a [Makefile](https://github.com/orlandorode97/gofetch/blob/main/Makefile) with built-in commands to create the binary for your current os and architecture.
3. Type the command to build the project.
- `build`
5. Example
    ```sh
    $ make build
    --> Building gofetch binary for linux:amd64
    --> gofetch for linux:amd64 built at /usr/home/orlandoromo/go/src/gofetch
    ```
5. The previous command generates the binary `gofetch` at the root of the project. Copy it to the folder under `/usr/local/bin` either on `linux` or `mac`:
```sh
$ sudo cp gofetch /usr/local/bin
```

6. Go for it:
<img width="600" alt="Screen Shot 2022-02-14 at 15 43 03" src="https://user-images.githubusercontent.com/34588445/177675584-51c65cac-538c-4d30-b77a-75d8c8c6f042.png">
