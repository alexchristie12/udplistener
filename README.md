# udplistener

This is the test rig that I have written for CC3501 Practical 05. It's purpose
is to listen for UDP messages on a particular port, and print them to the
terminal.

## How to use

You can send UDP messages to the port configured by a `PORT` environment
variable. This is normally configured in a `.env` file in the root project
directory.

## Installation

There is only one dependency for this project, the Go programming language.
You can download it from [the official website](https://go.dev/). Once you
have done that you can clone this repository:

```bash
git clone https://github.com/alexchristie12/udplistener.git
```

Then create a `.env` file in the project directory and define the `PORT`
environment variable:

```bash
cd udplistener
touch .env
# Then define the PORT variable

# In .env
PORT=":4321"
```

We then collect the dependencies and build it:

```bash
go mod tidy
go run .
```

You should get something like:

```bash
2024/08/18 13:56:38 UDP server listening on: udp://127.0.0.1:4312
```

When we send a message to `127.0.0.1:4312`, this program will print it out to
the terminal:

When we send the message:

```bash
cowsay "Hello CC3501" | nc -u -w1 127.0.0.1 4312
```

When the server receieves the message:

```bash
2024/08/18 14:19:10 Received message from 127.0.0.1:48713:
 ______________
< Hello CC3501 >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
