package main

type Fetcher interface {
	fetch() ([]Car, error)
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
