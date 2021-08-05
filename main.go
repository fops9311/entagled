package entangled

func Entangle() (in func(interface{}), out func() interface{}) {
	var data interface{}
	in = func(d interface{}) {
		data = d
	}
	out = func() interface{} {
		return data
	}
	return in, out
}

func ToWay() (in ToWayPipeline, out ToWayPipeline) {
	in.emiter, out.reciever = Entangle()
	out.emiter, in.reciever = Entangle()
	return
}

type ToWayPipeline struct {
	emiter   func(interface{})
	reciever func() interface{}
}
