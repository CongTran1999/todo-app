# Todo CLI app
The todo cli application using GO programing languague to list what you need to do and what have been done.

# Build
```
make build
```
If you want to run test, you can do this
```
make test
```
# Run
Build the app and run, currently save to a json file, maybe we will add a csv file or DB in config
```
./todo -add Do homework
```
# Usage
```
./todo <flags> input
  -add
        add a new todo
  -compelte int
        mark a todo complete
  -del int
        delete a todo
  -list
        list all todo
```
