package cmd_handler

import (
	"bytes"
	"strings"
	"unicode"
)

type profanityWord struct {
	words map[string]bool
}

var _profanityWords profanityWord

const (
	wordsContent       = "4r5e\n5h1t\n5hit\na55\nanal\nanus\nar5e\narrse\narse\nass\nass-fucker\nasses\nassfucker\nassfukka\nasshole\nassholes\nasswhole\na_s_s\nb!tch\nb00bs\nb17ch\nb1tch\nballbag\nballs\nballsack\nbastard\nbeastial\nbeastiality\nbellend\nbestial\nbestiality\nbi+ch\nbiatch\nbitch\nbitcher\nbitchers\nbitches\nbitchin\nbitching\nbloody\nblow job\nblowjob\nblowjobs\nboiolas\nbollock\nbollok\nboner\nboob\nboobs\nbooobs\nboooobs\nbooooobs\nbooooooobs\nbreasts\nbuceta\nbugger\nbum\nbunny fucker\nbutt\nbutthole\nbuttmunch\nbuttplug\nc0ck\nc0cksucker\ncarpet muncher\ncawk\nchink\ncipa\ncl1t\nclit\nclitoris\nclits\ncnut\ncock\ncock-sucker\ncockface\ncockhead\ncockmunch\ncockmuncher\ncocks\ncocksuck \ncocksucked \ncocksucker\ncocksucking\ncocksucks \ncocksuka\ncocksukka\ncok\ncokmuncher\ncoksucka\ncoon\ncox\ncrap\ncum\ncummer\ncumming\ncums\ncumshot\ncunilingus\ncunillingus\ncunnilingus\ncunt\ncuntlick \ncuntlicker \ncuntlicking \ncunts\ncyalis\ncyberfuc\ncyberfuck \ncyberfucked \ncyberfucker\ncyberfuckers\ncyberfucking \nd1ck\ndamn\ndick\ndickhead\ndildo\ndildos\ndink\ndinks\ndirsa\ndlck\ndog-fucker\ndoggin\ndogging\ndonkeyribber\ndoosh\nduche\ndyke\nejaculate\nejaculated\nejaculates \nejaculating \nejaculatings\nejaculation\nejakulate\nf u c k\nf u c k e r\nf4nny\nfag\nfagging\nfaggitt\nfaggot\nfaggs\nfagot\nfagots\nfags\nfanny\nfannyflaps\nfannyfucker\nfanyy\nfatass\nfcuk\nfcuker\nfcuking\nfeck\nfecker\nfelching\nfellate\nfellatio\nfingerfuck \nfingerfucked \nfingerfucker \nfingerfuckers\nfingerfucking \nfingerfucks \nfistfuck\nfistfucked \nfistfucker \nfistfuckers \nfistfucking \nfistfuckings \nfistfucks \nflange\nfook\nfooker\nfuck\nfucka\nfucked\nfucker\nfuckers\nfuckhead\nfuckheads\nfuckin\nfucking\nfuckings\nfuckingshitmotherfucker\nfuckme \nfucks\nfuckwhit\nfuckwit\nfudge packer\nfudgepacker\nfuk\nfuker\nfukker\nfukkin\nfuks\nfukwhit\nfukwit\nfux\nfux0r\nf_u_c_k\ngangbang\ngangbanged \ngangbangs \ngaylord\ngaysex\ngoatse\nGod\ngod-dam\ngod-damned\ngoddamn\ngoddamned\nhardcoresex \nhell\nheshe\nhoar\nhoare\nhoer\nhomo\nhore\nhorniest\nhorny\nhotsex\njack-off \njackoff\njap\njerk-off \njism\njiz \njizm \njizz\nkawk\nknob\nknobead\nknobed\nknobend\nknobhead\nknobjocky\nknobjokey\nkock\nkondum\nkondums\nkum\nkummer\nkumming\nkums\nkunilingus\nl3i+ch\nl3itch\nlabia\nlmfao\nlust\nlusting\nm0f0\nm0fo\nm45terbate\nma5terb8\nma5terbate\nmasochist\nmaster-bate\nmasterb8\nmasterbat*\nmasterbat3\nmasterbate\nmasterbation\nmasterbations\nmasturbate\nmo-fo\nmof0\nmofo\nmothafuck\nmothafucka\nmothafuckas\nmothafuckaz\nmothafucked \nmothafucker\nmothafuckers\nmothafuckin\nmothafucking \nmothafuckings\nmothafucks\nmother fucker\nmotherfuck\nmotherfucked\nmotherfucker\nmotherfuckers\nmotherfuckin\nmotherfucking\nmotherfuckings\nmotherfuckka\nmotherfucks\nmuff\nmutha\nmuthafecker\nmuthafuckker\nmuther\nmutherfucker\nn1gga\nn1gger\nnazi\nnigg3r\nnigg4h\nnigga\nniggah\nniggas\nniggaz\nnigger\nniggers \nnob\nnob jokey\nnobhead\nnobjocky\nnobjokey\nnumbnuts\nnutsack\norgasim \norgasims \norgasm\norgasms \np0rn\npawn\npecker\npenis\npenisfucker\nphonesex\nphuck\nphuk\nphuked\nphuking\nphukked\nphukking\nphuks\nphuq\npigfucker\npimpis\npiss\npissed\npisser\npissers\npisses \npissflaps\npissin \npissing\npissoff \npoop\nporn\nporno\npornography\npornos\nprick\npricks \npron\npube\npusse\npussi\npussies\npussy\npussys \nrectum\nretard\nrimjaw\nrimming\ns hit\ns.o.b.\nsadist\nschlong\nscrewing\nscroat\nscrote\nscrotum\nsemen\nsex\nsh!+\nsh!t\nsh1t\nshag\nshagger\nshaggin\nshagging\nshemale\nshi+\nshit\nshitdick\nshite\nshited\nshitey\nshitfuck\nshitfull\nshithead\nshiting\nshitings\nshits\nshitted\nshitter\nshitters \nshitting\nshittings\nshitty \nskank\nslut\nsluts\nsmegma\nsmut\nsnatch\nson-of-a-bitch\nspac\nspunk\ns_h_i_t\nt1tt1e5\nt1tties\nteets\nteez\ntestical\ntesticle\ntit\ntitfuck\ntits\ntitt\ntittie5\ntittiefucker\ntitties\ntittyfuck\ntittywank\ntitwank\ntosser\nturd\ntw4t\ntwat\ntwathead\ntwatty\ntwunt\ntwunter\nv14gra\nv1gra\nvagina\nviagra\nvulva\nw00se\nwang\nwank\nwanker\nwanky\nwhoar\nwhore\nwillies\nwilly\nxrated\nxxx\n"
	replaceAsterisk = "*"
)

func init() {
	_profanityWords = profanityWord{words: make(map[string]bool, 1000)}
	words := strings.Split(wordsContent, "\n")
	for _, word := range words {
		_profanityWords.words[word] = true
	}
}

func (filter *profanityWord) asteriskWords(input string) string {
	begin, curr := -1, 0
	var buffer bytes.Buffer
	var strLen = len(input)
	for curr < strLen {
		// search begin of word
		for curr < strLen {
			if unicode.IsSpace(rune(input[curr])) {
				buffer.WriteByte(input[curr])
				curr++
			} else {
				begin = curr
				curr++
				break
			}
		}

		// loop over
		if begin < 0 {
			break
		}

		for curr < strLen {
			if unicode.IsSpace(rune(input[curr])) {
				break
			} else {
				curr++
			}
		}

		// word
		word := input[begin:curr]
		if _, ok := filter.words[word]; ok {
			var wl = len(word)
			for l := 0; l < wl; l++ {
				buffer.WriteString(replaceAsterisk)
			}
		} else {
			buffer.WriteString(word)
		}
		begin = -1
	}
	return buffer.String()
}
