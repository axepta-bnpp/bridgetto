# smallbridge
A small and simple self-contained, config-less TCP bridge. Mostly some CLI sugar over the code written in this article: http://blog.evilissimo.net/simple-port-fowarder-in-golang

Example:
	smallbridge 0.0.0.0:9999 mongodbremoteip:27017  #bridges the local 9999 port to a remote mongodb host on port 27017
