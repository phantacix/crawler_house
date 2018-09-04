package engine

import "log"
import "../fetcher"

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		webDriver := fetcher.Fetch(r.Url)

		parseResult := r.ParserFunc(webDriver)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("got item %v", item)
		}
	}
}