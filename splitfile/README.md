## GetParts example use

```
  parts, err := splitfile.GetParts(filePath, workerCount)
	if err != nil {
		return time.Duration(0), err
	}

  resultChan := make(chan result)
	for _, part := range parts {
		go processPart(filePath, part.Offset(), part.Size(), resultChan)
	}
```
