package ast

import (
    "foo/token"
)

type Attrib interface {}

type World struct {
    Name string
}

func NewWorld(id Attrib) (*World, error) {
    return &World{string(id.(*token.Token).Lit)}, nil
}

func (this *World) String() string {
    return "hello " + this.Name
}