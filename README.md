To run, build and launch the app:
```
make && ./chat
```
or
```
go build -o chat main.go && ./chat
```
or 
```
go run main.go
```

This app uses two ports and they are currently hardcoded:
* 8081 is used for web files
* 8082 is used for web sockets

These ports must be available and not blocked by a firewall.

Once the app is running, use a browser to connect to http://{hostname}:8081. For example:
```
http://localhost:8081
```

Things that could be added to make this more professional:
* Logging
* Better error handling
* Better UI
* Configurable ports
* More tolerant sockets. Could detect disconnections and automatically reconnect. Could possible store missed messages for a period of time as well.
* Unit tests
* End to End tests
* Buffer sizes - currently unspecified so in theory they can grow to any size. The upside is that there shouldn't be wasted buffer space. The downside is frequent memory allocation.
* Hostname display - locally I sometimes see connection addresses as 127.0.0.1 or [::1]. These should probably be altered to display an external address
* This readme could use formatting, quick start, etc.

Extra Credit:
* TLS - not supported  
* Makefile supports cross-compiling for raspberry pi. It builds, but I don't have  Pi for testing.
* Debian package - I know of these, but I'm not familiar enough to tackle this, especially without a Debian instance running for testing.
* There are additional dependencies and can be installed with "make dep".
* Scaling - This shouldn't have any significant scaling issues. 10,000+ users should already work. Potentially there might be a concurrency issue: When broadcasting to many clients, the a mutex may be locked for a long period of time. If excessive, a potential solution would be to make a list of socket pointers we will need for sending, and then send to that list after unlocking the mutex. If the user count grows to a point that the OS can't support enough open sockets, then the solution will need to be more creative, anda possibility would involve multiple servers available to clients, which in turn communicate with a centralized server. This would require a significant rewrite.
* Pi lights - I don't have a Pi for testing so I didn't attemp this. I did find code that should work and it looks simple enough, but I don't have a way to test the code.
 
