# Go Command Broadcaster

A simple Go tool that lets a server broadcast shell/command-line instructions to multiple connected clients.

## Features
- Server can send commands to all connected clients
- Clients execute commands and return output
- Works on Windows, Linux, and macOS

## Usage

### 1. Build
Server
```
go build -o server.exe main.go server.go
```
Client
```
go build -o client.exe main.go client.go
```

### 2. Run
Server
```
./server.exe
```
The server will start on port 5000 and wait for clients to connect.

Client
```
./client.exe
```
Once connected, the client will execute commands from the server and print the output.
