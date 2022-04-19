module github.com/NubeIO/nubeio-rubix-app-rubix-broker-go

go 1.17

//replace github.com/NubeIO/nubeio-rubix-lib-helpers-go => /home/aidan/code/go/nube/nubeio-rubix-lib-helpers-go

require (
	github.com/mochi-co/mqtt v1.1.2
	github.com/sirupsen/logrus v1.8.1
	go.etcd.io/bbolt v1.3.6
)

require (
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/asdine/storm/v3 v3.2.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
