# Player Interface

Interface used to create a new player

The interface must implement the Update and Create methods.

## Basic Setup
Create a branch off main before beginning.

Create a file with a unqiue name and place it in the `players` directory.  At the top of the file, make sure you have the correct package listed: `package players`.

All players will need to implement the PlayerInterface and create a struct.

### Update
This method is called on every 'tick' (cycle of the game).  This allows the user to control movement of the tank as well as call any actions.  

Two variables are supplied to the user.  First being the tank itself that is being controlled.  The tank has a number of functions that facilitate movement.  Additionally the user is supplied with the positions of other tanks on the field.

At the very end of the Update function, the tank object must call the Update method it has.  This will ensure your updates are registered.
``` golang
func (sampleTank *SampleTank) Update(t *tank.Tank, pos []tank.TankPosition) {

	t.Update() // This always needs to be called within the Update method
}
```

### Create
The create method must return a tank object with the name and color fields filled out.  An example can be found below.
``` golang
	return &tank.Tank{
		Name:  "SampleTank",
		Color: colors.RED,
	}
```
### Struct
A struct, ideally with the name of your tank, should be created as well within the same fill as your implementation of the interface.  An example can be found below:
``` golang
type SampleTank struct {}
```

## API Basics
### **Tank**
The tank has a number of functions that allow the user to move or perform actions.

**GetPosition()**
    - Returns a TankPosition object

- **GetScore()**
    - Returns current score as an integer

- **GetHealth()**
    - Returns current health as an integer

- **Accelerate()**
    - Increases the speed of the tank.  There is a maximum speed that the tank can reach and will stop accelerating.

- **Deaccelerate()**
    - Decreases the speed of the tank.  There is a maximum speed that the tank can reach and will stop decelerating.

- **RotateLeft()**
    - Rotates the tank left.

- **RotateRight()**
    - Rotates the tank right.

- **Fire()**
    - Fires a bullet in the direction the tank is facing.  The tank can only have one bullet on the screen at a time.

- **Stop()**
    - Sets the tanks speed to zero immediately.

***NOTE***

There are a number of other functions available to the user that should not be used.  These functions are being used to control and keep track of the tank.  

### **Tank Position**
This is a struct that contains the following information within
``` golang
X: Tank X position
Y:  Tank Y position
Width: Tank Width
Height: Tank Height
BulletX: Bullet X position
BulletY: Bullet Y position
BulletSize: Bullet Size (not needed by players)
Facing: Direction facing in Radians (Is a running total, can be pos and negative)
```

## Rules
Tanks begin in preestablished spots in the corners of the map.  They all start with 3 health and a score of 0.  You gain score by shooting and hitting another tank on the map.  Your tank is removed from the map when your reach a health of 0.

### Collisions
If your tank runs off the map, collides with another tank or collides with an object on the map - you will lose one point of health.  Additionally, if a bullet collides with your tank, you will lose one point of health.

After colliding with or being shot, you will respawn in your original position.  There is a short period of immunity, but that is not very long.  

### Win Conditions
The tank that is the last one standing wins the event.

### Losing
You lose by having 0 health.


