package flyweight

type FlyWeight struct {
	Name string
	Age  int
}

func NewFlyWeight(name string, age int) *FlyWeight {
	return &FlyWeight{
		Name: name,
		Age:  age,
	}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string, age int) FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name, age)
		f.pool[name] = weight
	}
	weight.Age = age
	return *weight
}
