#gofetch

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
4. Example
    ```sh
    $ make build
    --> Building gofetch binary for linux:amd64
    --> gofetch for linux:amd64 built at /usr/home/orlandoromo/go/src/gofetch
    ```
5. The previous command generates the binary `gofetch` at the root of the project. Copy it to the folder under `/usr/local/bin` either on `linux` or `mac`:
```sh
$ sudo cp gofetch /usr/local/bin
```
For Windows users it generates the binary `gofetch.exe` at the root of the project. Make a binary folder to place it with `make C:\bin`, move `gofetch.exe` to the created folder `move gofetch.exe C:\bin.`, add it to PATH with `setx PATH "C:\bin\;%PATH%"` and restart your terminal.

6. Go for it (animation created with [vsh](https://github.com/charmbracelet/vhs)):
<img width="600" alt="Screen Shot 2022-02-14 at 15 43 03" src="gofetch.gif">

# Contributions
Contributions are very welcome :shipit:
