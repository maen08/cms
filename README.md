# GoLang CMS

![tests](https://github.com/gouniverse/cms/workflows/tests/badge.svg)

PREVIEW ONLY. NOT STABLE

A "plug-and-play" content managing system (CMS) for GoLang that does its job and stays out of your way.

## Introduction

All of the existing GoLang CMSs require a full installations from scratch. Its impossible to just add them to an exiting Go application, and even when added feel like you don't get what you hoped for.

This package allows to add a content management system as a module dependency, which can be easily updated or removed as required to ANY Go app. It is fully self contained, and does not require any additional packages or dependencies. Removal is also a breeze just remove the module.

## Features

- Templates (CMS)
- Pages (CMS)
- Blocks (CMS)
- Menus (CMS)
- Settings (CMS)
- Custom Types

# Simple Initialization

In its simplest initialization the CMS package accepts a GORM DB instance

```
cms.Init(cms.Config{
		DbInstance: gormDB,
})
```
