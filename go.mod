module github.com/NubeIO/nubeio-rubix-app-rubix-broker-go

go 1.18

replace github.com/NubeIO/nubeio-rubix-lib-helpers-go => /home/aidan/code/go/nube/nubeio-rubix-lib-helpers-go


require (
	github.com/NubeIO/nubeio-rubix-lib-helpers-go v0.2.4
	github.com/mochi-co/mqtt v1.1.2
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)
