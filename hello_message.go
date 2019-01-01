// hello_message.go

package main

import (
    "myitcv.io/react"
)

//go:generate reactGen

// Step 1
// Declare a type that has (at least) an anonymous embedded react.ComponentDef
// (it can have other fields); this type must have the suffix 'Def', which corresponds to
// 'Definition'
//
type HelloMessageDef struct {
    react.ComponentDef
}

// Step 2
// Optionally declare a props type; the naming convention is *Props
//
type HelloMessageProps struct {
    Name string
}

// Step 3
// Optionally declare a state type; the naming convention is *State
//
type HelloMessageState struct {
    count int
}

// hello_message.go continued....

// Step 4
// Declare a function to create instances of the component, i.e. an element. If
// your component requires props to be specified, add this to the function
// signature. If the props are optional, use a props pointer type.
//
// buildHelloMessageElem is code generated to wrap a call to react.CreateElement.
//
// Convention is that this function is given the name of the component, HelloMessage
// in this instance. Because this component has props, we also accept these as part
// of the constructor.
//
func HelloMessage(p HelloMessageProps) *HelloMessageElem {
    return buildHelloMessageElem(p)
}

// Step 5
// Define a Render method on the component's non-pointer type
//
func (r HelloMessageDef) Render() react.Element {
    return react.Div(nil,
        react.S("Hello "+r.Props().Name),
    )
}