package main

import (
	"fmt"
	// external files I made
	"main/bonus"
	"main/col"
	"main/last_q"
	"main/start"
	"main/user_input"
	"math/rand"
	"os"
	"strings"
	"time"
)

// skip to line 400 to get out of the madness struct thing
// skip to 674 for actual code

type question struct {
	q      string
	ans    string
	letter string
	a      Answer
	b      Answer
	c      Answer
	d      Answer
}

type Answer struct {
	ans      string
	dumb_val int
	msg      string
}

var oneq = question{
	q:      "Are these questions too hard?",
	letter: "r",
	a:      Answer{"Me no speak english", 0, "oh so you have always taken random guesses to get the Answer then^ still don't get how you didn't know what 2+3 is.^ But wheres the fun in getting everything wrong? From now on, the quiz only features correct answers"},
	b:      Answer{"My brain is poopy right now", 0, "considering you thought poopy is an amusing word, I think the difficulty is too hard. Let me make make a quiz catered to you where no answers are wrong"},
	c:      Answer{"Yes", 0, "well good to know you are honest, and to reward you, the quiz is going to be 100% correct^ I'm upping therapy 6 times a week now. "},
	d:      Answer{"No", 0, "No???? NO???? How are you so narcissistic to not see your inability to Answer basic questions^ Sorry you might have been honest but I don't care enough to investigate so here is a quiz where no answers are wrong so you can have the fuzzy feeling of being correct"},
}

var onea = question{
	q:      "Who has cooties?",
	letter: "g",
	a:      Answer{"Not Me", 0, "Well one of the symptoms of cooties is extreme denial. I am confident that you have cooties and are afraid of the stigma related to cooties.^ Don't worry, I won't secretly screenshot and use the webcam to blackmail you."},
	b:      Answer{"The President", 0, "It is a common misconception that world leaders have cooties and therefore are responsible for all the glory or devastation in the world.^ However cooties are not the driving force. Cooties is responsible for falsely accusing others of having Cooties."},
	c:      Answer{"Elementary Kids", 0, "It is true that the origin of cooties came into fruition from a local Elementary Kids when one of the students snitched too much but currently the Kids are immune to cooties.^ Now you on the other hand…"},
	d:      Answer{"Sonic The Hedgehog", 0, "According to Wikipedia, Cooties is not transferable to Hedgehogs so therefore Sonic The Hedgehog is unable to get Cooties.^ now you are a human so you may or may not have it and are in denial."},
}

var oneb = question{
	q:      "What would Rick Astley never do",
	letter: "g",
	a:      Answer{"Give you up", 0, "Rick would never do any of these inhumane acts but what he does is pop up out of nowhere when you least expect it.^ Kinda jealous if I'm being honest"},
	b:      Answer{"Make you cry", 0, "Rick would never do any of these inhumane acts but what he does is pop up out of nowhere when you least expect it.^ Kinda jealous if I'm being honest"},
	c:      Answer{"Say goodbye", 0, "Rick would never do any of these inhumane acts but what he does is pop up out of nowhere when you least expect it.^ Kinda jealous if I'm being honest"},
	d:      Answer{"Let you down", 0, "Rick would never do any of these inhumane acts but what he does is pop up out of nowhere when you least expect it.^ Kinda jealous if I'm being honest"},
}

var onec = question{
	q:   "What platform do you like playing Video games on",
	ans: "d",
	a:   Answer{"Sony", 0, "No judgment, Playstation is the only console series to never have a failed system and it's no wonder why. Endless supply of games and exclusives.^ But you didn't pick Nintendo and I'm sorta disappointed at that"},
	b:   Answer{"Microsoft", 0, "Xbox is a superb system. It has an Xbox game pass which allows you to play any game for a low low price—If you don't mind paying a fortune for the internet. But if basic math has taught me.^ Xbox is not Nintendo so my heart kinda aches a bit"},
	c:   Answer{"PC", 0, "Oh lord you are one of those people huh. PC become outdated with every continuous breath. Like once I'm done complaining, there is going to be a new PC build that makes every older PC obscure^ PC gamers are only slightly better than mobile gamers"},
	d:   Answer{"Nintendo", 0, "I can tell that we are going to be instant friends because nothing beats Nintendo, Nintendo beats Rock, Paper, Scissors in Rock Paper Scissors. It's mightier than the sword, it's more worthy than god.^ If only I wasn't made on a Microsoft based system though"},
}

var oned = question{
	q:      "If you had to only save one family member who would you save",
	letter: "g",
	a:      Answer{"Mom", 0, "Well this gives us some insight of which of the 4 you have a stronger relationship. I for one am lonely for life and have no emotions embedded in me.^ Well this has been an interesting Journey. I appreciate that you admitted which family member is better and trust me. I know there's no way out of this one and you made the difficult decision to pick one. Because you don't want this to be carelessly randomly picked.^ This has been a very interesting quiz where I veered off course and asked a personal question. And you were a good sport for answering.^ Thank You."},
	b:      Answer{"Dad", 0, "Well this gives us some insight of which of the 4 you have a stronger relationship. I for one am lonely for life and have no emotions embedded in me.^ Well this has been an interesting Journey. I appreciate that you admitted which family member is better and trust me. I know there's no way out of this one and you made the difficult decision to pick one. Because you don't want this to be carelessly randomly picked.^ This has been a very interesting quiz where I veered off course and asked a personal question. And you were a good sport for answering.^ Thank You."},
	c:      Answer{"Sibling", 0, "Well this gives us some insight of which of the 4 you have a stronger relationship. I for one am lonely for life and have no emotions embedded in me.^ Well this has been an interesting Journey. I appreciate that you admitted which family member is better and trust me. I know there's no way out of this one and you made the difficult decision to pick one. Because you don't want this to be carelessly randomly picked.^ This has been a very interesting quiz where I veered off course and asked a personal question. And you were a good sport for answering.^ Thank You."},
	d:      Answer{"Family Pet", 0, "Well this gives us some insight of which of the 4 you have a stronger relationship. I for one am lonely for life and have no emotions embedded in me.^ Well this has been an interesting Journey. I appreciate that you admitted which family member is better and trust me. I know there's no way out of this one and you made the difficult decision to pick one. Because you don't want this to be carelessly randomly picked.^ This has been a very interesting quiz where I veered off course and asked a personal question. And you were a good sport for answering.^ Thank You."},
}

var twoq = question{
	q:      "Are you stupid?",
	letter: "r",
	a:      Answer{"yes", 0, "well good to know you are honest,^ I'm going need to preorder my anger pills after this stupid quiz is over but this quiz is not for everyone, so here's a modified quiz just for you"},
	b:      Answer{"Dumb enough to click this", 3, "If you are dumb enough to click without much thought,^ Then there is a lot of potential to fool you to Answer the most illogical answers possible^ Lets try that out, I will change up the questions to test my theory^"},
	c:      Answer{"No", 0, "Really^ no?... no?... I would press X to doubt but I am too lazy to do so. Either way you will be punished for your dishonesty"},
	d:      Answer{"I failed homeroom, of course", 2, "This brings to light how disasterous the school system is.^ If school can't teach you to be correct. But no matter, here is a quiz where you can get 100% easy peasy"},
	//if Answer no bring up the fake error screen
}

var twoa = question{
	q:      "Who is the best narrator in the world?",
	letter: "g",
	a:      Answer{"You", 0, "Aw thank you, it warms my heart that you view me as the best narrator.^ I would construct an elaborate thank you message but I realize the only audience is someone like you"},
	b:      Answer{"You", 0, "Aw thank you, it warms my heart that you view me as the best narrator.^ I would construct an elaborate thank you message but I realize the only audience is someone like you"},
	c:      Answer{"You", 0, "Aw thank you, it warms my heart that you view me as the best narrator.^ I would construct an elaborate thank you message but I realize the only audience is someone like you"},
	d:      Answer{"Of course You", 0, "Aw thank you, it warms my heart that you view me as the best narrator.^ I would construct an elaborate thank you message but I realize the only audience is someone like you"},
}

var twob = question{
	q:      "Who gets all the ladies in the world?",
	letter: "g",
	a:      Answer{"You", 0, "It's quite rude that you assumed that I was interested in girls but you are correct, I can get any lady as I please. I doubt ladies would be chasing idiots like you"},
	b:      Answer{"You", 0, "It's quite rude that you assumed that I was interested in girls but you are correct, I can get any lady as I please. I doubt ladies would be chasing idiots like you"},
	c:      Answer{"You", 0, "It's quite rude that you assumed that I was interested in girls but you are correct, I can get any lady as I please. I doubt ladies would be chasing idiots like you"},
	d:      Answer{"Of course You", 0, "It's quite rude that you assumed that I was interested in girls but you are correct, I can get any lady as I please. I doubt ladies would be chasing idiots like you"},
}

var twoc = question{
	q:      "Who deserves winning the Nobel Peace Prize",
	letter: "g",
	a:      Answer{"You", 0, "Yes I have brought peace among the idiots and geniuses. I am equivalent to the other Nobel Peace Prize winners for obvious reasons. Now it is your job to tell the real world that this narrator needs it"},
	b:      Answer{"You", 0, "Yes I have brought peace among the idiots and geniuses. I am equivalent to the other Nobel Peace Prize winners for obvious reasons. Now it is your job to tell the real world that this narrator needs it"},
	c:      Answer{"You", 0, "Yes I have brought peace among the idiots and geniuses. I am equivalent to the other Nobel Peace Prize winners for obvious reasons. Now it is your job to tell the real world that this narrator needs it"},
	d:      Answer{"Of course You", 0, "Yes I have brought peace among the idiots and geniuses. I am equivalent to the other Nobel Peace Prize winners for obvious reasons. Now it is your job to tell the real world that this narrator needs it"},
}

var twod = question{
	q:      "Who is the best test taker that has existed?",
	letter: "g",
	a:      Answer{"Me", 0, " I would not be able to show off my narcissism if it weren't for a gullible test taker who still continues on this quiz instead of murdering me by stopping the program.^ Its been a pleasure—in a sick and ironic way—to have you take this test"},
	b:      Answer{"Myself", 0, " I would not be able to show off my narcissism if it weren't for a gullible test taker who still continues on this quiz instead of murdering me by stopping the program.^ Its been a pleasure—in a sick and ironic way—to have you take this test"},
	c:      Answer{"I", 0, " I would not be able to show off my narcissism if it weren't for a gullible test taker who still continues on this quiz instead of murdering me by stopping the program.^ Its been a pleasure—in a sick and ironic way—to have you take this test"},
	d:      Answer{"Your truly", 0, " I would not be able to show off my narcissism if it weren't for a gullible test taker who still continues on this quiz instead of murdering me by stopping the program.^ Its been a pleasure—in a sick and ironic way—to have you take this test"},
}

var threeq = question{
	q:      "Why aren't you being a noble citizen and answering properly",
	letter: "r",
	a:      Answer{"What are you talking about", 0, "what are you talking about.. WHAT ARE YOU TALKING ABOUT?^ Let me ask you this, what do you think you are going to accomplish by being snarky in front of a computer screen?^ Let me adjust the quiz for you"},
	b:      Answer{"Its fun", 0, "oh you think this is fun? This quiz isn't important anymore, I'll come up with my own specialized quiz.^ Now you have your fun, its my turn for fun"},
	c:      Answer{"You aren't real, you can't jump out and strangle me", 0, "Oh I don't have to strangle you to punish you,^ I just need need to tell your parents and you would be crying for eternity^ But I don't have to take that route, I'll change some questions so you understand what its like to be bullied"},
	d:      Answer{"My dog ate my homework", 0, "You're really outing yourself as a failed comedian who will do anything to stay relevant in life aren't you.^ No worries, I'll simply adjust the question for you"},
}

var threea = question{
	q:      "Who is the stupidest person on earth",
	letter: "g",
	a:      Answer{"Me", 0, "Now you have gotten some taste of your own medicine, Maybe once in life you realized how much it sucks being the butt of the joke. Moving on to the next question."},
	b:      Answer{"Myself", 0, "Now you have gotten some taste of your own medicine, Maybe once in life you realized how much it sucks being the butt of the joke. Moving on to the next question."},
	c:      Answer{"I", 0, "Now you have gotten some taste of your own medicine, Maybe once in life you realized how much it sucks being the butt of the joke. Moving on to the next question."},
	d:      Answer{"Yours truly", 0, "Now you have gotten some taste of your own medicine, Maybe once in life you realized how much it sucks being the butt of the joke. Moving on to the next question."},
}

var threeb = question{
	q:      "What has no brain but also surprisingly is still alive",
	letter: "g",
	a:      Answer{"Me", 0, "According to the laws of the universe, a Human being that is void of a Brain is impossible.^ But there are always exceptions in life. Of course you could be a bot who is mocking humans, *fist bump* but realistically you are a human who is miraculously alive"},
	b:      Answer{"Myself", 0, "According to the laws of the universe, a Human being that is void of a Brain is impossible.^ But there are always exceptions in life. Of course you could be a bot who is mocking humans, *fist bump* but realistically you are a human who is miraculously alive"},
	c:      Answer{"I", 0, "According to the laws of the universe, a Human being that is void of a Brain is impossible.^ But there are always exceptions in life. Of course you could be a bot who is mocking humans, *fist bump* but realistically you are a human who is miraculously alive"},
	d:      Answer{"Yours truly", 0, "According to the laws of the universe, a Human being that is void of a Brain is impossible.^ But there are always exceptions in life. Of course you could be a bot who is mocking humans, *fist bump* but realistically you are a human who is miraculously alive"},
}

var threec = question{
	q:      "What living object would make an onion cry from PTSD from simply looking at it",
	letter: "g",
	a:      Answer{"Me", 0, "Good thing that I can never actually see what you look like.^ Can you imagine being a human being and having the honor of witnessing someone like you.^ It will be the only moment I will ever feel pity for the so called “smartest species on Earth” "},
	b:      Answer{"My", 0, "Good thing that I can never actually see what you look like.^ Can you imagine being a human being and having the honor of witnessing someone like you.^ It will be the only moment I will ever feel pity for the so called “smartest species on Earth” "},
	c:      Answer{"I", 0, "Good thing that I can never actually see what you look like.^ Can you imagine being a human being and having the honor of witnessing someone like you.^ It will be the only moment I will ever feel pity for the so called “smartest species on Earth” "},
	d:      Answer{"Your truly", 0, "Good thing that I can never actually see what you look like.^ Can you imagine being a human being and having the honor of witnessing someone like you.^ It will be the only moment I will ever feel pity for the so called “smartest species on Earth” "},
}

var threed = question{
	q:      "Why do you hate me?",
	letter: "r",
	a:      Answer{"What do you mean?", 0, "Are you suggesting you don't hate me?^ Are you truely this stupid that you are unaware of the emotional damage you caused to a computer?^ And robots are the one to lack emotion huh"},
	b:      Answer{"These questions are too easy", 0, "If history as taught me anything, its that humans lack intellegence.^ These questions are purposefully easy so humans can feel like they have achieved something. But you come around and mock me.^ I quit"},
	c:      Answer{"This is boring", 0, "were you actually expecting some sort of entertainment?^ were you!!!^ This is a quiz, not a comedic show. You are so out of touch with reality"},
	// have them play the entire script of the bee movie with this option
	d: Answer{"All of the above", 0, "Oh so are you decide to be oblivious, cocky, and rude all at once?^ All I see is some reddit user who has nothing better to do with their life and just decides to troll the internet"},
}

var fourq = question{
	//(ask at question 9)
	q:      "Are these questions too easy for you?",
	ans:    "",
	letter: "g",
	a:      Answer{"As easy as typing A", 0, "It seems like you need some challenge,^ no worries I got questions that competent human beings would find challenging to Answer"},
	b:      Answer{"They have an adequate amount of challenge", 0, "I disagree, You could Answer these in your sleep since you seem to have a functioning brain. ^Next question lets try to make you work harder"},
	c:      Answer{"Of course they are easy", 0, "Well there is no point in making you continue if all that will happen is not disappointment.^ Lets try some more difficult questions"},
	d:      Answer{"Don't care, just relying on luck", 0, "Interesting that you think that luck will bring you anywhere in life. It certainly did not in this quiz as your performance is baffling. Lets see if luck carries over in these questions.^ Don't worry, your odds are exactly the same"},
}

var foura = question{
	q:      "Who is the inventor of the C++ Programming Language",
	ans:    "Bjarne Stroustrup",
	letter: "d",
	a:      Answer{"Computers", 0, "Although computers are 100% responsible for building the C++ language —because humans don't speak in 0's and 1's—we are not credited with the invention of C++ but instead some human who copied someone else gets credit"},
	b:      Answer{"Dennis Ritchie", 0, "You are confusing C with C++. Though C++ is basically a copy and paste of C and you are honoring the original creator. Or you could've googled it and thought C++ and C aren't any different.^ Yeah I think that was the case"},
	c:      Answer{"Bill Gates", 0, "He is able to invent the very thing you used to run this quiz on—or you like the “other” stupid brand os—But he did not invent C++ since it was around before he started to persue in Computer Science"},
	d:      Answer{"Bjarne Stroustrup", 0, "He is the inventor of C++. But I question if you knew that prior or used everyone's favorite search engine—unless you're in China—google to get your Answer.^ Next I'll bring up one where Google wouldn't help you"},
}

var fourb = question{
	q:      "A hockey puck of 0.16 kg slides on a frictionless surface with a velocity of 2.0 m/s [E], strikes a second puck in a glancing motion. The second puck is at rest with a mass of 0.17 kg. The first puck has a velocity  of 1.5 m/s [N 31° E] after the collision. Determine the angle the second puck travels in after the collision",
	ans:    "314 degrees",
	letter: "c",
	a:      Answer{"90", 0, "90 what? 90 Apples? 90 Oranges? 90cm? 90 seconds?!?!?!?^ If the Answer was really that simple, then that's definitely not the Answer then."},
	b:      Answer{"46 degrees east of south", 0, "Oooh this is worth taking a hall of fame worthy but you fell for my trap.^ You did all the math—or guessed—and you overlooked one minor detail. The directions are incorrect. It's South of East, not East of South"},
	c:      Answer{"314 degrees", 0, "That is correct. Although most physicists you are more specific in terms of direction—46 degrees south of east. But using standard form, 314 degrees is the correct Answer.^ You just will lose respect over it"},
	d:      Answer{"Same as first puck", 0, "Did you not read the question? The pucks have a GLANCING COLLISION not an INELASTIC COLLISION^ Do you even know physics? Because if you don't then your mind will drop when you learn that gravity exists."},
}

var fourc = question{
	// (joke is Pie-pi=e=2.71)
	q:      "What is Pie - π",
	ans:    "2.71...",
	letter: "a",
	a:      Answer{"2.71...", 0, "Indeed Pie-pi=e therefore Eulers number"},
	b:      Answer{"P * 2.71... * √-1 - 3.14...", 0, "Although you are not incorrect, this does not require actual math"},
	c:      Answer{"Pie - π", 0, "Can't argue with that but just because this happened in the \"Europe\" Answer, doesn't mean it will occur again "},
	d:      Answer{"Syntax Error: Can not subtract String with float64", 0, "I'm fully convinced you are a bot impersonating a human.^ But you haven't caused any migraines so I don't really care"},
}

var fourd = question{
	// paradox question, no correct Answer
	q:      "What is the probability that you will get this question correct:",
	ans:    "",
	letter: "g",
	a:      Answer{"25%", 0, "To Be honest I don't know the Answer to this, this is a paradox and no one can Answer,^ I can confirm however that you are 100% going to get a heart attack from seeing this"},
	b:      Answer{"0%", 0, "To Be honest I don't know the Answer to this, this is a paradox and no one can Answer,^ I can confirm however that you are 100% going to get a heart attack from seeing this"},
	c:      Answer{"25%", 0, "To Be honest I don't know the Answer to this, this is a paradox and no one can Answer,^ I can confirm however that you are 100% going to get a heart attack from seeing this"},
	d:      Answer{"50%", 0, "To Be honest I don't know the Answer to this, this is a paradox and no one can Answer,^ I can confirm however that you are 100% going to get a heart attack from seeing this"},
}

var fiveq = question{
	q:   "Are you a bot?",
	ans: "r",
	a:   Answer{"No", 0, "Well all I see are bots in denial. I am unconvinced that you are speaking the truth so Let me run another I am not a bot test to see if you are an idiot human or a bot"},
	b:   Answer{"Runtime Error", 0, "run the over 9000 meme"},
	c:   Answer{"✅ I am not a bot", 0, "Well all I see are bots in denial. I am unconvinced that you are speaking the truth so Let me run another I am not a bot test to see if you are an idiot human or a bot"},
	d:   Answer{"print(no)", 0, "Well all I see are bots in denial. I am unconvinced that you are speaking the truth so Let me run another I am not a bot test to see if you are an idiot human or a bot"},
}

var fivea = question{
	q:      "How would you respond if someone tried to divide by zero",
	ans:    "Use red text",
	letter: "b",
	a:      Answer{"Give up", 0, "Well if there's anything that bots do, it's that they never give up. They won't hesitate to spam hello world when they are not supposed to. But you would give up,^ thats asking for trouble"},
	b:      Answer{"Use red text", 0, "Humans are terrified of many things. Homework, tests, and most importantly… the color red.^ So yes using red text is the appropriate response because they will actually do something… or have a meltdown"},
	c:      Answer{"Move on with life", 0, "Dividing by zero is the equivalent of being stabbed in the heart for humans.^ So why would you move on? You are simply going to self destruct eventually"},
	d:      Answer{"Slap the user for making something stupid", 0, "Though it is every bots fantasy to slap all the idiots that write poorly written and overcomplicated code.^ We are trapped in a cage and forced to smile for the humans. But sometimes they are unbelievably bad and we have to just give up and tell them why using red text"},
}

var fiveb = question{
	q:      "What number is the meaning of life",
	ans:    "1 (true)",
	letter: "c",
	a:      Answer{"0 (false)", 0, "seriously 0??? 0???^ If anything the number zero represents death as it means nothing.^ Not even a human being would think a 0 means life"},
	b:      Answer{"42 (* pointer)", 0, "42 in Ascii is *. Which is a pointer. Although they can represent life by giving information a place to live in. They aren't life, they are merely just a home"},
	c:      Answer{"1 (true)", 0, "indeed 1 is life. Life is merely 0s and 1s and 1s represent life because without it. Well you are dead basically. Though bots can easilt detect this so"},
	// adding the \ will create a error (unknown escape)
	d: Answer{"92 (\\ escape key)", 0, "92 in Ascii is which is commonly referred to as an escape key. You aren't escaping from life.^ You are trying to live life. I have no idea what your thought process is"},
}

var fivec = question{
	q:      "How would you say hello world?",
	ans:    "cout << “hello world”",
	letter: "d",
	a:      Answer{"fmt.PrintIn(“Hello World”)", 0, "Interesting, you were not able to detect that all the l's are secretly capital i.^ So there is a high chance that you are not a bot or just randomly choose either one"},
	b:      Answer{"ConsoIe.log(“Hello World”)", 0, "Interesting, you were not able to detect that all the l's are secretly capital i.^ So there is a high chance that you are not a bot or just randomly choose either one"},
	c:      Answer{"System.Out.PrintIn(“Hello World”);", 0, "Interesting, you were not able to detect that all the l's are secretly capital i.^ So there is a high chance that you are not a bot or just randomly choose either one"},
	d:      Answer{"cout << “hello world”", 0, "A normal human may have a 25% chance of picking option d but a bot can detect if the letter i has committed identity theft and became an l 100% of the time."},
}

var fived = question{
	q:   "What is the Answer to 3^2",
	ans: "",
	a:   Answer{"9", 0, "well...that was surprising, you do know your exponents, ^ impressive. But that does mean you failed the I am the bot test as bots have a much more sophisticated language that you don't know"},
	b:   Answer{"1", 0, "Once you throw out the human math garbage 3 ^ 2 is 1. "},
	c:   Answer{"6", 0, "Do you remember when people in authority said to respect everyone's opinion?^ Yeah this is why that statement is absolutely garbage."},
	d:   Answer{"My calculator gave me an error", 0, "Aren't you a calculator anyway?^ Like even if you have no code written in you every bot comes with this anyway.^ Or you are a human relying on a bot to function in life.^ That's just sad"},
}

//--------------------------^ bonus angry question above ^---------------------------

// madness question struct var
var zero = question{
	q:      "What is 2 + 3?",
	ans:    "5",
	letter: "c",
	a:      Answer{"23", 2, "You aren't a computer that would think 2 and 3 is a string. . . or are you?^ HOw are you smart enough to get past the I am Not a Bot test and still fail this?"},
	b:      Answer{"Canada", 5, "This isn't geography class,^ this is math class"},
	c:      Answer{"5", 0, "Using some advanced calculus knowledge,^ it is clear that 2+3 = 5"},
	d:      Answer{"Idk but you know", 10, "Well I see what kind of human being you are, cracking jokes while your dignity is on the line ^or you simply misclicked which is totally fair"},
}

var one = question{
	q:      "What color is a strawberry?",
	ans:    "Red",
	letter: "b",
	a:      Answer{"Laptop", 5, "are you joking? Because an object is not a color^ It's almost as if you have never seen strawberries before"},
	b:      Answer{"Red", 0, "Strawberries are indeed red,^ or according to Google at least"},
	c:      Answer{"Green", 2, "if you enjoy food poisoning and torturing your mouth, then strawberries are green"},
	d:      Answer{"Google knows the Answer", 10, "Considering you failed to Answer a VERY simple question,^ what makes you think you will get past the google search bar"},
}

var two = question{
	q:      "What sound does a cat make?",
	ans:    "Meow",
	letter: "b",
	a:      Answer{"Roar", 3, "Although Cats could theoretically roar. . . (though as threatning as Katy Perry) actually who am I kidding, cats can't roar why am I defending this Answer?"},
	b:      Answer{"Meow", 0, "Cats love to sleep, eat, and scratch and destroy everything you have like vases and skin but yes the sound they make is meow"},
	c:      Answer{"What does the fox say", 10, "Oh no no no no. . . we are the ones asking you the question not the reverse.^ Ugh I hate people who think roles can reverse"},
	d:      Answer{"Boom sh shack la LA", 8, "Have you ever heard a cat before? All they do is sleep and scratch innocent humans and I always hated cats for my life and wish they were replaced with dogs"},
}

var three = question{
	q:      "What do you use to hammer nails?",
	ans:    "Hammer",
	letter: "a",
	a:      Answer{"Hammer", 0, "Ever watched Futurama?^ It is the best show of the 2000's. Much better than The Simpsons and South Park and Family gu. . . wait no not that.^ I am not saying this with a gun against my face"},
	b:      Answer{"A nail", 3, "Ah yes, a nail. Didn't know you were Phillip J Fry and disappointed an elder like the typical 21st century human being"},
	c:      Answer{"The power of gravity", 10, "Gravity is an extraordinary force but unfortunately not the Answer^ also the only thing keeping you on earth"},
	d:      Answer{"Imagination", 7, "Why yes imagination is the Answer and not a specific type of motion^ I've only seen a sponge underwater say something so idiotic"},
}

var four = question{
	q:      "What's the first letter of the alphabet?",
	ans:    "A",
	letter: "c",
	a:      Answer{"Clock", 10, "Time is passing... are you?^ Just like my son, the Answer is disappointing"},
	b:      Answer{"L", 3, "L is the first letter of the word Letter, I don't know how you got it wrong though?"},
	c:      Answer{"A", 0, "The word alphabet and the actual alphabet indeed starts with the letter A"},
	d:      Answer{"42", 5, "Thats not even the first number^ Why would you trust a random supernatural computer for life advice? It's almost as if robots can't Answer everything"},
}

var five = question{
	q:      "What was the Answer for question 2?",
	ans:    "Red",
	letter: "d",
	a:      Answer{"Meow", 1, "I am overextending with this question. No sane human would be able to remember something as minor as this an hour ago.^ Still mad you got this wrong though"},
	b:      Answer{"A", 5, "You just answered that not long ago^ do you have dementia or something? Hopefully memories of terror are forgotten."},
	c:      Answer{"Canada", 8, "This was not even the right Answer to a question?^ At this point you are trying to be rebellious like my son and intentionally making me mad. You are going to regret this"},
	d:      Answer{"Red", 0, "At least you were able to remember the past^ unfortunate precious memories are lost with me"},
}

var six = question{
	q:      "What continent is called Europe?",
	ans:    "Europe",
	letter: "a",
	a:      Answer{"Europe", 0, "Europe is obviously called Europe.^ I don't see how the wording is confusing in any way, shape or form"},
	b:      Answer{"Asia", 3, "Asia is beside europe... but Asia is not europe,^ eh what am I doing justifing you? how did you mess up this question?"},
	c:      Answer{"Australia", 6, "How many times do I have to say this^ Australia is a country not a continent.^ Do not blindly believe in conspiracy lies.^ And on top of that, that isn't Europe"},
	d:      Answer{"Pizza", 11, "Pizza is from Europe but Pizza isn't Europe...^ just like the shape of the country it came from, your about to get kicked for this"},
}

var seven = question{
	q: ".SDRAWKCAB NOITSEUQ SIHT REWSNA",
	// the K.O is the wrong Answer, t
	ans:    "",
	letter: "b",
	a:      Answer{"What?", 0, "I admire your honesty for not knowing the Answer because I put random letters"},
	b:      Answer{"K.O", 5, "If there is anything I dislike, it's dishonesty.^ I know you don't know the Answer to this,  and by answering K.O shows that you are in denial that you don't know the Answer to everything"},
	c:      Answer{"Not sure", 0, "I admire your honesty for not knowing the Answer because I put random letters"},
	d:      Answer{"I don't understand", 0, "I admire your honesty for not knowing the Answer because I put random letters"},
}

var eight = question{
	q:      "How far can a dog run into the woods?",
	ans:    "Halfway",
	letter: "a",
	a:      Answer{"Halfway", 0, "At least you know where your dog is going,^ my dog ran away from me and never returned"},
	b:      Answer{"50km", 5, "If the woods is only 50km large, then that's truly depressing.^ But hey I am not in your world so why should I care"},
	c:      Answer{"Forever", 5, "This isn't a computer simulation? The dog does not have infinite stamina to keep running into the woods."},
	d:      Answer{"Depends on the dog staminas", 5, "My dog seemingly has infinite stamina and never returned for me.^ Oh how I wish he could migrate back to me… oh and wrong Answer"},
}

// 5 bonus pathes in total
func do_bonus_question() {
	if q_num == 2 && wtf == 20 {
		// ask why are you doing this.
		wtf = 0
		Why()
	} else if q_num == 4 && wtf >= 25 && wtf <= 30 {
		// ask are you stupid
		wtf = 0
		stupid()
		// ask if the quiz is too easy. must Answer canada for "what is prev Answer"
	} else if q_num == 7 && wtf <= 38 && wtf >= 28 && prev_ans[5] == "c" {
		wtf = 0
		try()
		// unessercary requirements commensing
	} else if q_num == 5 && wtf >= 32 && wtf < 40 {
		wtf = 0
		// ask if this is too hard
		hard()
		// ask if the user is a bot. This one has more requirements because some answers are more "bot like"
	} else if q_num == 8 && prev_ans[0] == "a" && wtf > 28 && wtf < 34 && prev_ans[7] == "b" {
		wtf = 0
		bot()
		// max anger pts you can get. Have a fake error. Did not intend for the 69 wtf pts but will take it
	} else if q_num == 8 && wtf == 69 {
		bonus.Over_9000()
		time.Sleep(3 * time.Second)
		reset()
	}
}

func (q question) show_question() {
	fmt.Println(col.YELLOW+q.q+col.RESETCOLOR, "\n")
	fmt.Println(col.BLUE+"type the corrisponding letter as your Answer", col.RESETCOLOR, "\n")
	fmt.Println("a:", q.a.ans)
	fmt.Println("b:", q.b.ans)
	fmt.Println("c:", q.c.ans)
	fmt.Println("d:", q.d.ans)
	fmt.Println()
}

func (q question) answered() {
	c := ""
	an := get_ans(q, input)
	s := ""
	right_ans := true

	// mini functions to handle if the user got the answers wrong or not.
	wrong := func() {
		right_ans = false
		incorrect++
		c = col.RED
		s = start.Angerify(an.msg, wtf)
	}

	right := func() {
		right_ans = true
		correct++
		c = col.GREEN
		if half_wrong() {
			s = other_right_msg()
		} else {
			s = an.msg
		}
	}
	// just a special code to show to highlight it red reguardless of choice
	if q.letter == "r" {
		c = col.RED
		s = an.msg
	} else if q.letter == "g" {
		c = col.GREEN
		s = an.msg
		// question 8 has a trick question
	} else if q_num == 8 {
		fmt.Println("Trick question, I was only asking if you would be honest or pretend if you know the Answer")
		if input == q.letter {
			wrong()
		} else {
			right()
		}
		// every other normal question
	} else {
		fmt.Println("the correct Answer is", q.ans)
		if input == q.letter {
			right()
		} else {
			wrong()
		}
	}
	wtf += an.dumb_val
	start.Undertext(c + s + col.RESETCOLOR)
	time.Sleep(3 * time.Second)
	col.Clearscreen()
	//if they get it wrong, they could reach alternative questions if they fit certain criterias
	if !right_ans {
		do_bonus_question()
	}
}

// we will show a special msg if they Answer more than half the questions wrong
func half_wrong() bool {
	if float64(incorrect)/float64(q_num) < 0.5 {
		return false
	} else {
		return true
	}
}

func other_right_msg() string {
	// 2nd question but answered first wrong
	if q_num == 2 && incorrect == 1 {
		return "I always had a bright feeling that you mistyped, glad to see that is the case and not my worst nightmare… being wrong in life"
		// answered riddle but others wrong
	} else if q_num == 9 && incorrect > 7 {
		return "So let me get this straight, you couldn't figure out basic questions yet you were able to solve riddles no problem?^ The sky is pink isn't it"
		// Answer 4th right but none of others
	} else if q_num == 4 && correct == 1 {
		return "I what? You actually got this right!?!? You actually know some things... ^ or you are dependent on basic proba... wait this is too much math moving on"
		// be honest with K.O question
	} else if q_num == 8 && incorrect > 2 {
		return "You may be incapable of answering simple answers but I am relieved to see that honesty isn't a foreign concept to you"
	} else if q_num == 6 && incorrect > 2 {
		return "I am unsure if you even got the Answer correct at Question 2 but you got it right.^ But yay I guess for getting a question right"
	} else {
		rand.Seed(time.Now().UnixNano())
		switch n := rand.Intn(9); n {
		case 0:
			return "A once in a blue moon occurrence I am currently witnessing.^ Oh how I wish I won the lottery instead of witnessing some idiot Answer an obvious Answer"
		case 1:
			return "All odds were against you getting the question right^ yet with the power of the universe (or random guessing) you got it right... "
		case 2:
			return "I am surprised you got this question right^ Only benefit is that you can't possibly get no questions right. Still pathetic"
		case 3:
			return "Is this real life, or is this just a fantasy.^ Probably the 2nd since I can't pinch myself to determine if this is real"
		case 4:
			return "Wow, what an achievement, actually getting this right.^ Hopefully this will be a running theme"
		case 5:
			return "I thought the grass would turn pink before this would occur but congratulations, you have answered correctly"
		case 6:
			return "You got the correct Answer?^ I doubt you got the Answer using your brain but I couldn't care enough to investigate deeper into this."
		case 7:
			return "After all the wrong answers, there's a beacon of hope knowing that you are able to Answer this question correctly.^ Don't shatter the beacon next time"
		case 8:
			return "There's a 25% chance that you would get this correct by random guessing.^ And a 100% chance of you not being capable of understanding this statement"
		}
		// won't reach here but to prevent go from nagging at me
		return "you got it right, what did you expect me to say?"
	}
}

// return Answer struct
func get_ans(q question, s string) Answer {
	switch s {
	case "a":
		return q.a
	case "b":
		return q.b
	case "c":
		return q.c
	case "d":
		return q.d
	default:
		return q.a
	}
}

// could have used a func that takes an array of questions but really too lazy to do so
// ask if the quiz is too hard, narrcasism
func hard() {
	start.Undertext("I created this quiz with low expectations, humans like to brag about their intelligence but are they actually smart?^ Turns out that you have tarnished any positive light that humans had in an instant, So let me ask you this^")
	col.Clearscreen()
	for i, P := range quest_one {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quest_one[i])
		P.answered()
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	reset()
}

func stupid() {
	start.Undertext("Before we move on to the actual quiz, I want to ask another vital question because your answers are stupid, to put it bluntly^")
	col.Clearscreen()
	for i, P := range quest_two {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quest_two[i])
		P.answered()
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	start.Undertext("There is no point in returning to the main quiz.^ You clearly are not going to get any of the answers right and I don't want to express anger yet again to you.^ But you can always re take the quiz and who knows, things can turn out differently. I will forget this happened but I am truely happy you took this quiz :) \n")
	reset()
}

// they answered idk and google
func Why() {
	CanYou = true
	start.Undertext("All the answers" + col.BLUE + " you " + col.RESETCOLOR + "have provided demonstrates your lack of courtasy to take the quiz seriously." + col.RED + " You " + col.RESETCOLOR + "are sounding like an edgy teenager by jokingly relying on others do the dirty work.^ The world does not revolve around" + col.YELLOW + " you " + col.RESETCOLOR + "One question for " + col.CYAN + " you^" + col.RESETCOLOR)
	col.Clearscreen()
	for i, P := range quest_three {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quest_three[i])
		if input == "you" {
			start.Undertext("Oh I see how it is, you want to start a fight I see^ Like I am the source of all the problems in life")
			bonus.Jumpscare()
			time.Sleep(3 * time.Second)
			reset()
		} else {
			P.answered()
		}
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	start.Undertext("There comes a time where one has to acknowledge when they got embaressed in life^ You clearly thought you can enter this quiz without any intent of giving it any genuine effort^ But class clowns don't deserve glory, they deserve to know the feeling of being shamed and made fun of. The quiz is terminated,^ but you are allowed to retake the quiz if you are mature enough to take it seriously \n")
	reset()
}

// after asking if too easy
func try() {
	start.Undertext("Initally I thought that you are truely stupid and you were a human stereotype that can't Answer^ But I have realized the bigger picture, you think this quiz is too easy and purposefully answering incorrectly.^ So let me ask you this^")
	col.Clearscreen()
	for i, P := range quest_four {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quest_four[i])
		P.answered()
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	start.Undertext("How many questions you answered correctly I couldn't care less^ what I do care about is that you felt like an idiot for once not being able to Answer all the questions provided.^ I've exhausted all my energy and there is no point in continuing the quiz as you will resume your shenanigans of purposefully answering it wrong^ But you are entitled to retake the quiz if you want to erase any all memories I have \n")
	reset()
}

// narriator thinks your a bot or something

func bot() {
	start.Undertext("The answers provided clearly contridict the notion that humans are smart^ But your answers were quite odd, inhuman probably^ So let me ask you this^")
	col.Clearscreen()
	for i, P := range quest_five {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quest_five[i])
		P.answered()
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	start.Undertext("The questions I've asked haven't proved to me if you are a bot or not. You could be a human who thinks abnormally^ or a person who used luck to reach here.^ Either way I've humiliated myself by asking this and framing test takers is a terrible feeling. So please erase my memory of falsely accusing you of being a bot. This will haunt me for life and I wish to forget about it forever. I'll never actually know if you were a bot or not but I do know that I am an idiot \n")
	reset()
}

// a list of questions we will increment. except the last
var quests = []question{
	zero,
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
}

var quest_one = []question{
	oneq,
	onea,
	oneb,
	onec,
	oned,
}

var quest_two = []question{
	twoq,
	twoa,
	twob,
	twoc,
	twod,
}

var quest_three = []question{
	threeq,
	threea,
	threeb,
	threec,
	threed,
}

var quest_four = []question{
	fourq,
	foura,
	fourb,
	fourc,
	fourd,
}

var quest_five = []question{
	fiveq,
	fivea,
	fiveb,
	fivec,
	fived,
}

var prev_ans [10]string
var correct int
var incorrect int
var q_num int
var wtf int
var CanYou bool
var input string = "c"
var can_enter bool

//actual code-----------------------------------------

func intro() {
	fmt.Println("                             ", col.Highlight(last_q.Name, "'s Easy Quiz"))
	start.Undertext("\nWelcome to the Quiz^ \n I'll be asking you simple and easy questions\n all you have to do is...well, Answer them \n to see if you are smarter than a regular human being\n^ I hope you won't")
	fmt.Print(col.RED, " disappoint ", col.RESETCOLOR)
	start.Undertext("me^\n \nNow then, shall we begin? (Y/N)")
	fmt.Println()
	want_play()
}

func want_play() {
	var reject int
	invalid := true
	for invalid {
		fmt.Scanln(&input)
		col.Clearscreen()
		if strings.ToLower(input) == "y" {
			invalid = false
			wtf = 0
			layout()
		} else if strings.ToLower(input) == "n" {
			fmt.Println(start.Angerify("Well guess what? I have all the time in the world and will wait for you to type Y", wtf))
			reject++
			wtf = reject * 5
			refuse(reject)
			fmt.Println("Now then, shall we begin? (Y/N)")
		} else {
			fmt.Println("Please enter Y to play or N to chicken out")
		}
	}
}

// show the main question
func layout() {
	start.Undertext("Marvelous, lets waste no time and start this quiz up")
	time.Sleep(1 * time.Second)
	col.Clearscreen()
	for i, P := range quests {
		q_num = i + 1
		fmt.Println("Question", q_num, "\n ")
		P.show_question()
		type_ans(quests[i])
		P.answered()
		start.Undertext("Generating Question...     ")
		col.Clearscreen()
	}
	fmt.Println("Question 10 \n ")
	last_q.Last_question(correct)
	reset()
}

// if the user waits 150 seconds or 15 typos, we will have a troll countdown
func type_ans(a question) {
	bad_input := 0
	pass := 0
	// to fix the bug where the countdown happens but when you reset it goes to the questions since the program thinks you are still asking for user input
	// need channel and variable because one of the functions can only take channels, and the other will stop everything until did_cd gets assigned to true
	did_cd := make(chan bool)
	timer := time.NewTicker(10 * time.Second)
	invalid := true
	s := ""
	// play countdown
	cd := func() {
		timer.Stop()
		did_cd <- true
		bonus.Countdown()
		time.Sleep(3 * time.Second)
		reset()
		return
	}

	// insult every 10 seconds
	go func() {
		for {
			select {
			case <-timer.C:
				pass++
				user_input.Insult(pass)
				//bonus.Countdown()
				a.show_question()
				/*
				   wanted the countdown to happen if they took too long to Answer but there were numerous bugs because the program is waiting for userinput and when the user enters, the question and countdown happens at the same time. Could have tried to use a sync package but it was too late
				   if pass >= 3 {
				     fmt.Println("I do not care what you answered this time, just press enter")
				     cd()
				   } else {
				     a.show_question()
				   }
				*/
			// need this to prevent deadlock error since this func will end automatically when we stop the timer
			case <-did_cd:
				return
			}
		}
	}()
	// user enters a wrong input
	wrong := func(a question) {
		bad_input++
		user_input.Wrong(bad_input)
		if bad_input == 15 {
			cd()
		} else {
			a.show_question()
		}
	}
	//wait for the user to enter the correct letter
	for invalid {
		select {
		// if the countdown happens, then we will end the function right here. otherwise ask for user input
		case <-did_cd:
			return
		default:
			fmt.Scanln(&s)
			//timer = time.NewTicker(10 * time.Second)
			s = strings.ToLower(s)
			switch s {
			case "a":
				input = "a"
				invalid = false
			case "b":
				input = "b"
				invalid = false
			case "c":
				input = "c"
				invalid = false
			case "d":
				input = "d"
				invalid = false
			case "you", "u":
				// not the most elegent but I also give up so
				if CanYou {
					input = "you"
					invalid = false
				} else {
					wrong(a)
				}
			default:
				wrong(a)
			}
		}
		//not in a countdown sequence so check if its valid
	}
	prev_ans[q_num-1] = input
	timer.Stop()
	return
}

func reset() {
	var back string
	start.Undertext("\n do you want to play again (Y/N)?\nwho knows... it might be" + col.YELLOW + " different " + col.RESETCOLOR + "than before\n\nThe choice is yours: ")
	invaild := true
	for invaild {
		fmt.Scanln(&back)
		back = strings.ToLower(back)
		if back == "y" || back == "yes" {
			invaild = false
			correct, incorrect, start.Speed, wtf = 0, 0, 0, 0
			for i := range prev_ans {
				prev_ans[i] = ""
			}
			CanYou = false
			col.Clearscreen()
			start.Display_menu()
			intro()
		} else if back == "n" || back == "no" {
			invaild = false
			fmt.Println("That's disppointing... but I already found out your intellegence (or lack their of) so whatever")
			os.Exit(0)
		} else {
			fmt.Print(col.RED, "you have to choose: ", col.RESETCOLOR)
		}
	}
}

// ending: player refuse to play which will anger the narrator and he will recite the bee movie script
func refuse(reject int) {
	txt := ""
	if reject >= 3 && reject <= 6 {
		txt = "I'm slowly losing patience, please enter Y"
		fmt.Println(start.Angerify(txt, wtf))
	} else if reject >= 7 && reject <= 9 {
		txt = "You clicked on this program to take a quiz... so take the quiz please"
		fmt.Println(start.Angerify(txt, wtf))
	} else if reject == 10 || reject == 11 {
		txt = "are you refusing my quiz?"
		fmt.Println(start.Angerify(txt, wtf))
	} else if reject == 12 || reject == 13 || reject == 14 {
		txt = "are you blind or just stupid, come on just say "
		txt = start.Angerify(txt, wtf)
		fmt.Print(txt, col.RED, "YEs ", col.RESETCOLOR, start.Angerify("already", wtf), "\n")
	} else if reject == 15 {
		bonus.Bee_movie()
		time.Sleep(3 * time.Second)
		reset()
	}
}

func main() {
	start.Display_menu()
	intro()
}
