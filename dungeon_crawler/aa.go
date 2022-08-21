package main

const (
	all = `
L       /
#L     /#
#|L _ /|#
#|#|#|#|#
#|#|_|#|#
#|/   L|#
#/     L#
/       L`

	frontLeftNorth = `
         
         
  _      
 |#|     
 |_|     
         
         
         `

	frontRightNorth = `
         
         
      _  
     |#| 
     |_| 
         
         
         `

	frontNorth = `
         
         
    _    
   |#|   
   |_|   
         
         
         `

	frontWest = `
         
         
 |L      
 |#|     
 |#|     
 |/      
         
         `

	frontEast = `
         
         
      /| 
     |#| 
     |#| 
      L| 
         
         `

	leftNorth = `
         
_        
#|       
#|       
#|       
_|       
         
         `

	rightNorth = `
         
        _
       |#
       |#
       |#
       |_
         
         `

	north = `
         
  _____  
 |#####| 
 |#####| 
 |#####| 
 |_____| 
         
         `

	west = `
L        
#L       
#|       
#|       
#|       
#|       
#/       
/        `

	east = `
        /
       /#
       |#
       |#
       |#
       |#
       L#
        L`
)

type location = int

const (
	locationFrontLeft location = iota
	locationFrontRight
	locationFront
	locationLeft
	locationRight
	locationCenter
	locationMax
)

var aaTable = map[location][directionMax]string{
	locationFrontLeft: {
		frontLeftNorth, "", "", "",
	},
	locationFrontRight: {
		frontRightNorth, "", "", "",
	},
	locationFront: {
		frontNorth, frontWest, "", frontEast,
	},
	locationLeft: {
		leftNorth, "", "", "",
	},
	locationRight: {
		rightNorth, "", "", "",
	},
	locationCenter: {
		north, west, "", east,
	},
}

var locations = map[direction][locationMax]vec2{
	directionNorth: {
		{x: -1, y: -1},
		{x: 1, y: -1},
		{x: 0, y: -1},
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: 0},
	},
	directionWest: {
		{x: -1, y: 1},
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 0, y: 0},
	},
	directionSouth: {
		{x: 1, y: 1},
		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: -1, y: 0},
		{x: 0, y: 0},
	},
	directionEast: {
		{x: 1, y: -1},
		{x: 1, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 0, y: 0},
	},
}
