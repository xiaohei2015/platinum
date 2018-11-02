# platinum - A simple Go API extension for Datetime

## Introduction

A very very simple Datetime API writen by golang

## Install

```
go get -u github.com/xiaohei2015/platinum
```

## Usage

```go
fmt.Println("Now:", platinum.Now())
// prints: Now: 2018-11-02 11:07:19
fmt.Println("Right now:", platinum.Now().ToDateTimeString())
// prints: Right now: 2018-11-02 11:07:19
fmt.Println("Date:", platinum.CreateFromDate(2018, 10, 31))
// prints: Date: 2018-10-31 00:00:00
fmt.Println("DateTime:", platinum.CreateFromDateTime(2018, 10, 31, 13, 22, 33))
// prints: DateTime: 2018-10-31 13:22:33

fmt.Println("Today:", platinum.Today())
// prints: Today: 2018-11-02 00:00:00
fmt.Println("Tomorrow:", platinum.Tomorrow())
// prints: Tomorrow: 2018-11-03 00:00:00
fmt.Println("Yesterday:", platinum.Yesterday())
// prints: Yesterday: 2018-11-01 00:00:00

fmt.Println("Location Change:", platinum.Now().Location(time.UTC))
// prints: Location Change: 2018-11-02 03:07:19
fmt.Println("Format Change:", platinum.CreateFromDate(2018, 10, 31).Format("2006/01/02"))
// prints: Format Change: 2018/10/31
```
