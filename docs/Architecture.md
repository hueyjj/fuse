# Sunder architecture

![Architecture](https://raw.githubusercontent.com/hueyjj/fuse/master/docs/draw.io/SunderArchitecture.png)

## Web view
The front end will be implemented through the browser (Chrome, Firefox).

## Front module
A front module will be any bundle of HTML/CSS/JS that can be loaded dynamically.

## REST API server
n/a

## Router
A router will route the frontend requests to the proper back module.

## Back module
A back module is a Golang executable that will process JSON objects from stdin and raw stdout from the components it connects to. This is the meat of this architecture.  

## Post processor
A post processor is a part of a back module where it simply executes some commands after processing stdout from the back component.

## Back component
A Back component is anything that can communicate with a Back Module. This can be literally almost everything. A python/c++ server, executable, ruby script. The Back Component must be spawnable as a child process.

# Main conditions to satisfy
## Module extensibility
Front module and back module must be decoupled. They must be replacable in any project. These should provide the bare minimum to interface with the server and the back component, respectively.

## Stateless front end
The front end should not persist any data after a system startup. Business logic should be moved to the server and their respective module, intefacing some back component.

# Developing
## What a developer should do
This architecture is meant to be the loosest way for any developer to add their own technology. Mainly, a developer would need to add three things.

1. Frontend component (ex. React)
2. Back module (ex. executable written in Go)
3. Backend component (ex. grep, youtube-dl, python server)

# Notes
## Why Sunder architecture?
Trying to find a way to build local/desktop applications without having to deal with QT or Electron (note: Sunder architecture is not made to be a remote server-client relationship, it's local server-client relationship).

## Why server?
The front end needs to be decoupled so that any front end can be implemented through REST API. Also, all operations will be faster when moved to the backend.

## Why not [Electron](https://github.com/electron/electron) or [QML](https://en.wikipedia.org/wiki/QML)?
Might as well use the browser since most people have them up all the time anyway (?). Though, any frontend can be used since everything is just network calls anyway. It is up to the frontend on how to handle the data. 

## Do not use post process to handle large data streams
Video and audio streaming, or transporting any large amount of data should not be done through a back module. Most likely, a back module must change communication method/protocol with the server and back component to achieve optimum resource consumption and low latency.

## Golang "plugin" not available on Windows
The plugin package allows a Go program allows plugin to be loaded at run time. Unfortunately, there is only support for Mac and Linux with no further intended development (at this time, June 23, 2018 2:31 PM PST).

We opt to just communciate through a Golang executable (back module) with JSON data.

## We do not want to modify/restart the main Golang server everytime a module is installed
We should not assume a user has any build tools installed. Rebuilding the project would be out of the question then.
 
## Why Golang?
Goroutines are apparently very sexy.

## Why not another communication protocol like JSON RPC?
This might change. Everything is experimental. We suspect that trying to force each component to adhere to the Sunder communcation methods (stdin, stdout, JSON) would be too limited, so most likely a more general approach for communication will be defined later.

