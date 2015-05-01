package main

type Room struct {
	Id string `json:"id"`
	Files []File `json:"files"`
}
