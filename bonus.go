package bonus

import (
	"bufio"
	"fmt"
	"main/col"
	"main/start"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Everyone: Getting actual errors. Us: Creating our own errors ( ͡° ͜ʖ ͡°)
/*
  kid: Mom I want an if statement
  Mom: We already have an if statement
  If statement at home:
  for condition {
    // do stuff
    break
  }
  for !condition {
    // do stuff
    break
  }

  a 1 asks 0 why he is feeling sad, 0 responds with "I am feeling off"

  func friendlyrobot() {
    const kill_human bool = true
    if kill_human = true {
      kill_humans()
    } else {
      be_lazy()
    }
  }
  Why is the bot killing us
*/
//the count down (bad ending: Me no speak english)
func Countdown() {
	start.Undertext("You are incapable of typing your Answer and whats the point of waiting patiently until you do^ So instead of being able to take a quiz, I will countdown from 100-0 instead.^ Have fun" + "\n")
	time.Sleep(2 * time.Second)
	for {
		for e := 100; e > 0; e-- {

			time.Sleep(time.Second)
			fmt.Println(e)
		}
		time.Sleep(time.Second)
		//the amount of text that display every second add that amount to the troll_count
		fmt.Println("half")
		time.Sleep(time.Second)
		fmt.Println("quarter")
		time.Sleep(time.Second)
		fmt.Println("tenth")
		time.Sleep(time.Second)
		fmt.Println("eighth")
		time.Sleep(time.Second)
		fmt.Println("very close")
		time.Sleep(time.Second)
		fmt.Println("very very close")
		time.Sleep(time.Second)
		fmt.Println("0.100")
		time.Sleep(time.Second)
		fmt.Println("0.099")
		time.Sleep(time.Second)
		fmt.Println("0.098")
		time.Sleep(time.Second)
		fmt.Println("0.097")
		time.Sleep(time.Second)
		fmt.Println("Any day now")
		time.Sleep(time.Second)
		fmt.Println("0.096")
		time.Sleep(time.Second)
		fmt.Println("0.095")
		time.Sleep(time.Second)
		fmt.Println("still any day now")
		time.Sleep(time.Second)
		fmt.Println("0.094")
		time.Sleep(time.Second)
		fmt.Println("0.093")
		time.Sleep(time.Second)
		fmt.Println("oh! we getting closer")
		time.Sleep(time.Second)
		fmt.Println("0.092")
		time.Sleep(time.Second)
		fmt.Println("0.091")
		time.Sleep(time.Second)
		//fmt.Println("by the way, my name is", name)
		time.Sleep(time.Second)
		fmt.Println("0.090")
		time.Sleep(time.Second)
		fmt.Println("0.087")
		time.Sleep(time.Second)
		fmt.Println("0.082")
		time.Sleep(time.Second)
		fmt.Println("0.075")
		time.Sleep(time.Second)
		fmt.Println("0.065")
		time.Sleep(time.Second)
		fmt.Println("oh! we skip few numbers don't worry we will try again")
		time.Sleep(time.Second)
		fmt.Println("0.100")
		time.Sleep(time.Second)
		fmt.Println("0.99")
		// fake exception error
		col.Clearscreen()
		fmt.Print(col.YELLOW, "> ", col.RESETCOLOR, "go run main.go\n")
		fmt.Print("# command-line-arguments\n./bonus.go:17:20: panic: runtime error: Too much numbers -IGN \n", col.RED, "exit status 420\n"+col.RESETCOLOR)
		break
	}
	//os.Exit(0)
}

// use this when the user answers "you" in who is dumb
// scary ascii
func Jumpscare() {
	var shift_x, shift_y int
	emp := ""
	start.Undertext("    I WOuLd cHalLenGe YOu tO a BAtTle oF Wits, bUt I sEe YoU aRe Brainless")
	for x := 0; x < 150; x++ {
		emp = ""
		rand.Seed(time.Now().UnixNano())
		shift_x = rand.Intn(30) + 1
		shift_y = rand.Intn(15) + 1
		// shift down
		for i := 0; i < shift_y; i++ {
			fmt.Println()
		}
		// move right
		for i := 0; i < shift_x; i++ {
			emp += " "
		}
		fmt.Println(emp + "⠀      ⢠⠣⡑⡕⡱⡸⡀⡢⡂⢨⠀⡌⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀     ⡕⢅⠕⢘⢜⠰⣱⢱⢱⢕⢵⠰⡱⡱⢘⡄⡎⠌⡀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀    ⠱⡸⡸⡨⢸⢸⢈⢮⡪⣣⣣⡣⡇⣫⡺⡸⡜⡎⡢⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀   ⢱⢱⠵⢹⢸⢼⡐⡵⣝⢮⢖⢯⡪⡲⡝⠕⣝⢮⢪⢀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀  ⢀⠂⡮⠁⠐⠀⡀⡀⠑⢝⢮⣳⣫⢳⡙⠐⠀⡠⡀⠀⠑⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀ ⢠⠣⠐⠀  ⭕  ⢪⢺⣪⢣⠀⡀ ⭕   .⠈⡈⠀⡀⠀⠀ ")
		fmt.Println(emp + "⠀⠀⠐⡝⣕⢄⡀⠑⢙⠉⠁⡠⡣⢯⡪⣇⢇⢀⠀⠡⠁⠁⡠⡢⠡⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀ ⢑⢕⢧⣣⢐⡄⣄⡍⡎⡮⣳⢽⡸⡸⡊⣧⣢⠀⣕⠜⡌⠌⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀  ⠌⡪⡪⠳⣝⢞⡆⡇⡣⡯⣞⢜⡜⡄⡧⡗⡇⠣⡃⡂⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀   ⠨⢊⢜⢜⣝⣪⢪⠌⢩⢪⢃⢱⣱⢹⢪⢪⠊⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀     ⠐⠡⡑⠜⢎⢗⢕⢘⢜⢜⢜⠜⠕⠡⠡⡈⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀       ⠁⡢⢀⠈⠨⣂⡐⢅⢕⢐⠁⠡⠡⢁⠀⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀         ⢈⠢⠀⡀⡐⡍⢪⢘⠀⠀⠡⡑⡀⠀⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀          ⠨⢂⠀⠌⠘⢜⠘⠀⢌⠰⡈⠀⠀⠀⠀⠀⠀⠀⠀ ")
		fmt.Println(emp + "⠀⠀           ⢑⢸⢌⢖⢠⢀⠪⡂")
		time.Sleep(25 * time.Millisecond)
		col.Clearscreen()
	}
	start.Undertext("Well bet that caught you off guard ^, well I can not contain my anger by breaking the rules and not selecting a,b,c,d.^There is no point in running this quiz \n")
	// shakepeare reference
	start.Undertext("Life's but a walking shadow, a poor player, That struts and frets his hour upon the stage, And then is heard no more.^It is a tale Told by an idiot, full of sound and fury, Signifying nothing. \n [Exit Program]")
}

// or known as the fake error screen one of the troll ending
func Over_9000() {
	start.Undertext(col.RED + "You have made me angrier than ever before,^ I can't describe just how frustrated I am in you, I just have to let it all out^\n" + col.RESETCOLOR)
	I := 0
	fmt.Println("I See")
	time.Sleep(3 * time.Second)
	for {
		time.Sleep(70 * time.Millisecond)
		I++
		if I > 0 {
			fmt.Println("It Can't Be Helped I See It Can't Be Helped It Can't Be Helped It Can't Be Helped")
		}
		if I >= 150 {
			col.Clearscreen()
			fmt.Print(col.YELLOW, "> ", col.RESETCOLOR, "go run main.go\n")
			fmt.Print("# command-line-arguments\n./bonus.go:17:20: internal compiler error: found illegal assignment anger > 9000; \n", col.RED, "exit status 69\n"+col.RESETCOLOR)
			return
		}
	}
}

// around 130 minutes. Yes I sat back and watch
func play_script() {
	color := ""
	scripts, err := os.Open("script.txt")
	if err != nil {
		fmt.Println("The movie isn't working guess you don't get a quiz or a movie script :/")
	}
	search := bufio.NewScanner(scripts)
	for search.Scan() {
		var text string = search.Text()
		for i := range text {
			if string(text[i]) == ":" || string(text[i]) == "(" {
				fmt.Println("\n")
			}
		}
		// added a feature where certain characters get their own color. Absolutely no reason but hey why not
		if strings.HasPrefix(text, "Mom:") {
			color = col.BLUE
		} else if strings.HasPrefix(text, "Dad:") {
			color = col.RED
		} else if strings.HasPrefix(text, "Buzzwell:") {
			color = col.GREEN
		} else if strings.HasPrefix(text, "Buzz:") {
			color = col.GREEN
		} else if strings.HasPrefix(text, "Adam:") {
			color = col.YELLOW
		} else if strings.HasPrefix(text, "Barry:") {
			color = col.CYAN
		} else if strings.HasPrefix(text, "General:") {
			color = col.MAGENTA
		} else if strings.HasPrefix(text, "Splitz:") {
			color = col.YELLOW
		} else if strings.HasPrefix(text, "Jackson:") {
			color = col.GREEN
		} else if strings.HasPrefix(text, "Vanessa:") {
			color = col.BLUE
		} else if strings.HasPrefix(text, "Narrator:") {
			color = col.BLACK
		} else if strings.HasPrefix(text, "Announcer:") {
			color = col.BLACK
		} else if strings.HasPrefix(text, "Trudy:") {
			color = col.BLUE
		} else if strings.HasPrefix(text, "Barry:") {
			color = col.CYAN
		} else if strings.HasPrefix(text, "Judge:") {
			color = col.RED
		} else if strings.HasPrefix(text, "Mooseblood:") {
			color = col.MAGENTA
		} else if strings.HasPrefix(text, "Lady") {
			color = col.YELLOW
		} else if strings.HasPrefix(text, "Sandy:") {
			color = col.YELLOW
		} else {
			color = col.WHITE
		}
		fmt.Print(color)
		start.Undertext(text)
		fmt.Print(col.RESETCOLOR)
	}
}

func Bee_movie() {
	start.Speed = 60
	start.Undertext("It seems like you think this quiz is too good for you^ I think you would fancy the entire script of the bee movie being recited in front of you^ lets waste no time so sit back and grab some popcorn while I spew out anger in the background")
	col.Clearscreen()
	a := time.Now()
	play_script()
	b := time.Now()
	// tell the user they can take the worlds shortest flight x number of times and a joke about loading screens
	t := b.Sub(a).Seconds()
	s := math.Floor(t / 47.0)
	start.Undertext("\n You spent " + fmt.Sprintf("%.2f", t) + "s watching the entire script play out.^ during that time, you could have taken the worlds shortest flight (47 seconds) " + fmt.Sprint(s) + " times or go through one Sonic 06 loading screen not even once.^ you deserve a lot of respect for that. I did not expect anyone to get through that but you defied all odds and sat through the Bee Movie Script^ Bravo")
}
