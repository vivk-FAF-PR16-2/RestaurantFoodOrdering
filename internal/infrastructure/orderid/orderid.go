package orderid

// region Singleton

type IProvider interface {
	Get() int
}

var instance IProvider

func Get() IProvider {
	if instance == nil {
		instance = newProvider()
	}
	return instance
}

// endregion

type idProvider struct {
	id int
}

func newProvider() IProvider {
	return &idProvider{
		id: 0,
	}
}

func (provider *idProvider) Get() int {
	defer func() {
		provider.id += 1
	}()
	return provider.id
}
