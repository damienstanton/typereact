## TypeReact

| health | [![CircleCI](https://circleci.com/gh/damienstanton/typereact/tree/master.svg?style=shield)](https://circleci.com/gh/damienstanton/typereact/tree/master) [![Go report card](https://goreportcard.com/badge/github.com/blinkanalytics/typereact)](https://goreportcard.com/report/github.com/blinkanalytics/typereact)|
| ---------| ----- |
| env | ![Go version](https://img.shields.io/badge/go-1.7.3-blue.svg) ![Node version](https://img.shields.io/badge/npm-3.10.8-orange.svg) ![TS version](https://img.shields.io/badge/typescript-2.0.7-brightgreen.svg)|

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
npm run dev
```
This will install all dependencies and start the local Go server.

## Roadmap:

A common request has been a hot-reload server, so [issue #4](https://github.com/blinkanalytics/typereact/issues/4) tracks my progress on this effort. In the meantime, consider installing something like [webpack-dev-server](https://www.npmjs.com/package/webpack-dev-server).

## Issues:

Please follow [the issues section](https://github.com/blinkanalytics/typereact/issues) to see where it will go from here and to report any bugs/suggestions. Contributions and forks are most welcome!

## License

Copyright (C) 2016 Blink Analytics (TM), distributed per [the MIT license](https://github.com/blinkanalytics/typereact/blob/master/LICENSE).
