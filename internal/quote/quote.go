package quote

import (
	"math/rand"
	"time"
)

func Random() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	quotes := []string{
		"Why did the chicken cross the playground? To get to the other slide!",
		"The early bird might get the worm, but the second mouse gets the cheese.",
		"I used to play sports. Then I realized you can buy trophies. Now I’m good at everything.",
		"If at first you don’t succeed, skydiving is not for you.",
		"I’m on a whiskey diet. I’ve lost three days already.",
		"Teamwork is important; it helps to put the blame on someone else.",
		"I told my wife she was drawing her eyebrows too high. She looked surprised.",
		"I’m reading a book about anti-gravity. It’s impossible to put down!",
		"I told my computer I needed a break. Now it won’t stop sending me Kit-Kats.",
		"Two antennas met on a roof, fell in love and got married. The ceremony wasn’t much, but the reception was excellent.",
		"I couldn't figure out how to put my seatbelt on. Then it clicked.",
		"I used to be a baker, but I couldn't make enough dough.",
		"I'm friends with all electricians. We have good current between us.",
		"Don't trust atoms. They make up everything.",
		"Parallel lines have so much in common. It’s a shame they’ll never meet.",
		"Did you hear about the mathematician who’s afraid of negative numbers? He'll stop at nothing to avoid them.",
		"I told my friend 10 jokes to make him laugh. Sadly, no pun in ten did.",
		"I'm reading a horror story in Braille. Something bad is going to happen, I can feel it.",
		"I asked the librarian if the library had any books on paranoia. She whispered, 'They're right behind you!'",
		"Why don't scientists trust atoms? Because they make up everything.",
		"Why did the scarecrow win an award? Because he was outstanding in his field.",
		"What do you call fake spaghetti? An impasta.",
		"Why did the bicycle fall over? Because it was two-tired.",
		"What do you call an alligator in a vest? An investigator.",
		"How do you make a tissue dance? You put a little boogie in it.",
		"Why did the golfer bring two pairs of pants? In case he got a hole in one.",
		"What do you call a bear with no teeth? A gummy bear.",
		"What do you call a cow with no legs? Ground beef.",
		"What did one ocean say to the other ocean? Nothing, they just waved.",
		"Why did the tomato turn red? Because it saw the salad dressing.",
		"What do you call a fish with no eyes? Fsh.",
		"How do you organize a space party? You planet.",
		"How do you catch a squirrel? Climb a tree and act like a nut.",
		"Why did the math book look sad? Because it had too many problems.",
		"What do you call a snowman with a six-pack? An abdominal snowman.",
		"How do you make holy water? You boil the hell out of it.",
		"What do you call cheese that isn't yours? Nacho cheese.",
		"Did you hear about the cheese factory that exploded? There was nothing left but de-brie.",
		"Why couldn't the leopard play hide and seek? Because he was always spotted.",
	}

	return quotes[r.Intn(len(quotes))]
}
