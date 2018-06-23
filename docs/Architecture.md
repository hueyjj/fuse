# Sunder Architecture

![Architecture](https://raw.githubusercontent.com/hueyjj/fuse/master/docs/draw.io/SunderArchitecture.png)

## Back Module
A Back Module is a Golang executable that will process JSON objects from stdin and raw stdout from the components it connects to. This is the meat of this architecture.  

## Post Processor
A post processor is a part of a back module where it simply executes some commands after processing stdout from the back component.

## Back component
A Back component is anything that can communicate with a Back Module. This can be literally almost everything. A python/c++ server, executable, ruby script. The Back Component must be spawnable as a child process.

## What a developer should do
This architecture is meant to be the loosest way for any developer to add their own technology. Mainly, a developer would need to add four things.

1. Frontend component (ex. React)
2. Back module (ex. executable written in Go)
3. Backend component (ex. grep, youtube-dl, python server)