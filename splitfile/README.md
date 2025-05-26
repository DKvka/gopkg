# GetParts:

## Get parts:
```go
	parts, err := splitfile.GetParts(filePath, partCount, maxLineLength)
	if err != nil {
		// Stop process, error in GetParts means you need to start over
		log.Fatal(err)
	}
```

## Start processing:
```go
	// ResultType is an empty interface, so it can be any type
	results := make(chan splitfile.ResultType)
	for _, part := range parts {
		// You need to define processLine according to documentation
		go splitfile.ProcessPart(filePath, part, results, processLine)
	}
```

## Wait for results:
```go
	for i := 0; i < len(parts); i++ {
		res := <- results

		// Process result as needed
	}
```

You could also use a range over the channel, but you need to implement a way for the results channel to close after every part is processed or the loop will block forever. This way is easier.

# ProcessPart:

