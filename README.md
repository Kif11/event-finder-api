# Event Finder API

Event Finder is a microservice that aggregate events from different sources and expose it via GraphQL endpoint. Currently Eventbrite, Meetup and Couchsurfing adapters are implemented.

To run:
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