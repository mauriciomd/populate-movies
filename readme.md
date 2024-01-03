# Populate Movies

## About
This repo implements a simple CLI written in Go for reading a CSV file and importing that data into a Postgres database. An example of a CSV file can be found in ```data.csv```.
It can be used in Sequential and Concurrent mode. When the concurrent flag is used, it uses the Pool Pattern to insert the data. 

## How to run
First, clone this repository
```
    git clone https://github.com/mauriciomd/populate-movies.git
```

Next, starting the Postgres by using Docker
```
    docker-compose up
```

Then, to execute the program, execute
```
    go run cmd/main.go 
```