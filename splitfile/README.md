## GetParts examples:

### Get parts:
```
	parts, err := splitfile.GetParts(filePath, workerCount, maxLineLength)
	if err != nil {
		log.Fatal(err)
	}
```

### Start processing:
```
	// resultType can be anything you need to get back from the file
	results := make(chan resultType)
	for _, part := range parts {
		go processPart(filePath, part.Offset(), part.Size(), results)
	}
```

### Wait for results:
```
	for i := 0; i < len(parts); i++ {
		res := <- results
		// Process result
	}
```

You could also use a range over the channel, but you need to implement a way for the results channel to close after every part is processed or the loop will block forever. This way is easier.

### Split up file processing function boilerplate:
```
	func processParts(filePath string, offset, size int64, results chan resultType) {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	
		_, err = file.Seek(fileOffset, io.SeekStart)
		if err != nil {
			panic(err)
		}
		
		f := io.LimitedReader{R: file, N: fileSize}
	
		scanner := bufio.NewScanner(&f)
		for scanner.Scan() {
			// Process
		}
		resultChan <- stationStats
	}
```
