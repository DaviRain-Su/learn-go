package main

import (
	"fmt"
	"headfirstgo/src/chapter11"
)

func playList(device chapter11.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called!")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func DummyPlay(n NoiseMaker) {
	n.MakeSound()
}

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	var player chapter11.Player = chapter11.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player = chapter11.TapeRecorder{}
	playList(player, mixtape)

	var value MyInterface

	value = MyType(4)

	value.MethodWithoutParameters()
	value.MethodWithParameter(3)
	fmt.Println(value.MethodWithReturnValue())
	//fmt.Println("hello, world!")

	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	DummyPlay(toy)
	//toy.MakeSound()
	toy = Horn("Toyco Blaster")
	//toy.MakeSound()
	DummyPlay(toy)

	s := Switch("off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()

	var noiseMaker NoiseMaker = Robot("Botco Amber")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()

	var palyer chapter11.Player = chapter11.TapePlayer{}
	fmt.Println("palyer is ", palyer)
	recorder, ok := player.(chapter11.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapRecorder")
	}

	var err error
	err = ComedyError("What's a programmer's favoite beer? Logger!")
	fmt.Println(err)

	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func TryOut(palyer chapter11.Player) {
	palyer.Play("Test Track")
	palyer.Stop()
	recorder, ok := palyer.(chapter11.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
