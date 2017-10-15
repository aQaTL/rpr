# RPR - Remote Program Runner

Listens on given port at given address and executes given command when request is made. 
I made it for my VirtualBox Host-Only Network to run command in my VM from the host.

## Warning!

It has absolutely no security, do not use it in public networks (probably not even in personal). 

## Usage

```
Usage of rpr:
  -address string
    	IP address from which the app listens (default "127.0.0.1")
  -cmd string
    	Command to run on detected connection
  -port uint
    	Port on which the app listens (default 1120)
```

