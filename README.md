# plural - Simple Go API for Pluralisation.

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/plural)
[![Build Status](https://travis-ci.org/rickb777/plural.svg?branch=master)](https://travis-ci.org/rickb777/plural)

Package plural provides simple support for localising plurals in a flexible range of different styles.
There are considerable differences around the world in the way plurals are handled. This API is
a simple but competent API for catering with these differences when formatting text to
present to people.

This package is able to format **countable things** and **continuous values**. It can handle integers
and floating point numbers equally and this allows you to decide to what extent each is appropriate.

For example, "2 cars" might weigh "1.6 tonnes"; both categories are covered.

This API is deliberately simple; it doesn't address the full gamut of internationalisation. If that's
what you need, you should consider products such as https://github.com/nicksnyder/go-i18n instead.
