# _MEGAVENTORY_

MEGAVENTORY API

# Overview

Using net/http and gin-gonic at the same time for creating rotes

# Prerequisites

golang 1.20.2 and postman required to run the project and test the routes

# Installation

Just run main.go and test the giving routes using postman

# Usage

HOST_HTTP=:8080
GET: ("/get")
POST: ("/post")
GET: ("/getinventory")

HOST_GIN=:9090
GET: ("/get/:sku")
POST: ("/product")
POST: ("/client")
POST: ("/inventorylocation")
POST: ("/tax")
POST: ("/discount")

# Branches

v8_Final
v7_FilesSeperated
v6_HttpGinTogether
v5_MoreMethods
v4_OnionArchitecture
v3_StructsInterfacesCreated
v2_FunctionsSeperated
v1_AllInMain
