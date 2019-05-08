package utils

//等泛型吧
func Map(f func(v interface{}) interface{}, l []interface{}) []interface{} {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}

func Reduce(f func(interface{}, interface{}) interface{}, l []interface{}) interface{} {
	v1 := l[0]
	for i := range l {
		v1 = f(v1, l[i+1])
	}
	return v1
}
