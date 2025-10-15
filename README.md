# Citadel
Self-hosted mini search engine. Under early development.

## Adding new data
Add/modify websites with the db.csv file,
following the same structure for each line:

`Site's name,Site link,Description,Tags (separated by semicolon (;))`

## Building
```ksh
make
```

## Is it scalable?
All search results are stored in your RAM,
and no advanced optimization algorithm is being used.
So, no.

Feel free to give suggestions on the topic.
