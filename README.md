# 400kanji

## overview

1. acquire a sorted list of 400 kanji by frequency
2. get 10 sample words for each character
3. produce 3 versions of the word list
- english, kanji+hiragana
- kanji, english+hiragana
- audio, kanji+hiragana+english
4. ask chatgpt for a sample sentence for each word on the word list:
`i have a list of japanese words. for any given word, please generate a short sample sentence in japanese that uses this word. output the sample sentence in two saparate versions. the first is a kanji version. the second is a hiragana version. next give an english translation of this japanese sentence. repeat this task for each of these words. continue generating sentences until you have performed this task for every item on the list. the list follows. 期待, 招待,...`
5. produce multiple versions of the sentence list
- kanji, english+hiragana
- audio, kanji+hiragana+english
6. use shell scripts to generate the audio files via macOS `say` + `ffmpeg`: 
- `shell-scripts/kanji-audio_conversion.sh`
- `shell-scripts/sent-audio_conversion.sh`
7. move audio files into `collection.media`
8. use `400kanji` application to interleave the lists from above: 
- `go build`
- `mv 440kanji-builder ./merging`
- `cd ./merging`
- `./400kanji-builder`
9. use `split` to split the merged product into chunks: `shell-scripts/splitter.sh`
10. import each chunk into `anki`
11. export the whole back out from `anki` to an `.apkg`

## caveats

1. irregularities in the `chatgpt` output have not been smoothed over
2. `chatgpt` output; so, buyer beware
3. the number of sentences is out of sync with the number of kanji (*fixme...*)

## fyi

1070 distinct characters in the list of sample words

々いうえおかがきぎくけげこさしじすずせたちっつてでとなにねのはばびぶべまみむめもやょよらりるれろわんアェカグゴシテフプマムラリルレー一七万丈三上下不与世両並中丹主乗九了予争事二互五井亡交京人今介仕付代令以仲件任企伊会伝伸位低住佐体何作使例供価侵便係促保信俣修俳俵個倍倒候値倶停側備催債傾働像僚儀億償優元先光児党入全八公六共兵具典内円再冗写冬冷凍処出分切刊刑列初判別利到制券則削前副剰割創劇力功加助努労効勉動務勝募勢勤勧化北区医十千午半卓協南単博占印危却厚原去参友双反収取受叙口古可台史右号司各合同名后向告周味命和品員商問善喚喪営器四回因団困囲図国圏園土圧在地均型垣城域執基報場塁境増壌士声売変夕外多夜大天太夫央失奏契女好妥妻始委姿娯婚婦子字存季学宅守安完宗官定宝実客宣室宮害家容宿寄密富察審対寿専射将導小少就尿局居屋展属山岳岸島川州巡工左差市布帝師席帯帰常幅幕幡平年幹幻庁広店府度座庫庭庵廃延廷建弁式引張強当形彩役彼待律後徒従得御復徴心必忌志応快念思急性恩息恵患悪情惑想意愛感態慎慮慰憲憶懇懐懲懸成戒戦戸房所手才打払批承技投抗抽担拘招拝拠拡括持指挑挙振捜据掃授掘掛採探接推措提換揮援損搭携摘撃撲擁支改攻放政故救敗教敢散数整敷文料断新方施旅族日早旬昇明易星映春昨昭是昼時普景暗暫暮暴曜曲更書最月有望朝期木未末本札材村束条来東林果枝枠架染柔柳査校株核根格案桑械棄棚棟植検業極楽概構様標模権横樹橋機欠次欧欲歌歓止正武歩歯歳歴死殊残殖殺殿毎比氏民気水求汚江決沢河油治沼沿況法波注泳洋活派流海消清済渉減測港湾満源準滑滞演漫潤激火災炭点為無焦然照熱燃爆物特犬状独献獲率玉王珠現球理環生産用田由甲申男町画界畑留略番異畿疑疲病症療癒発登白百的皮益盗盟監盤目盲直相省看県真着知石研砕硬確礎示社神祥票禁秀私秋科秘移程税種穀積穫究空窓窮立章端競笑符第等筋筒答策算管築簿米粧精糖糧系紀約納紙紛素終組経結絡給統絵絶継続維綱緊総線締編緩緯練縮績織罪置署美義習老考者聞聴職肉育背能脳臓臨自舎舗舞航般良色芝芸若苦英草荒荘菜落葉著葛葬蓄蔓蔵薬藤虐融血衆行術街衛表被袴裁装裏補製複西要見規視覚覧親観角解言計討託記訪設許訴診証評試話詳誉認誘語説読課調談請論諭謀講謝識警議護谷象負財貨販責貴買貸費賀賄資賛賞質購贅贈起超越足跡路躍身車軌軍転軸軽較輪輸辞農辺込近返述追退送逆途通速造連週進遇遊運過道達違遠遣適選遺避還邦邸郊部郵都配酬酸里重野量金針鉄銀錠録長門開間関閣閥闘阪防阻降限院陣除陳陸険隊階際障隠隣集雇離難雰電需青非面革靴韓音響項順領頭頼題額顔類顧風食飲飼養館首馬駅駆駐験高鮮齢
