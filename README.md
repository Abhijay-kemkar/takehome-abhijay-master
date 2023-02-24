## beam takehome

Welcome! The high level goal for this exercise is to build a client/server application in go-lang that synchronizes files from the client to the server (in one direction). A websocket server/client and basic protocol have been included.

### setup

- install go
- install [air](https://github.com/cosmtrek/air) for hot reloading
- to start client: `make client`
- to start server: `make server`

### overview

This repo includes a simple websocket client/server, with a simple protocol that enables them to send and receive serialized messages back and forth. A demo is included that shows a basic ECHO request that returns the same string that was sent by the client. Currently,
all the client does is infinitely send and receive the ECHO request.

The goal of this exercise to implement a very simple file synchronization protocol using the same protocol. For example, given an input directory, the client should scan that input directory, and serialize the files into messages containing the base64 contents of the file. On the server side, it should be able to handle that message and write the file to disk.

### goals

- Design a SYNC request that can send a file over the websocket

- Implement a basic asynchronous 'file watcher' that takes in an input directory, and detects when any files change. When a file changes the client should send the entire file over the websocket using the SYNC request

- Implement a SYNC handler on the server side that can read in the SYNC request and write the file to disk. It should return a response saying whether this process was successful
