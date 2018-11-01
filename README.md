# Platinum - A simple Go API extension for Datetime

## Introduction

A very very simple Datetime API writen by golang

## Install

```
go get -u github.com/xiaohei2015/Platinum
```

## Usage

```go
// Default locale is based on env vars or en_US if none are set.
fmt.Println("Now:%s", Platinum.Now())
// prints: Now:2018-11-01 03:58:31

fmt.Println("Right now:%s", Platinum.Now().ToDateTimeString())
// prints: Right now:2018-11-01 04:04:08
```
