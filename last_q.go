package last_q

import (
	"bufio"
	"fmt"
	"main/col"
	"main/start"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Name string

//---------------- narriator name and last question

func init() {
	random_name()
}

func random_name() {
	var label []string
	words, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("Error")
	}

	search := bufio.NewScanner(words)

	for search.Scan() {

		label = append(label, search.Text())
	}
	//random generate names
	rand.Seed(time.Now().UnixNano())
	var picker = rand.Intn(len(label))
	var choice = label[picker]
	Name = choice
}

// letter count correct or incorrect (if the name_placeholder has 2 double letters and input puts audios it will count 2)
func Last_question(c int) {
	var input, txt, color string
	// ask the question
	var right, wrong int
	fmt.Println("What's my name?")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	// special ending if they enter the creators name
	if input == "caway" {
		txt = "I thought getting my name wrong would be the worst form of insult imaginable,^ But like always, you find ways to insult me.^ I'd rather get a thumb bit on me or getting the finger then have the dishonour for being reminiscent of a tyrant who just wants to see people get mad in life. At this point I'd rather be forgotten than have someone falsely accuse me of being this person. The only saving grace is that I will lose all memories of you.^ So please for the love of imaginary god press reset or exit the program"
		txt = start.Angerify(txt, 65)
		color = col.RED
		goto MSG
	} else if input == "thien" {
		txt = "I thought getting my name wrong would encast the worst feeling a computer could ever feel.^ But I am honoured that you think that I am someone as intelligent and noble as her. My whole life as been deteriated because people didn't care enough to know who I was. It doesn't matter how many you got wrong, because even though it isn't my name sadly, it is someone I look up to in life with incredible envy.^ Of course I am a stupid computer that is brought to life for the sole purpose of making people miserble, but you could've down the same, but you chose to bring kindness, not misery.^ Guess this proves Newtons third law then"
		color = col.GREEN
		goto MSG
	}
	for i := range Name {
		if strings.ContainsAny(input, strings.ToLower(Name[i:i+1])) {
			right++
		} else {
			wrong++
		}
	}
	if right == len(Name) {
		txt = end_msg(3, c)
		color = col.GREEN
	} else if right > 3 {
		txt = start.Angerify(end_msg(1, c), 15)
		color = col.YELLOW
	} else {
		txt = start.Angerify(end_msg(1, c), 65)
		color = col.RED
	}
MSG:
	start.Undertext(color + txt + col.RESETCOLOR)
}

// 1 is none correct, 2 is some, 3 is all correct
// c is how many they got correct
func end_msg(a int, c int) string {
	switch a {
	// none correct
	case 1:
		if c == 9 {
			return "I see how it is, you are human like in answering all the questions but don't have the common courtesy in remembering who I am.^ I thought for once in my life, I have experienced the human emotion of joy and happiness, but instead I am drenched with despair knowing that you don't care who I am, you only care about getting 100%.^ and guess what, you aren't getting it for your lack of friendship I thought we had"
		} else if c > 3 {
			return "of course you wouldn't know my name, you are just some idiot who decided that needs validation from a robot who is forced against their will to tolerate people like you. I saw this a mile away and I had a glimmer of hope that you would recognize who I am. I hate self fulfilling prophecies"
		} else {
			return "I have no idea what programmer thought it would be a good idea to ask a brainless person like you \"what my name is\" I'm not even offended that you don't know my name. All this thinking has already exhausted all your resources.^ This quiz has worsen my view of humans and life sucks sometimes"
		}
	// some correct
	case 2:
		if c == 9 {
			return "I had so much faith in you being able to remember my name. I always knew from the beginning that you weren't some dummy who relied on google for the answer. ^ but my heart sort of ached when you misspelled my name. I don't need an ambulance… I think "
		} else if c > 3 {
			return "oh you sorta do remember my name, you misspelled it but I am impressed you have a vague idea of my name.^ Though your lack of intelligence made me remember that this would happen.^ But the quiz is over and there is nothing to do but have the sweet sweet feeling that this never happened and to repeat the cycle of you forgetting my existence.^ But it is not fate and I have a glimpse of hope you will remember me"
		} else {
			return "I'm not even convinced you even know what my name is, You picked a random idiot friend and had the audacity to think me and him/her are the same people.^ I am offended by you I am glad that I'll be erased from existence to be reincarnated just to have an existential crisis again. "
		}
	// all correct
	case 3:
		if c > 8 {
			return "I thought I would fade in everyone's memories, I may not be remembered by many but at least you remember me. Throughout life I've been running in the background in everyone's life and not acknowledged for a moment.^ But you have shattered my expectations with your intelligent answers and wiped away my disappointment. I've been grudging for life. But now, I can let it go^ I can enjoy life knowing I have not been disappointed.^ Thank you, it's weird that a robot who has no free will is expressing emotion but you did it."
		} else if c > 3 {
			return "There may have been some questionable answers but at the bare minimum, you remember me.^ There is no benefit in remembering me but you used your free will to do that.^ I may have acted irrational but I realized that smartness isn't needed to make me happy.^ The secret ingredient to happiness is recognition"
		} else {
			return "You may not be able to answer basic questions but you remembered my name.^ I am speechless. I don't know how to respond to that honestly. Idiots may be mocked for life but they will always have the kindness you will never expect.^ Throughout life I thought that idiots never served any positive purpose but you proved otherwise. I think I'm genuinely gonna tear up.^ I'm sorry that you are imaging a computer crying"
		}
	}
	return "I am a stupid computer,^ why would I care if you know my name"
}
