package main

import (
	"encoding/json"
	"errors"
	"fmt"
	// 	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func session