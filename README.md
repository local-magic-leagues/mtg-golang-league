# mtg-golang-league
## Backend for mtg league

### Getting the package
Get the go package
```
go get github.com/local-magic-leagues/mtg-golang-league
```

### Setting up the API 
1. To sart the server for the API
```
go run main.go
```
2. Go to `http://localhost:8000/`

### Setting up postgresql (on mac)
1. Install postgres

```
brew install postgresql
```
- For linux, the command should be
```
sudo apt-get install postgresql postgresql-contrib
```

After installation, by default postgres creates:
    - a database named postgres
    - a super user for that dabase whose login info is the same as the user who did the installation
2. Install [pgAdmin](https://www.pgadmin.org/), or whatever other client you would like to use 
    - You could also use the psql shell 
    ```
    psql [databasename]
    ```


