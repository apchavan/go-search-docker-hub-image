
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) ![](https://img.shields.io/badge/OS-Linux-orange) ![](https://img.shields.io/badge/OS-macOS-black) ![](https://img.shields.io/badge/OS-Windows-blue)

## Search for Docker Hub Images using [Go](https://go.dev/)!

It's a helper mini-tool that can be used to search images available on [Docker Hub](https://hub.docker.com/) within terminal. It works the same way as [`docker search`](https://docs.docker.com/engine/reference/commandline/search/) command but adds some interactive controls while searching.

**Note**:

- [Docker](https://www.docker.com/) must be installed & running to use this application, because it uses [official Docker Engine's Go SDK](https://docs.docker.com/engine/api/sdk/) to communicate with the Docker Daemon.
