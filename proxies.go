package main

import "fmt"

type Proxy struct {
	IP   string
	Port string
}

type ProxyFetcher struct{}

func (pf *ProxyFetcher) Get(limit int, protocols []string, anonymities []string) ([]Proxy, error) {
	// Implement the logic to fetch proxies based on the given criteria.
	// This is a placeholder for the actual implementation.
	return nil, fmt.Errorf("not implemented")
}

var PROXY bool = true // Assuming the package is imported and available.

func fetchProxy() []map[string]string {
	if !PROXY {
		return nil
	}

	fetcher := ProxyFetcher{}
	var proxies []Proxy
	var err error

	// Try to fetch anonymous proxies.
	proxies, err = fetcher.Get(10, []string{"HTTP"}, []string{"ANONYMOUS"})
	if err != nil {
		fmt.Println("No Anonymous proxies found. Switching to normal proxies ...")
		// If fetching anonymous proxies fails, try to fetch normal proxies.
		proxies, err = fetcher.Get(10, []string{"HTTP"}, nil)
		if err != nil {
			// Handle the error if fetching normal proxies also fails.
			fmt.Println("Failed to fetch proxies:", err)
			return nil
		}
	}

	// Convert the proxies to the desired format.
	var proxyList []map[string]string
	for _, proxy := range proxies {
		proxyMap := map[string]string{
			"http": "http://" + proxy.IP + ":" + proxy.Port,
		}
		proxyList = append(proxyList, proxyMap)
	}
	return proxyList
}
