package broker

func (inst *Broker) SetPort(port int) *Broker {
	inst.Port = port
	return inst
}
