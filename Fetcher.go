package main

type Fetcher interface {
	fetch() ([]Car, error)
	GetMiniCooperInfo() ([]Car, error)
}

type FakeFetcher struct {
	nextResponse      []Car
	nextResponseError error
}

func (f *FakeFetcher) fetch() ([]Car, error) {
	if f.nextResponseError != nil {
		return nil, f.nextResponseError
	} else {
		return f.nextResponse, nil
	}
}

func (f *FakeFetcher) GetMiniCooperInfo() ([]Car, error) {
	return make([]Car, 0, 0), nil
}
