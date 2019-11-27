![](https://github.com/peterwade153/ivents/workflows/Go/badge.svg)
# ivents

### ivents is a one stop center for all venues for all types events

### Installation.
Should have Go and Postgres installed

Clone the repository.
<pre>
git clone https://github.com/peterwade153/ivents.git
</pre>

### Environment variables
Create a database and create a `.env` from the `.env-sample` and replace its values with the actual values.

### Running application
Change directory into ivents then
<pre>
$ go run main.go
</pre>

API endpoint can be accessed. Via http://localhost:5000/

### Endpoints

Request |       Endpoints                 |       Functionality
--------|---------------------------------|--------------------------------
POST    |  /register                      |   User Signup   ( firstname, lastname, email, password)
POST    |  /login                         |   User Login    ( email, password)
POST    |  /api/venues                    |   Add Venue     ( name, description, location, capacity, category)
GET     |  /api/venues                    |   View Venues
GET     |  /api/venues/id                 |   View Venue
PUT     |  /api/venues/id                 |   Update Venue  ( name, description, location, capacity, category)
DELETE  |  /api/venues/id                 |   Delete Venue

### Running tests
<pre>
$ go test ./...
</pre>