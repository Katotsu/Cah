package cardsagainstdiscord

func init() {
        pack := &CardPack{
	    	        Name:        "america",
	    	        Description: "America expansion pack",
		            Prompts: []*PromptCard{
                        &PromptCard{Prompt: `Oh shit! I just got %s.`},
                        &PromptCard{Prompt: `They’re bringing drugs. They’re bringing crime. They’re rapists. And some, I assume, are %s.`},
                        &PromptCard{Prompt: `When asked about the biggest threat facing the nation, 60% of Americans said %s.`},
                        &PromptCard{Prompt: `Cry havoc and let slip %s!`},
                        &PromptCard{Prompt: `The fault, dear Brutus, is not in our starts, but in %s.`},
                        &PromptCard{Prompt: `Now is the winter of our discontent, made glorious summer by %s.`},
                        &PromptCard{Prompt: `You see, son, baseball is like %s. Don't overthink it.`},
                        &PromptCard{Prompt: `We want a pitcher, not %s!`},
                 },
                 
                 Responses: []ResponseCard{
                        `Our dump asshole president.`,
                        `Mexicans and Muslims.`,
                        `How easy is it to climb over a wall with a ladder.`,
                        `Getting deported.`,
                        `Guacamole`,
                        `A baby panda that just can't keep its eyes open.`,
                        `Two baby otters holding hands.`,
                        `A laundry basket full of baby bunnies.`,
                        `Two hamsters nibbling on a tiny burrito.`,
                        `A stack of puppies in a trench coat.`,
                        `A really good dog.`,
                        `My money stolen by Cards Against Humanity and redistributed to people poorer than me.`,
                        `A full refund from Cards Against Humanity, paid for by people richer than me.`,
                        `A $1,000 check from Cards Against Humanity!`,
                        `The truth.`,
                        `White people becoming a minority.`,
                        `Slowly shaking my head at a sad statistic.`,
                        `A gay Republican somehow.`,
                        `The shocking stupidity of the American public.`,
                        `Speaking in iambic pentameter.`,
                        `The significance of eyes in King Lear.`,
                        `Taking a minute to really unpack what Shakespeare is getting at in this scene.`,
                        `Teaching my son to love the Red Sox and hate minorities.`,
                        `Lou Gehrig's Disease.`,
                        `Yelling "I got it, I got it!"`,
                        `My drunk father screaming from the bleachers.`,
                   },
	}

	AddPack(pack)
}
                        
