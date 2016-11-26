# zeus-server
Ragnarök Online server in Golang.

Board (Portuguese): http://zeusproject.com.br/board/
Discord: https://discord.gg/94APcTE

# Features and Goals

* All mechanics implemented in scripts
* Capability of removing all RO mechanics and using its engine to create another game
* Provide a "2016-century" architecture
* Support for load balancing
* Multi-zone support
* Ease of use
* Performance

# Technical Documentation

See [here](docs/index.md).

# Building

Just use standard Go convention:

    go build

# Requirements

* Golang 1.7
* PostgreSQL 9.5
* Redis 3.2

# Running

## With Docker (recommended)

Use the provided [compose file](docker-compose.yml) as base.

You can start all servers and databases as follow:

    docker-compose up

## Manually

Start all four servers:

    ./zeus-server server inter
    ./zeus-server server account
    ./zeus-server server char
    ./zeus-server server zone

## Client

Just connect with a standard client. Currently, only 2015-11-04 is supported.

# Contributing

See [here](CONTRIBUTING.md)

# License

Copyright (c) 2016 Zeus Project team

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

