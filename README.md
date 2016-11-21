## TypeReact
---

Environment:

![Go version](https://img.shields.io/badge/go-1.6.2-blue.svg)

![Node version](https://img.shields.io/badge/npm-3.10.8-orange.svg)

![TS version](https://img.shields.io/badge/typescript-2.0.7-brightgreen.svg)

![make version](https://img.shields.io/badge/make-3.8.1-lightgrey.svg)

Current Health:

[![CircleCI](https://circleci.com/gh/blinkanalytics/typereact.svg?style=shield&circle-token=083647c2df4d3b0478290e959e70c1fd9efd38c7)](https://circleci.com/gh/blinkanalytics/typereact)

[![Go report card](https://goreportcard.com/badge/github.com/blinkanalytics/typereact)](https://goreportcard.com/report/github.com/blinkanalytics/typereact)

# Why?

Due to the great flux of JS community efforts, most starter kits we have found were too complex or outdated, and none had a simple means of dealing with Babel/ES6, Flow, Eslint and Webpack.

There are also some design decisions that we wanted to make for all internal projects, that are detailed in the following section. The goal of this starter kit is to simply provide a rapid way to get started using the core technologies of React, TypeScript and Go.

# Assumptions / Decisions

+ We want strongly-typed code for the front end. Thus [React](https://facebook.github.io/react/) modules are built using [TypeScript](https://www.typescriptlang.org) and bundled using [Webpack](https://webpack.github.io/).
+ We want immutable data structures for their advantage in building reactive architectures. Thus we have installed [Immutable.js](https://facebook.github.io/immutable-js/).
+ We want strongly-typed, scalable, performant and easy-to-deploy code for the server. Thus we have written it in [Go](https://golang.org).
+ We presume that the front end may be built using a predictable state-container like [Redux](http://redux.js.org/), but since this adds complexity we leave the decision up to the developer.

# Instructions:

At your terminal, simply type
```sh
make dev
```
This will install all dependencies and start the local Go server.

If you prefer hot-reloading or only need to work on the client, consider also installing something like [webpack-dev-server](https://www.npmjs.com/package/webpack-dev-server).

## Future iterations:

Please follow [the issues section](https://github.com/blinkanalytics/typereact/issues) to see where it will go from here and to report any bugs/suggestions. Contributions and forks are most welcome!

## License

Copyright (C) 2016 Blink Analytics (TM), distributed per [the MIT license](https://github.com/blinkanalytics/typereact/blob/master/LICENSE).
