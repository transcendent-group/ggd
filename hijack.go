package main

import (
    "time"

    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {

    time.Sleep(10*time.Second)

    player2 := tello.NewDriver("8888")

    work := func() {
        gobot.After(5*time.Second, func() {
            player2.TakeOff()
        })
    }

    robot := gobot.NewRobot("tello",
        []gobot.Connection{},
        []gobot.Device{player2},
        work,
    )

    robot.Start()
}