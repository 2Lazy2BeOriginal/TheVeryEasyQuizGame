package start

import (
	"fmt"
	"main/col"
	"math"
	"math/rand"
	"strings"
	"time"
)

// default start speed is 90 or 11 letters per second
var Speed time.Duration = 90

// I change the speed '90' in the undertext function to 'speed' so if you want to test something without display_speed function then set or remove speed to original speed (90)
func display_speed() {
	input := 0
	var text = "Is this fast enough for you?"
	fmt.Println("                                  ", col.Highlight("Adjust Speed")+"\n\n                      How fast do you want the text to go?\n\n                                 \n                                     75% [1]\n                                     50% [2]\n                                     25% [3]\n                                      0% [4]\n                                 Instant [5]\ntype number here (1, 2, 3, 4, 5)")

	// sample is only useful in this function
	sample := func() {
		Undertext(text)
		fmt.Println()
		if ask_again() {
			return
		}
	}

	for {
		fmt.Scanln(&input)
		if input == 5 {
			Speed = 0
			sample()
			break
		} else if input == 1 {
			Speed = 50
			sample()
			break
		} else if input == 2 {
			Speed = 90
			sample()
			break
		} else if input == 3 {
			Speed = 130
			sample()
			break
		} else if input == 4 {
			Speed = 150
			sample()
			break
		} else {
			fmt.Println("Type in the numbers that are shown above please")
		}
	}
}

func ask_again() bool {
	input := ""
	for {
		fmt.Println("Are you sure you want to start with this speed? (Y/N)")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "y" || input == "yes" {
			//move on to the next function/part
			Undertext("Continuing to the intro...")
			col.Clearscreen()
			//intro
			return true
		} else if input == "n" || input == "no" {
			col.Clearscreen()
			display_speed()
			return false
		} else {
			fmt.Println("Please enter Y (for yes) or N for no")
		}
	}
}

// menu screen can't escape exit
func Display_menu() {
	input := ""
	// mini function that tells the program to go to the actual game
	move_on := func() {
		Speed = 90
		//intro()
		col.Clearscreen()

	}
	// got an import loop and I was too lazy to work around it
	fmt.Println("                                ", col.Highlight("A really simple quiz"))
	fmt.Println("\n                           Are you smarter than a human?\n\n\n                                       Play  [P]\n\n                                      Speed  [S]\n\n                                       Exit  [E]")
	fmt.Println("\nWhat do you want to do? (P/S/E)")
	invalid := true
	for invalid {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "p" || input == "play" {
			invalid = false
			move_on()
		} else if input == "s" || input == "speed" {
			invalid = false
			col.Clearscreen()
			display_speed()
		} else if input == "e" || input == "exit" {
			invalid = false
			Undertext("Free will doesn't exist in this world,^ moving on to the actual quiz")
			fmt.Println()
			move_on()
		} else {
			fmt.Println(col.RED + "Enter in the letter that are shown (ex: p) or type out the word (ex: play)" + col.RESETCOLOR)
		}
	}
	return
}

// show text like undertale. ^ serve as dramatic points to pause
func Undertext(s string) {
	const r string = "^"
	for i := range s {
		char := s[i : i+1]
		if char == "//" {
			continue
		}
		//original speed was 90
		time.Sleep(Speed * time.Millisecond)
		// ^ acts a key character to do a dramatic pause
		if char == r {
			time.Sleep(2 * time.Second)
		} else {
			fmt.Print(char)
		}
	}
}

// p is the percentage of text is uppercase, l is how many letters in total, c is the index.
func Angerify(s string, wtf int) string {
	p := 0.0
	l := 0.0
	c := 0
	// if user gets 65 wtf points than they get basically all caps
	p = math.Min(1, float64(wtf)/float64(65))
	l = float64(len(s)) * p
	// not gonna implement a system checking if its empty or already uppercase since a random lowercase is c0ol
	for i := 0; i < int(l); i++ {
		rand.Seed(time.Now().UnixNano())
		c = rand.Intn(len(s))
		// so basically I am keeping everything the same except the randomly uppercased letter. Achieved chunking off the "lucky" char and then piecing the rest back
		s = s[0:c] + strings.ToUpper(s[c:c+1]) + s[c+1:]
	}
	return s
}
