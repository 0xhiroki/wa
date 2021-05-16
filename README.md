This program will return a random Japanese pseudonym ("given name" + "family name").  
You can specify a prefix for both given name, and family name.

Example Usage:

```
go build

./wa -family=nakamoto -given=s
```

TODOs:  
- [ ] Store generated names in a file 
- [ ] Add more name data; some common ones are missing (e.g. "jiro" for first name) 
- [ ] Add a flag to specify the number of syllables 
- [ ] Add a check for .eth & .crypto domain availability for the generated name 
