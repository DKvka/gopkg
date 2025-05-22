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
```
