package model

import (
	"github.com/jiangrx816/gopkg/db"
	"gorm.io/gorm"
)

func DefaultPicture() *gorm.DB { return db.MustGet("picture") }

func DefaultPoetry() *gorm.DB { return db.MustGet("poetry") }

func DefaultMarket() *gorm.DB { return db.MustGet("market") }

func DefaultStory() *gorm.DB { return db.MustGet("story") }
