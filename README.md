# Top ten words extractor
Given a body of text, this servlet will return the top ten words together with the number of times
it was used.

The struct that is returned:
```golang
    type kv struct{
        Key string
        Value int
    }
```

## Notes
This was compiled using GoLang 1.10.2.

You need at least GoLang 1.8 and above for this to work.
It uses the new(relative) `sort.Slice`.
