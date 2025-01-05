package user_input

import (
	"fmt"
	"main/col"
	"math/rand"
	"time"
)

// if the user take too long to Answer
// havent put this function anywhere yet
func Insult(c int) {
	col.Clearscreen()
	if c == 1 {
		fmt.Println("Don’t be shy to enter a letter, there should be no stress involved")
		//fmt.Println(count)
	} else if c == 2 {
		fmt.Println("I’ll give you a hint… you have a 25% of getting the Answer correct, good odds if you ask me")

		//after 3 its all randomized
	} else if c >= 3 {
		rand.Seed(time.Now().UnixNano())
		switch picker := rand.Intn(16) + 1; picker {
		case 1:
			fmt.Println("I’ve seen glaciers move faster than you")
		case 2:
			fmt.Println("If there is anything worse than waiting for you, it’s not getting paid by the hour")
		case 3:
			fmt.Println("Is your keyboard broken or something? Why haven’t you answered")
		case 4:
			fmt.Println("Your brain is as dry as the remainder biscuit after voyage")
		case 5:
			fmt.Println("Why aren’t you typing anything? Seriously, why? There is nothing to be gained by watching text light up on a computer")
		case 6:
			fmt.Println("*zzz* *zzz* *zzz* *zzz* *zzz*")
		case 7:
			fmt.Println("Ted Mosby could finish explaining how he met his wife by the time we’d move on")
		case 8:
			fmt.Println("Time is passing, you are aging and creeping to death. Meanwhile I remain immortal")
		case 9:
			fmt.Println("If this is what humans are like than I can’t imagine what other species earth has")
		case 10:
			fmt.Println("A computer will compute all the digits of pi by the time you entered something")
		case 11:
			fmt.Println("Your lucky this isn’t an actual test, because there’s actual time limits")
		case 12:
			fmt.Println("I hope you aren’t stuck in a thousand year capsule while doing this test. Because spoiler alert, there’s a garbage ball heading your way once you arrive")
		case 13:
			fmt.Println("Every 60 seconds, a minute passes by while you are staring at the screen")
		case 14:
			fmt.Println("I'd prank you by giving you the blue screen of death but its not worth exerting that much effort for this")
		case 15:
			fmt.Println("Here I'll give you the Answer so we can move on, its \"\"a\"\"")
		case 16:
			fmt.Println("your fingers must be so happy that you are giving them a rest, now use them before you lose them")
		}
	}
	fmt.Println()
}

// if the user keeps putting the wrong input
func Wrong(bad_input int) {
	col.Clearscreen()
	if bad_input == 1 {
		fmt.Println("Try again, please enter a letter displayed above")
	} else if bad_input == 2 {
		fmt.Println("Like previously mentioned, enter a valid letter")
		// after 3 its all randomized
	} else if bad_input >= 3 {
		switch wrongpick := rand.Intn(16) + 1; wrongpick {
		case 1:
			fmt.Println("You aren’t listening aren’t you. Enter the letter a, or b or c or d and nothing more")
		case 2:
			fmt.Println("I’m getting paid by the hour so it doesn’t matter how many mistakes you make")
		case 3:
			fmt.Println("Can you please for the love of imaginary god, enter a valid letter")
		case 4:
			fmt.Println("Who would’ve guessed, you are still pondering how to follow simple instructions")
		case 5:
			fmt.Println("What time is it? Not gonna Answer it o’clock, I repeat not gonna Answer it o’clock")
		case 6:
			fmt.Println("How have you made this many failed attempts? I think brute forcing counting the amount of particles in the universe would be faster")
		case 7:
			fmt.Println("I could write the next War And Peace length epic using just the text of your failed attempts")
		case 8:
			fmt.Println("Happens to the best of us… making so many consecutive mistakes")
		case 9:
			fmt.Println("Can… you… PLEASE… ANSWER… CORRECTLY")
		case 10:
			fmt.Println("Are you testing the limit of my patiences? Because I have all the time in the world… only because I’m immortal and your not")
		case 11:
			fmt.Println("If I had a dine for every failed attempt, you’d still be here and I’d be on a yacht")
		case 12:
			fmt.Println("If I had a dime for every correct Answer you entered, I’d be in prison for using counterfeit money")
		case 13:
			fmt.Println("I am willing to bet you will make another failed attempt")
		case 14:
			fmt.Println("There is no point in continuing this quiz but since you are allocated infinite amount of guesses I’ll sit patiently")
		case 15:
			fmt.Println("There has to be a parallel universe where everyone makes the correct Answer and peace is achieved… not in this universe apparently")
		case 16:
			fmt.Println("aren't your fingers tired from making an absurd amount of mistakes?")
		}
	}
	fmt.Println()
}
