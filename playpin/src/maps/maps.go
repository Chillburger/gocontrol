package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}


// a map maps keys to values
// a zero value map is nil, nil map has no keys nor can keys be added

var m map[string]Vertex

// map literal are like struct literals but keys are required
var mapLiteral = map[string]Vertex {
  "Bell Labs": Vertex {
    40.68433, -7439967,
  },
  "Google": Vertex {
    37.42202, -122.08408,
  },
}

// if top level type is just a type name, you can ommit it from elements of the literals
var mapOmit = map[string] Vertex {
  "Bell Labs": { 40.94848, -74.994949},
  "Google": {37.42224, -122.889},
}

func main() {
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex { 40.68433, -74.39967}

  fmt.Println(m["Bell Labs"])

  // print the map literal above
  fmt.Println(mapLiteral)

  // print the ommited types
  fmt.Println(mapOmit)


  // mutating maps
  m := make(map[string]int)
  fmt.Println("The value:", m["Answer"])
  // insert or update an element of a map
  m["Answer"] = 48
  // retrieve an element
  fmt.Println("The value:", m["Answer"])

  // delete an element in a map
  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  // test that a key is present with a two value assignment
  v, ok := m["Answer"]
  fmt.Println("The Value:", v, "Present?", ok)
}
