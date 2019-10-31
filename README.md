# Event Finder API

Event Finder is a microservice that aggregate events from different sources and expose it via GraphQL endpoint. Currently Eventbrite, Meetup and Couchsurfing adapters are implemented.

## Running the server
```bash
# Pull the repo
git pull https://github.com/Kif11/event-finder-api
cd event-finder-api

# Init the project
make init

# Run the server
make serve
```

You will need to populate config.json files with your API keys.

## Making queries

By default the following endpoint is exposed `http://localhost:8080/graphql`. You can send POST request with the following GraphQL query

```
{
    events( lat: 37.770712, lon: -122.419565, totalTime: "3h" ) {
      name
      link
      lat
      lon
    }
}
```