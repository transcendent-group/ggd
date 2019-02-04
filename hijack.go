// Hijacks a drone mid-air, performs some acrobatics, and lands

package main

import (
    "time"

    "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
    drone := tello.NewDriver("8888")

    drone.Start()

    time.Sleep(2*time.Second)
    
    drone.LeftFlip()

    time.Sleep(5*time.Second)
    
    drone.RightFlip()

    time.Sleep(5*time.Second)
    
    drone.Land()
}
