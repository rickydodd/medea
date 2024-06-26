# Medea

![Go 1.21](https://github.com/rickydodd/medea/actions/workflows/go-version-1.21.yml/badge.svg)
![Go 1.22](https://github.com/rickydodd/medea/actions/workflows/go-version-1.22.yml/badge.svg)
![Tests](https://github.com/rickydodd/medea/actions/workflows/tests.yml/badge.svg)

Medea is a media rating API.

## Design

### API Capabilities Canvas

![API Capabilities Canvas](./design/API-Capabilities-Canvas.png)

## To-Do

- [ ] In-memory implementation of basic functionalities.
  - `GET` all media.
  - `GET` a specific piece of media.
  - `POST` a specific piece of media.
  - `PATCH` the rating of a specific piece of media, calculating the new rating.
  - Not all basic functionalities need to be implemented before migrating to a database system.
- [x] Improve ergonomics of bug reporting and feature suggesting on GitHub issues.
  - Issue templates seem to be an appropriate route.
- [x] Basic continuous integration.
- [ ] Determine an appropriate database system.
- [ ] Basic authentication and authorisation.
- [x] Implement Makefile to ease the barrier to entry on contributing.
- [ ] `CONTRIBUTING.md` file or contributing section in `README.md` file.
