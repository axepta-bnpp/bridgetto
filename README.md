# bridgetto
A small and simple self-contained, config-less TCP bridge. Mostly some CLI sugar over the code written in this article: http://blog.evilissimo.net/simple-port-fowarder-in-golang

Examples:

* bridgetto 0.0.0.0:9999 remotehost:27017  #bridges the local 9999 port to a remote host on port 27017
* bridgetto 0.0.0.0:9999 unix:/var/some/socket  #bridges the local 9999 port to a local unix socket in /var/some/socket
