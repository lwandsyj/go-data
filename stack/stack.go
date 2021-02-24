package stack

//Statck ...
type Statck struct {
	data []interface{}
	size int
}

//New init ...
func New() *Statck {
	return new(Statck)
}

//Default ...
func Default() Statck {
	return Statck{
		make([]interface{}, 0),
		0,
	}
}

//Default1 ...
func Default1() Statck {
	return Statck{
		size: 0,
	}
	return Statck{
		data: make([]interface{}, 0),
		size: 0,
	}
}

//Push statck ...
func (statck *Statck) Push(data interface{}) {
	if statck == nil {
		statck = New()
	}
	statck.data = append(statck.data, data)
	statck.size++
}

//Pop ...
func (statck *Statck) Pop() interface{} {
	if statck == nil || statck.size == 0 {
		return nil
	}
	rtn := statck.data[statck.size-1]
	statck.data = statck.data[:statck.size-1]
	statck.size--
	return rtn
}

//Size ...
func (statck *Statck) Size() int {
	return statck.size
}
