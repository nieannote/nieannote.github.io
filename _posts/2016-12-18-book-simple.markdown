---
layout: post
title: 读书笔记之《极简设计》
date: 2016-12-18 16:50
tags:
  - book
---

**主流用户是上帝**。借助交互设计的四策略 —— 删除、组织、隐藏、转移，可以达到简约设计的目的。

本书出自实践，语言干练，要点明确，例子较多、易于理解，是不可多得的佳作。这篇读书笔记以原书大纲为主，同时摘录了书中的一些要点，希望对大家有所帮助。本文主要借鉴了[董佳奇](http://qinsman.com/resume/)的[读书笔记](http://qinsman.com/reading/jyzs.html)，在原文基础上做了一些调整。

![front.jpg](https://raw.githubusercontent.com/nieannote/nieannote.github.io/master/images/20161218/front.jpg)


## 1 话说简单
+ 简单的威力: 人们喜欢简单、值得信赖、适应性强的产品
+ 复杂的产品不可持续
	- 增加的功能越多，就越难发现真正对用户有价值的新功能。这样盲目添加的新功能早晚会成为垃圾功能。增加复杂性意味着遗留代码越来越沉重，导致产品维护成本越来越高，而且也越来越难以灵活应对市场变化
	- 用户也会对你的产品越来越不满意。因为增加的复杂性导致他们很难找到自己真正需要的功能
	- 况且，一想到为那么多没用的功能买单，他们就更加不高兴了。所有不必要的功能都是要付钱的！
+ 不是那种简单法: 在做技术产品的设计时，至少有3个角度: 管理人员、工程师和用户。而我们要讨论的是怎么让用户感觉用起来简单
+ 简单的特征
	- 简单并不意味着欠缺或低劣，也不意味着不注重装饰或者完全赤裸裸。而是说装饰应该紧密贴近设计本身，任何无关的要素都应该予以剔除
	- 简单并不意味着最少化。朴素的设计仍然具有自身的特征和个性
	- 简单的特征和个性应该源自你使用的方法、所要表现的产品，以及用户执行的任务
+ 貌似简单
	- 貌似简单的东西没有一个能够应验的。相反，它们的存在反而会让事情变得更复杂，效果更差
	- 例子1——操作向导: 把一件事分成几个步骤，企图以此达到简化的目的。但实际上，向导意味着剥夺了用户的控制权，所以，人们才会感觉受到了向导的限制
	- 例子2——让一个有魔力的角色来预测用户需求，并告诉用户该怎么做
+ 了解你自己: 你要搞清楚简化用户体验将会如何影响方程式中的每一项。以汽车公司为例: （汽车销售量）×（汽车单价）－（成本）＝（利润）


## 2 明确认识
+ 描述要点的两种方式
	- 无论是设计整个Web站点还是设计一个下拉菜单，都需要对什么是简单的体验有一个认识
	- 简单而迅速的方式: 用一句话把它写出来，包括我要设计什么、要遵循哪几条设计规则，尽量使用最简单的术语
	- 先理解用户，再思考设计: 更好而花费时间更长的方式是描述我希望用户拥有什么体验。具体一点儿说，就是描述用户的使用情景，以及我的设计怎么满足用户在该情景下的需求
+ 走出办公室
	- 软件会用环境是观察用户的最佳地点: 先到用户要使用你的软件的地方去做个调查。很少有用户是在这种安静的环境下使用软件的。用户体验是否简约必须要在纷乱、多变的环境中才能考察出来
	- 如果你不能到用户现场，那么就要跟用户多了解一些他们工作环境的情况，特别是要知道他们在使用你的软件的时候经常会发生什么事情
	- 无法控制用户使用软件的环境，而只能使软件设计符合环境需求
+ 观察什么: 在家里、在公司、在户外，你的设计必须能够适应各种干扰
+ **三种用户**
	- **专家型用户**: 愿意探索你的产品或服务，并且会给你提出各种改进建议
	- **随意型用户**: 他们可能使用过类似的产品或服务。他们有兴趣使用更高级更复杂的产品，但却不愿意接触全新的东西
	- **主流用户**: 最大的一个用户群体，他们自己不会因为你的技术而使用你的产品，使用你产品的目的是完成某项任务
	- 即便一个产品用了很多年，用户类型的标签也是不会变的
	- 主流用户占绝对的主体地位，专家型和随意型用户只是少数派。针对前两种类型的用户设计产品或许更有诱惑力（他们更识货），不过感觉简单的体验却是主流用户所喜爱的
+ 为什么应该忽略专家型用户和随意型用户
	- 专家并不是典型用户，他们的判断会出现偏差。他们不会体验到主流用户遇到的问题。他们追求主流用户根本不在乎的功能。专家想要的功能往往会吓倒主流用户
	- 随意型用户还不够典型。这类人数量有限，相对极端，他们的技术水平较好，而且比主流用户更具有忍耐力
+ 要为主流用户而设计: 简单的用户体验是初学者、新手的体验，或者是压力之下的主流用户的体验。想要吸引大众，必须关注主流
+ 感情需求: 理解感情需求能够帮你把握设计重点
+ 简单意味着控制
	- 从简单这个角度来看，最重要的是让用户感到自己在掌控一切
	- 首先，用户希望感觉是在掌控自己使用的技术。主流用户希望自己掌控起来容易、可靠、迅速
	- 其次，用户希望感觉是在掌控自己的生活
+ 正确选择"什么"
	- 正确的做法是只描述一定程度的细节，能说清楚就行了
	- 关键是不能遗漏用户体验过程中的任何一个步骤
	- 要保证通过用户的语言来描述行动的经过
+ 描述用户体验: 故事应该用三言两语把核心体验表达出来
+ 讲故事
	- 不要担心故事的表现形式，关键是把你的所有约束条件诉诸文字
	- 要通过一个小故事展示出每一个需求点，并确定满足该需求的功能（核心功能）
	- 展示，而不是讲
	- 要想可信必须以真人真事为原型
	- 好的用户故事应该简明、具体、可信，并且拥有相关细节
+ 环境、角色、情节
	- 可信的环境（故事中的"时间"和"地点"）
	- 可信的角色（"谁"和"为什么"）
	- 流畅的情节（"什么"和"怎么样"）
+ 极端的可用性
	- 什么是简单的体验: 能够适应极端条件
	- 争取你不可能达成的目标有一个重要的好处: 保持正确的方向
	- 如果你设定了一个极端的目标，你的产品就能随着时间推移越变越好（至少能够实现真正重要的目标）
+ 简便的方式
	- 我的目标是拿出一个简洁、清晰、完整的描述
	- 我甚至想只用一句简短的话来表述。如果这句话既能忽略细节而概括出主要活动，又能不让听众失去兴趣，那么就说明它已经达到了简洁的标准
+ 洞察力
	- 首先，回顾一下你从用户那里收集的素材、他们面对的问题、他们生活的世界。把那些对用户行为影响最大的事情放在前面。然后，从你的故事中寻找突破口，把这些设计要点按先后次序排列出来。最后，验证你的见解
	- 验证你的想法意味着还要花更多时间观察现实中的人，通常可以使用原型或者竞争性产品作为辅助
+ 明确认识
	- 太早开始设计意味着会遗漏重要的见解，甚至意味着设计思路完全错误
	- 乍一看到某个问题，你会觉得很简单，其实你并没有理解其复杂性。当你把问题搞清楚之后，又会发现真的很复杂，于是你就拿出一套复杂的方案来。实际上，你的工作只做了一半，大多数人也都会到此为止... 但是，真正伟大的人还会继续向前，直至找到问题的关键和深层次原因，然后再拿出一个优雅的、堪称完美的有效方案
+ 分享
	- 让最核心的理念随处可见，提醒人们时刻谨记
	- 与别人分享你的认识，即使你不在场也能保证作出正确的决定
	- 跟参与项目的每一个人复述你的故事，看见他们一次就讲一次。不要停下来，要天天讲，反复讲。直到你讲得自己都厌烦了，人们才会真正领悟你的认识


## 3 简约四策略
四个策略，以遥控器设计为例，描述如下:

+ 删除: 去掉所有不必要的按钮，直至减到不能再减
+ 组织: 按照有意义的标准将按钮划分成组
+ 隐藏: 把那些不是最重要的按钮安排在活动仓盖之下，避免分散用户注意力
+ 转移: 只在遥控器上保留具备最基本功能的按钮，将其他控制转移到电视屏幕上的菜单里，从而将复杂性从遥控器转移到电视


## 4 删除
+ 删除: 简化设计最明显的方式就是删除不必要的功能
+ 避免错删: 删除功能时要避免错删，而把一切难于实现的功能统统抹杀就是典型的错误做法。不要等着别人不分青红皂白地、无情地删除最有意思的功能。要总揽全局，保证只交付那些真正有价值的功能和内容
+ 关注核心: 在按照优先级对功能排序时，要时刻记住用户认为那些关系到他们日常使用体验的功能最有价值。另外，能够消除他们挫折感的功能同样也会受到欢迎。在描绘用户故事时，别忘了寻找常见的挫折和难题。解决这些问题的功能的优先级次之
+ 砍掉残缺功能: 删掉实现得不够理想的功能也是很重要的。问题绝非"为什么应该去掉它"，而是"为什么要留着它"
+ "假如用户想..."
	- "假如用户想..."只会刺激人们求全的心理
	- 搞清楚这个功能对用户是否真的重要。问一问: "我的目标用户经常会遇到这个问题吗?"
+ "但我们的用户想要..."
	- 对用户的要求做**逆向工程——搞清楚用户到底遇到了什么问题，仔细斟酌这个问题是不是应该由我们的软件来解决**
	- 增加功能不一定会让用户体验更简单，反而经常会导致更多的迷惑
	- 很多情况下，你都有可能拿出一个能够满足用户真正需求的替代方案（例如让他们在应用之间快速切换）
+ 方案，不是流程: 如果一个小的变化导致了复杂的流程，就应该退一步去寻找更好的解决方案——如果在设计的时候只盯住流程，那么结果很可能会创造更多的功能去处理出现的各种异常情况、问题和细节。要想避免这些复杂性，退一步想，把注意力集中到客户的目的上面，问自己"还有其他的解决方式吗?"
+ 如果功能不是必要的
	- 功能多对于没有机会试用的消费者有吸引力。但是，在消费者使用了产品之后，他们的偏好就会改变。一下子从重视功能，变成了更重视可用性
	- 你设计的产品如果承载过多的功能，更有可能降低主流用户的满意度，从而对产品的长期盈利能力造成损害
+ 删除功能真的有影响吗
	- 当用户离不开这项功能时，你再把它砍掉，即使是一个很不起眼的变化，都会激怒用户
	- 判断删除功能对用户的影响有多大可是一件需要技巧的事。删除一项功能对不同用户的影响不同
	- 要知道人们真正关心什么，探知他们对删掉某个功能后的产品有什么意见，最好的方法就是先做个模型出来让他们试用
+ 排定功能优先级
	- 确定用户想要达到的目的，并排定优先次序
	- 专注于寻找能够完全满足优先级最高的用户需求的解决方案。找到之后再考虑满足用户的其他目标
	- 确定用户在使用产品过程中最常见的干扰源，并将解决这些问题的功能按难易程度排出优先次序
	- 要知道能够满足主流用户的"足够好"的遥控器与只有专家才看得上眼的"精准的"遥控器有什么区别
	- 给那些轻易就能满足主流用户需求的功能排定优先次序
+ 负担: 界面中的各种小细节会增加用户的负担。可考虑的减负措施包括: 删除没人会看的文字，简化布局，去掉重复的链接，精简按钮和链接的样式，减少广告位和广告数量，去掉分散注意力的元素
+ 决策
	- 在为用户提供少量选择的情况下，用户购买的可能性要大于为他们提供大量选择的情况。在选择少的情况下，用户购买之后的满意度要高于选择多的情况
	- 给用户提供选择会让人感觉自己在把控着局面，而在某些情况下人们更愿意少一些选择。如果选择超过了一定的界限，特别是在很多选择都相似的情况下，选择反而变成了负担
+ 分心
	- 增加文档中的超链接会降低读者的理解力——即使读者不会打开链接也一样
	- 网页的右边栏经常会出现更多分散注意力的链接
	- 放置这些链接最好的地方是页面底部
+ 聪明的默认值
	- 选择聪明的默认值可以减少用户的选择
	- 聪明的默认值指的是适合大多数人口味的选择。通过分析客户信息（如日志文件），可以找到很多选择默认值的依据
	- 当一个客户再次光顾网站或应用，他通常愿意以前次离开的状态作为起点
	- 默认值是节省用户时间和精力的有效方式，也是清除设计蓝图中"减速带"的首选方式
+ 选项和首选项: 主流用户不喜欢为设置选项和首选项费心劳神
	- 在设计团队模棱两可的时候，选项和首选项比较容易泛滥
	- 简单的用户体验不会强迫用户去做这种选择，哪种方式最有效应该是设计团队考虑的问题。解决这个问题的最佳途径就是请一些用户来测试
+ 如果一个选项还嫌多
	- 在向用户提供选择时，务必要考虑周全一些，想清楚用户会不会因为这些选项而不知所措，或者这些选项会不会动摇他们的决心。例如，在结账这个环节，网售商删除了其他每个页面顶部和底部都有的导航链接
	- 向用户提供这些选项会不会因为追求完美而牺牲速度和简单？如果是，删除那些选项
+ 错误: 在设计简单的体验时，关键的一步是确定哪些地方需要错误消息，或者检查错误日志，从中找出常见的错误消息。消除错误的来源是简化体验的一个重要思路
+ 视觉混乱: 删除混乱元素很简单。观察设计方案中的每一个元素，想一想为什么需要它
	- 使用空白或轻微的背景色来划分页面，而不要使用线条
	- 尽可能少使用强调，别使用粗黑线，匀称、浅色的线更好
	- 控制信息的层次。如果页面中信息的层次超过了两或三个层次，就会导致用户迷惑
	- 减少元素大小、开头的变化
+ 删减文字
	- 删除文字有下列三大好处——①重要的内容"水落石出"；②消除了分析满屏内容的麻烦；③读者会对自己看到了什么更有自信
	- 多余文字常见的几个藏身之所: ①删除引见性文字；②删除不必要的说明；③删除烦琐的解释；④使用描述性链接（把标题本身作为链接可以让页面更清爽）
	- "把每一页中的文字删掉一半，然后再把剩下的再删掉一半"——《点石成金》第三条可用性法则
+ 精简句子: 几乎任何句子都能够精简
	- 不使用介词（"对于/根据/为了/基于/通过/关于"）。这些词会弱化句子的谓语，因此要尽量省略
	- 不使用is的动词形式（"正在消耗时间"），尽你所能使用其他表述方式（"花时间"）
	- 把被动句式（"时间是被这个项目所需要的"）转换为主动句式（"这个项目需要时间"）
	- 删掉索然无味的开头（"大家都很容易看到这一点，..."），开门见山
	- 减少废话。在表达相同意思的前提下，用"每天"代替"在每天的基础上"
+ 删减过多: 避免删减过多，关键在于让人们能够控制结果
	- 你能做到——说服干系人删除内容和功能: 在共同愿景的基础上，在关注主流用户的前提下，通过彻底重新设计是可以达到简约之效的
+ 焦点
	- 聚焦于对用户有价值的功能。这意味着专注于那些承载用户核心体验的功能，也意味着交付的功能必须能够消除用户的挫折感，能够消除他们的焦虑
	- 聚焦于可用资源，通过删除残缺的功能、不切题的元素和花里胡哨的东西为用户提供价值
	- 聚焦于达成用户的目标。纠结于流程会陷入细节的泥潭而无法自拔
	- 删除那些干扰性的、增加用户负担的"减速带": 错误消息、不知所云的文字、不必要的选项和造成视觉混乱的元素
	- 两个例外: 一个是不可避免的法律要件，必须包含特定的内容或信息。另一个例外是不能脱离环境删除某些功能



## 5 组织
+ 分块
	- 用户界面设计离不开分块
	- 有关分块的经典建议是把项组织到"7加减1"个块中。理论上讲，这个数字是人的大脑瞬间能够记住的最大数目
	- 应该尽可能少分几个块，这样才能让主流用户感觉更简单——分块越少，选择越少，用户负担越轻
+ 围绕行为进行组织
	- 着手组织之前首先要理解用户的行为: 他们想做什么，先做什么后做什么
	- 人们一般都希望按照某种特定的步骤做事。打乱这个步骤就会造成迷惑，令人沮丧
+ 是非分明
	- 在对一组性质相同的产品（如网上书店的书）进行分类时，确定清晰的分类标准对用户非常重要
	- 太多的重叠会导致困惑，但有时候确实无法避免。所谓最简单的分类，通常指的是重复交叉最少的分类方法
+ 字母表与格式
	- 按照字母表顺序排列，其实会把顺序搞乱。按照字母表顺序排列看起来简单，却经常不可行
	- 对于专有名词，按照字母表顺序建立索引是没有问题的，因为有一个"准确无误"的词在描述某个概念，
	- 按照格式（文字、图片、视频）来对内容进行排序，是另一种看起来简单实则费力不讨好的分类方法
+ 搜索
	- 有论调称有的用户认为搜索比浏览更容易，但测试表明没有一个人始终会把搜索作为第一选择。当然也有例外——设想用户需要从大量类似项中挑选一个已知项，比如在iTunes提供的成千上万条可下载音乐中找到某个特殊曲目
	- 浏览的一个不太为人所知的好处在于，当人们看到站点中的主链接，或者界面中的主控件时，他们能够理解当前程序可以做什么
	- 设计简单的搜索界面其实要困难得多。你必须考虑搜索关键词中的拼写错误和同义词，而且，还需要对搜索结果有效地分类组织
+ 时间和空间
	- 按照时间来组织活动是一种简单又通用的方式。对于那些持续时间相差不大的活动，按照时间排序是最合适的
	- 一些实体对象，如酒店和国家之类的，全都可以按照空间来组织，只要用户对排列方式不感到陌生即可
	- 通过图解形式来表示时间和空间可能会有一些问题——密度不均。能够形象地反映密集程度的变化很重要
+ 网格
	- 利用不可见的网格来对齐界面元素，是吸引用户注意力的一种有效方式。网格越是简单，效果就越明显
	- 哪怕少数几个元素没有放到位，都会破坏这种网格布局的引导效果
	- 网格布局也会让人感觉局促和受压制。要解决这个问题，可以设计一个不对称的布局
+ 大小和位置: 在利用网格来布局界面项时，请注意参考如下关于大小和位置的提示
	- 重要的元素要大一些，即便比例失调也可以考虑
	- 不太重要的界面元素应该小一些。要想办法表现出不同的重要性，否则用户就会被搞迷糊。记住这条规则: 如果一个元素的重要性为1/2，那就把它的大小做成1/4
	- 把相似的元素放在一起
	- 从来没有发现任何证据能够证明导航条靠在屏幕的左边或右边，甚至横跨在屏幕上方的效果肯定更好。真正应该担心的问题，是用户能不能轻易地找到想找的按钮
	- 触摸屏上，把导航放在左侧（或右侧）可能会给习惯用右（左）手的人带来麻烦
+ 分层
	- 利用感知分层技术，我们可以把一些元素放在另一些元素上方，或者把两组元素并排起来。甚至，还可以让散落在用户界面各个地方的元素之间建立联系
	- 感知分层借助于颜色很容易实现。除了颜色之外，使用灰色阴影、大小缩放，甚至形状变化，都可以实现感知分层
	- 关于分层的提示: 尽可能使用较少的层；考虑把某些基本元素放在常规背景层；对于同等重要的类别，使用明亮、高饱和度的颜色可以让它们在页面上更加突出；对于同等重要的类别，利用感知分层技术，使用相同的亮度和大小，只是色调有所区别
+ 色标
	- 分层信息中的颜色利用了人们的记忆原理，因此给人造成的负担很轻。而使用颜色来标记信息的代价却很明显: 与任何标记系统一样，需要人们花时间来学习和理解这些标记，因此需要用户花费更多心思
	- 临时的访客也许没有时间去学习和记忆
	- 在某个环节中众所周知的标记系统借用到其他环节时可能会造成问题
	- 在不必要的情况下添加颜色会导致困惑
	- 在确保人们会花很长时间学习，且他们会重复使用你的设计时，色标系统非常适合。当然，使用人们已经知道含义的色标也没问题
+ 期望路径: 描述用户使用软件的路径时，千万不要被自己规划图中清晰的线条和整洁的布局缩迷惑。人们并不总是走你铺好的路



## 6 隐藏
+ 隐藏的优势: 用户不会因不常用的功能分散注意力。但无论隐藏什么功能，都意味着你在用户和功能之间设置了一道障碍，必须仔细权衡要隐藏哪些功能
+ 不常用但不能少
	- 那些主流用户很少使用，但自身需要更新的功能，通常是适合隐藏的功能: 事关细节的、选项和偏好、特定于地区的信息
	- 最好把它们放在一个开放的页面，或者所有页面中，并且最好隐藏在网站的开始位置，或者应用程序的边缘
+ 自定义
	- 一般来说，不应该让用户自定义他们的软件。让用户根据自己的需求自定义界面显得设计人员懒惰、没有主见。自定义可能是一件非常耗费时间、令人讨厌的事
	- 让用户自定义自己的用户界面是假设用户知道如何布局最有效、最高效，如果用户只需添加几项即可完成自定义，如果不需要重排那么多，自定义还是很有价值的
+ 自动定制
	- 有些程序会根据用户的行为自动显示或隐藏某些功能，这种自动定制不会让界面变得简单，反而会把界面搞得很复杂，给用户带来极大不便
	- 缺点: 很难保证默认菜单的准确性；缩短菜单后，用户需要把每个功能看两变才能确定（先看短菜单，再看长菜单）；用户最终不知道去哪里找自己享用的命令
+ 渐进展示: 通常一项功能会包含少数核心的供主流用户使用的控制部件，另有一些专为专家级用户准备的扩展性精确控制部件。隐藏这些精确的控制部件是保持设计简单的不错选择。这种"核心功能+扩展功能"的模式，不仅能够简化设计，更是一种强大的交互手段
+ 阶段展示: 如果所有用户都会随着搜索的深入而寻找较为复杂的功能，就可以使用阶段展示（如登记表单通常都需要使用阶段展示，向导也是一种阶段展示的形式），但要遵循几条规则: 
	- 设定一种场景
	- 讲一个故事——用户希望每个环节都能像讲故事一样层层展开
	- 说用户的语言
	- 把信息分成小块展示
+ 适时出现: 首先，尽可能彻底地隐藏所有需要隐藏的功能。其次，只在合适的时机、合适的位置上显示相应的功能。例如，在选择单词后才显示词典功能
+ 提示和线索
	- 隐藏复杂性的一个原因，就是不想让用户产生自己什么都不懂的感觉，而为按钮打上"高级"的标签，显然就是在呵斥用户
	- 可以只对特定的用户使用一个标签
	- 优秀的例子——PS和AI工具箱的小三角形按住后可看到高级选项: 采用了应邀探索设计模式，而非一个试图介绍更多功能的标签，能让用户清楚地知道高级工具与基本工具能完成类似的任务
隐藏处理得好的界面会给人一种优雅的感觉: 界面中包含的线索尽管细微，却能恰到好处地提示出隐藏功能的位置和功用
+ 让功能容易找到: 把标签放在哪里比把标签做多大重要得多，即使是一个非常小的标签，只要放在了用户关注点上，也能收到良好的效果
+ 隐藏策略的要求: 
	- 隐藏一次性设计和选项
	- 隐藏精确控制选项，但专家用户必须能够让这些选项始终保持可见
	- 不可强迫或寄希望于主流用户使用自定义功能，不过可以给专家提供这个选项
	- 巧妙地隐藏——首先是彻底隐藏，其次是适时出现
	- 只要不让人找太久，隐藏就是有效的



## 7 转移
+ 转移策略: 把正确的功能放到正确的平台或者正确的系统组件上去
	- 在设备之间转移——移动平台与桌面平台: 任何设备都有自己的长处和不足，把某项任务的某些部分转移到不同平台上可能是一种更好的选择
	- 向用户转移: 简化界面，而将模棱两可信息的界定这类复杂的工作转移到每一位用户的头脑中
	- 用户最擅长做什么: 让用户和计算机各自去做最擅长的事
+ 创造开放式体验: 让一个组件、功能具有多种用途也是一种简化之道。找一个功能总比在几个类似的功能中选择容易，学习一个功能也比学习多个功能容易，且一个功能更容易维护
+ 简单界面的最高境界: 专家和主流用户都会感觉它非常好用
	- 专家和主流用户可以分别设置自己不同的目标
	- 开放性界面的秘诀: 尽量减少仅适合中级用户的"便捷"特性
+ 非结构化数据: 计算机应当有能力识别并将用户提供的数据结构化，例如用户可以用任意格式和人类语言来写邮件，计算机负责发现邮件中是否有需要结构化或进行后续操作的数据
+ 信任: 简单的体验需要信任，计算机之所以搞得用户不舒服，就是因为它们总是控制和指挥用户
	- 如果把一组任务分解为两部分，分别交给两个设备来完成，而且这两个设备必须以某种特定的方式配合使用，那么这种情况下最容易实现任务的转移
	- 在难以分清设备之间如何协同工作时，要实现功能的转移是比较困难的
	- 如果想把任务转移到用户一方，你必须相信用户有能力完成该任务
	- 构筑信任关系的唯一方式: 让用户参与测试原型或实物模型


## 8 最后的叮嘱
+ 顽固的复杂性
	- **任何应用程序都会有一些无法消除的复杂性，关键问题在于，由谁来面对这些复杂性**
	- **设计简单用户体验的最后，往往不是问"怎样才能把这个功能设计得简单"，而是问"到底应该把这个复杂功能放到哪里"**
+ 细节: 设计中微小的瑕疵都可能变成永远挥之不去的烦恼。花上半天时间重新设计一个解决方案，解决看似微不足道的小问题，也许就能把成千上万次投诉消弭于无形
+ 简单发生在用户的头脑中: **不要让你的设计干扰用户的思绪**，简单的设计能够为用户留出足够空间，他们会用自己的生活来填充这些空间，从而创造出更丰富、更有意义的体验



谨以此文，献给那些 "坚持极简设计、坚持以主流用户为本" 的同仁们!
