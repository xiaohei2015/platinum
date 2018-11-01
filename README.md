# platinum - A simple Go API extension for Datetime

## Introduction

A very very simple Datetime API writen by golang

## Install

```
go get -u github.com/xiaohei2015/platinum
```

## Usage

```go
// Default locale is based on env vars or en_US if none are set.
fmt.Println("Now:", platinum.Now())
// prints: Now:2018-11-01 03:58:31

fmt.Println("Right now:", platinum.Now().ToDateTimeString())
// prints: Right now:2018-11-01 04:04:08

fmt.Println("Date:", platinum.CreateFromDate(2018, 10, 31))
// prints: Date:2018-10-31 00:00:00
```
