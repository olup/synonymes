# French Synonyms API

This is a simple project in Go, the quest or a simple, efficient and free API to find french synonyms.

This works by opening an open office dictionary of french synonyms.

Then it either :

- Generate all JSON files in a `build` directory, to be served by netlify as a static site (synonyms don't change that often) (demo : <https://synonymes.netlify.com/démonstration>)

- Serve the right word with a serverless handle deployed at `now.sh` (demo : <https://synonymes-api.now.sh/démonstration>) and leverage long term cache on their CDN.
