package main

import "shoplist/pkg/storage"

type Config struct {
	DataStore  storage.Storage
	jwt_secret []byte
}
