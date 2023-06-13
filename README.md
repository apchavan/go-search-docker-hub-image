
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) [![GoReportCard](https://goreportcard.com/badge/github.com/apchavan/go-search-docker-hub-image)](https://goreportcard.com/report/github.com/apchavan/go-search-docker-hub-image) ![](https://img.shields.io/badge/OS-Linux-orange) ![](https://img.shields.io/badge/OS-macOS-black) ![](https://img.shields.io/badge/OS-Windows-blue) [![Go Reference](https://pkg.go.dev/badge/github.com/apchavan/go-search-docker-hub-image.svg)](https://pkg.go.dev/github.com/apchavan/go-search-docker-hub-image)

# Search for Docker Hub Images

It's a helper mini-tool that can be used to search images available on [Docker Hub](https://hub.docker.com/) within terminal. It works the same way as [`docker search`](https://docs.docker.com/engine/reference/commandline/search/) command but adds some interactive controls while searching.

**Note**:

- [Docker](https://www.docker.com/) must be installed & running to use this application, because it uses [official Docker Engine's Go SDK](https://docs.docker.com/engine/api/sdk/) to communicate with the Docker Daemon/Engine.

## Main Features

- Interactive terminal UI to search for any Docker Hub images.

- Easily set filters or limits for the search.

- Automatically result sorting based on highest to lowest stars count.

- At present, no credentials/sign-in is required! Just run Docker Daemon/Engine & have internet connection...

- Fully open source under MIT license!

## Working Demo

## Dependencies

Currently the project depends on:

1. [docker](https://docs.docker.com/engine/api/sdk/) - Docker provided API for interacting with the Docker daemon (called the Docker Engine API).

2. [tview](https://github.com/rivo/tview) - Terminal UI library with rich, interactive widgets — written in Golang

3. [go-pretty](https://github.com/jedib0t/go-pretty) - Table-writer and more in golang!

## Build Binary

After installing [Go](https://go.dev), enter below command from project's root,

- On Linux/UNIX,

    `go build -o search_docker_hub_image ./cmd/search_docker_hub_image.go`

- On Windows,

    `go build -o search_docker_hub_image.exe ./cmd/search_docker_hub_image.go`

## Run Directly with Source Code

After installing [Go](https://go.dev), clone/download this project & enter below command from project's root,

`go run ./cmd/search_docker_hub_image.go`
