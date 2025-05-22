## GetParts example use

```
	parts, err := splitfile.GetParts(filePath, workerCount)
	if err != nil {
		log.Fatal(err)
	}
	
	results := make(chan result)
	for _, part := range parts {
		go processPart(filePath, part.Offset(), part.Size(), results)
	}
	
	// Wait for results to start coming in
	for i := 0; i < len(parts); i++ {
		res := <- results
		// Process result
	}
	// You could also use a range over the channel, but you need to
	// implement a way for the results channel to close after every
	// part is processed or the loop will block forever
```
