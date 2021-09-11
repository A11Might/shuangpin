package util

var (
	naturalCodeKeyToPinyin map[string][]string
	naturalCodePinyinToKey map[string]string
	initials               []string
	jsonData               string
)

func getNaturalCodeKeyToPinyin() map[string][]string {
	return naturalCodeKeyToPinyin
}

func getNaturalCodePinyinToKey() map[string]string {
	return naturalCodePinyinToKey
}

func getInitials() []string {
	return initials
}

func getJsonDate() string {
	return jsonData
}

func init() {
	naturalCodeKeyToPinyin = map[string][]string{
		"q": {"q", "iu"},
		"w": {"w", "ia", "ua"},
		"e": {"e"},
		"r": {"r", "uan"},
		"t": {"t", "ue", "ve"},
		"y": {"y", "ing", "uai"},
		"u": {"sh", "u"},
		"i": {"ch", "i"},
		"o": {"o", "uo"},
		"p": {"p", "un"},

		"a": {"a"},
		"s": {"s", "iong", "ong"},
		"d": {"d", "iang", "uang"},
		"f": {"f", "en"},
		"g": {"g", "eng"},
		"h": {"h", "ang"},
		"j": {"j", "an"},
		"k": {"k", "ao"},
		"l": {"l", "ai"},

		"z": {"z", "ei"},
		"x": {"x", "ie"},
		"c": {"c", "iao"},
		"v": {"zh", "ui", "v"},
		"b": {"b", "ou"},
		"n": {"n", "in"},
		"m": {"m", "ian"},
	}

	naturalCodePinyinToKey = map[string]string{
		"q":   "q",
		"iu":  "q",
		"w":   "w",
		"ia":  "w",
		"ua":  "w",
		"e":   "e",
		"r":   "r",
		"uan": "r",
		"t":   "t",
		"ue":  "t",
		"ve":  "t",
		"y":   "y",
		"ing": "y",
		"uai": "y",
		"sh":  "u",
		"u":   "u",
		"ch":  "i",
		"i":   "i",
		"o":   "o",
		"uo":  "o",
		"p":   "p",
		"un":  "p",

		"a":    "a",
		"s":    "s",
		"iong": "s",
		"ong":  "s",
		"d":    "d",
		"iang": "d",
		"uang": "d",
		"f":    "f",
		"en":   "f",
		"eng":  "g",
		"g":    "g",
		"h":    "h",
		"ang":  "h",
		"j":    "j",
		"an":   "j",
		"k":    "k",
		"ao":   "k",
		"l":    "l",
		"ai":   "l",

		"z":   "z",
		"ei":  "z",
		"x":   "x",
		"ie":  "x",
		"c":   "c",
		"iao": "c",
		"zh":  "v",
		"ui":  "v",
		"v":   "v",
		"b":   "b",
		"ou":  "b",
		"n":   "n",
		"in":  "n",
		"m":   "m",
		"ian": "m",
	}

	initials = []string{"b", "p", "m", "f", "d", "t", "n", "l", "g", "k", "h", "j", "q", "x", "zh", "ch", "sh", "r", "z", "c", "s", "y", "w"}

	jsonData = `{
  "title":"学生会退会",
  "famous":[
    "爱迪生a，天才是百分之一的勤奋加百分之九十九的汗水。b",
    "查尔斯·史a，一个人几乎可以在任何他怀有无限热忱的事情上成功。b",
    "培根说过，深窥自己的心，而后发觉一切的奇迹在你自己。b",
    "歌德曾经a，流水在碰到底处时才会释放活力。b",
    "莎士比亚a，那脑袋里的智慧，就像打火石里的火花一样，不去打它是不肯出来的。b",
    "戴尔·卡耐基a，多数人都拥有自己不了解的能力和机会，都有可能做到未曾梦想的事情。b",
    "白哲特a，坚强的信念能赢得强者的心，并使他们变得更坚强。b",
    "伏尔泰a, 不经巨大的困难，不会有伟大的事业。b",
    "富勒曾经a, 苦难磨炼一些人，也毁灭另一些人。b",
    "文森特·皮尔a, 改变你的想法，你就改变了自己的世界。b",
    "拿破仑·希尔a, 不要等待，时机永远不会恰到好处。b",
    "塞涅卡a, 生命如同寓言，其价值不在与长短，而在与内容。b",
    "奥普拉·温弗瑞a, 你相信什么，你就成为什么样的人。b",
    "吕凯特a, 生命不可能有两次，但许多人连一次也不善于度过。b",
    "莎士比亚a, 人的一生是短的，但如果卑劣地过这一生，就太长了。b",
    "笛卡儿a, 我的努力求学没有得到别的好处，只不过是愈来愈发觉自己的无知。b",
    "左拉a, 生活的道路一旦选定，就要勇敢地走到底，决不回头。b",
    "米歇潘a, 生命是一条艰险的峡谷，只有勇敢的人才能通过。b",
    "吉姆·罗恩a, 要么你主宰生活，要么你被生活主宰。b",
    "日本谚语a, 不幸可能成为通向幸福的桥梁。b",
    "海贝尔a, 人生就是学校。在那里，与其说好的教师是幸福，不如说好的教师是不幸。b",
    "杰纳勒尔·乔治·S·巴顿a, 接受挑战，就可以享受胜利的喜悦。b",
    "德谟克利特a, 节制使快乐增加并使享受加强。b",
    "裴斯泰洛齐a, 今天应做的事没有做，明天再早也是耽误了。b",
    "歌德a, 决定一个人的一生，以及整个命运的，只是一瞬之间。b",
    "卡耐基a, 一个不注意小事情的人，永远不会成就大事业。b",
    "卢梭a, 浪费时间是一桩大罪过。b",
    "康德a, 既然我已经踏上这条道路，那么，任何东西都不应妨碍我沿着这条路走下去。b",
    "克劳斯·莫瑟爵士a, 教育需要花费钱，而无知也是一样。b",
    "伏尔泰a, 坚持意志伟大的事业需要始终不渝的精神。b",
    "亚伯拉罕·林肯a, 你活了多少岁不算什么，重要的是你是如何度过这些岁月的。b",
    "韩非a, 内外相应，言行相称。b",
    "富兰克林a, 你热爱生命吗？那么别浪费时间，因为时间是组成生命的材料。b",
    "马尔顿a, 坚强的信心，能使平凡的人做出惊人的事业。b",
    "笛卡儿a, 读一切好书，就是和许多高尚的人谈话。b",
    "塞涅卡a, 真正的人生，只有在经过艰难卓绝的斗争之后才能实现。b",
    "易卜生a, 伟大的事业，需要决心，能力，组织和责任感。b",
    "歌德a, 没有人事先了解自己到底有多大的力量，直到他试过以后才知道。b",
    "达尔文a, 敢于浪费哪怕一个钟头时间的人，说明他还不懂得珍惜生命的全部价值。b",
    "佚名a, 感激每一个新的挑战，因为它会锻造你的意志和品格。b",
    "奥斯特洛夫斯基a, 共同的事业，共同的斗争，可以使人们产生忍受一切的力量。　b",
    "苏轼a, 古之立大事者，不惟有超世之才，亦必有坚忍不拔之志。b",
    "王阳明a, 故立志者，为学之心也；为学者，立志之事也。b",
    "歌德a, 读一本好书，就如同和一个高尚的人在交谈。b",
    "乌申斯基a, 学习是劳动，是充满思想的劳动。b",
    "别林斯基a, 好的书籍是最贵重的珍宝。b",
    "富兰克林a, 读书是易事，思索是难事，但两者缺一，便全无用处。b",
    "鲁巴金a, 读书是在别人思想的帮助下，建立起自己的思想。b",
    "培根a, 合理安排时间，就等于节约时间。b",
    "屠格涅夫a, 你想成为幸福的人吗？但愿你首先学会吃得起苦。b",
    "莎士比亚a, 抛弃时间的人，时间也抛弃他。b",
    "叔本华a, 普通人只想到如何度过时间，有才能的人设法利用时间。b",
    "博a, 一次失败，只是证明我们成功的决心还够坚强。 维b",
    "拉罗什夫科a, 取得成就时坚持不懈，要比遭到失败时顽强不屈更重要。b",
    "莎士比亚a, 人的一生是短的，但如果卑劣地过这一生，就太长了。b",
    "俾斯麦a, 失败是坚忍的最后考验。b",
    "池田大作a, 不要回避苦恼和困难，挺起身来向它挑战，进而克服它。b",
    "莎士比亚a, 那脑袋里的智慧，就像打火石里的火花一样，不去打它是不肯出来的。b",
    "希腊a, 最困难的事情就是认识自己。b",
    "黑塞a, 有勇气承担命运这才是英雄好汉。b",
    "非洲a, 最灵繁的人也看不见自己的背脊。b",
    "培根a, 阅读使人充实，会谈使人敏捷，写作使人精确。b",
    "斯宾诺莎a, 最大的骄傲于最大的自卑都表示心灵的最软弱无力。b",
    "西班牙a, 自知之明是最难得的知识。b",
    "塞内加a, 勇气通往天堂，怯懦通往地狱。b",
    "赫尔普斯a, 有时候读书是一种巧妙地避开思考的方法。b",
    "笛卡儿a, 阅读一切好书如同和过去最杰出的人谈话。b",
    "邓拓a, 越是没有本领的就越加自命不凡。b",
    "爱尔兰a, 越是无能的人，越喜欢挑剔别人的错儿。b",
    "老子a, 知人者智，自知者明。胜人者有力，自胜者强。b",
    "歌德a, 意志坚强的人能把世界放在手中像泥块一样任意揉捏。b",
    "迈克尔·F·斯特利a, 最具挑战性的挑战莫过于提升自我。b",
    "爱迪生a, 失败也是我需要的，它和成功对我一样有价值。b",
    "罗素·贝克a, 一个人即使已登上顶峰，也仍要自强不息。b",
    "马云a, 最大的挑战和突破在于用人，而用人最大的突破在于信任人。b",
    "雷锋a, 自己活着，就是为了使别人过得更美好。b",
    "布尔沃a, 要掌握书，莫被书掌握；要为生而读，莫为读而生。b",
    "培根a, 要知道对好事的称颂过于夸大，也会招来人们的反感轻蔑和嫉妒。b",
    "莫扎特a, 谁和我一样用功，谁就会和我一样成功。b",
    "马克思a, 一切节省，归根到底都归结为时间的节省。b",
    "莎士比亚a, 意志命运往往背道而驰，决心到最后会全部推倒。b",
    "卡莱尔a, 过去一切时代的精华尽在书中。b",
    "培根a, 深窥自己的心，而后发觉一切的奇迹在你自己。b",
    "罗曼·罗兰a, 只有把抱怨环境的心情，化为上进的力量，才是成功的保证。b",
    "孔子a, 知之者不如好之者，好之者不如乐之者。b",
    "达·芬奇a, 大胆和坚定的决心能够抵得上武器的精良。b",
    "叔本华a, 意志是一个强壮的盲人，倚靠在明眼的跛子肩上。b",
    "黑格尔a, 只有永远躺在泥坑里的人，才不会再掉进坑里。b",
    "普列姆昌德a, 希望的灯一旦熄灭，生活刹那间变成了一片黑暗。b",
    "维龙a, 要成功不需要什么特别的才能，只要把你能做的小事做得好就行了。b",
    "郭沫若a, 形成天才的决定因素应该是勤奋。b",
    "洛克a, 学到很多东西的诀窍，就是一下子不要学很多。b",
    "西班牙a, 自己的鞋子，自己知道紧在哪里。b",
    "拉罗什福科a, 我们唯一不会改正的缺点是软弱。b",
    "亚伯拉罕·林肯a, 我这个人走得很慢，但是我从不后退。b",
    "美华纳a, 勿问成功的秘诀为何，且尽全力做你应该做的事吧。b",
    "俾斯麦a, 对于不屈不挠的人来说，没有失败这回事。b",
    "阿卜·日·法拉兹a, 学问是异常珍贵的东西，从任何源泉吸收都不可耻。b",
    "白哲特a, 坚强的信念能赢得强者的心，并使他们变得更坚强。 b",
    "查尔斯·史考伯a, 一个人几乎可以在任何他怀有无限热忱的事情上成功。 b",
    "贝多芬a, 卓越的人一大优点是：在不利与艰难的遭遇里百折不饶。b",
    "莎士比亚a, 本来无望的事，大胆尝试，往往能成功。b",
    "卡耐基a, 我们若已接受最坏的，就再没有什么损失。b",
    "德国a, 只有在人群中间，才能认识自己。b",
    "史美尔斯a, 书籍把我们引入最美好的社会，使我们认识各个时代的伟大智者。b",
    "冯学峰a, 当一个人用工作去迎接光明，光明很快就会来照耀着他。b",
    "吉格·金克拉a, 如果你能做梦，你就能实现它。b"
  ],
  "bosh":[
    "现在, 解决x的问题, 是非常非常重要的. 所以, ",
    "我们不得不面对一个非常尴尬的事实, 那就是, ",
    "x的发生, 到底需要如何做到, 不x的发生, 又会如何产生. ",
    "而这些并不是完全重要, 更加重要的问题是, ",
    "x, 到底应该如何实现. ",
    "带着这些问题, 我们来审视一下x. ",
    "所谓x, 关键是x需要如何写. ",
    "我们一般认为, 抓住了问题的关键, 其他一切则会迎刃而解.",
    "问题的关键究竟为何? ",
    "x因何而发生?",
    "每个人都不得不面对这些问题.  在面对这种问题时, ",
    "一般来讲, 我们都必须务必慎重的考虑考虑. ",
    "要想清楚, x, 到底是一种怎么样的存在. ",
    "了解清楚x到底是一种怎么样的存在, 是解决一切问题的关键.",
    "就我个人来说, x对我的意义, 不能不说非常重大. ",
    "本人也是经过了深思熟虑,在每个日日夜夜思考这个问题. ",
    "x, 发生了会如何, 不发生又会如何. ",
    "在这种困难的抉择下, 本人思来想去, 寝食难安.",
    "生活中, 若x出现了, 我们就不得不考虑它出现了的事实. ",
    "这种事实对本人来说意义重大, 相信对这个世界也是有一定意义的.",
    "我们都知道, 只要有意义, 那么就必须慎重考虑.",
    "既然如此, ",
    "那么, ",
    "我认为, ",
    "一般来说, ",
    "总结的来说, ",
    "既然如何, ",
    "经过上述讨论, ",
    "这样看来, ",
    "从这个角度来看, ",
    "我们不妨可以这样来想: ",
    "这是不可避免的. ",
    "可是，即使是这样，x的出现仍然代表了一定的意义. ",
    "x似乎是一种巧合，但如果我们从一个更大的角度看待问题，这似乎是一种不可避免的事实. ",
    "在这种不可避免的冲突下，我们必须解决这个问题. ",
    "对我个人而言，x不仅仅是一个重大的事件，还可能会改变我的人生. "
  ],
  "after":[
    "这不禁令我深思. ",
    "带着这句话, 我们还要更加慎重的审视这个问题: ",
    "这启发了我. ",
    "我希望诸位也能好好地体会这句话. ",
    "这句话语虽然很短, 但令我浮想联翩. ",
    "这句话看似简单，但其中的阴郁不禁让人深思. ",
    "这句话把我们带到了一个新的维度去思考这个问题: ",
    "这似乎解答了我的疑惑. "
  ],
  "before":[
    "曾经说过",
    "在不经意间这样说过",
    "说过一句著名的话",
    "曾经提到过",
    "说过一句富有哲理的话"
  ]
}`
}
