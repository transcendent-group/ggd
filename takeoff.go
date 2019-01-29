package main

import (
    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
    drone := tello.NewDriver("8888")

    work := func() {
        drone.TakeOff()
    }

    robot := gobot.NewRobot("tello",
        []gobot.Connection{},
        []gobot.Device{drone},
        work,
    )

    robot.Start(false)
}