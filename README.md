#Registry app

Simple client-server app for a registry.

## Client

Main function the client application is send form information to server for store in store. Connect to server open
on each button click, because application have low load and not necessary monitoring network disconnection in the background.

Client configuration file format :

```
srv = "localhost:8080"
log = "log.txt"
```

```srv``` - server network address
```log``` - path to log file

## Server

Server saved client information to db and generate Excel file with report. 

Server configuration file:

```
srv = "localhost:8080"
log = "log.txt"
db = "productdb.db"
```

```srv``` - server network address
```log``` - path to log file
```db``` - path to sqlite db
 