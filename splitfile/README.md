## GetParts example use

```
  parts, err := splitfile.GetParts(filePath, workerCount)
	if err != nil {
		log.Fatal(err)
	}

  resultChan := make(chan result)
	for _, part := range parts {
		go processPart(filePath, part.Offset(), part.Size(), resultChan)
	}
```
