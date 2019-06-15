package github.com/ciuak/jvozba

func Jvozba(tanru string) (string, error) {
	return Zbasu(tanru, Rafsi, false)
}

func Cmejvozba(tanru string) (string, error) {
	return Zbasu(tanru, Rafsi, true)
}

var Rafsi = map[string][]string{
	"barna": []string{"ba'a"},
	"balre": []string{"ba'e"},
	"banli": []string{"ba'i", "bal"},
	"banro": []string{"ba'o"},
	"bacru": []string{"ba'u"},
	"zbabu": []string{"bab"},
	"bancu": []string{"bac"},
	"bandu": []string{"bad"},
	"bakfu": []string{"baf"},
	"bargu": []string{"bag"},
	"bapli": []string{"bai", "bap"},
	"bajra": []string{"baj"},
	"bakni": []string{"bak"},
	"jbama": []string{"bam"},
	"bangu": []string{"ban", "bau"},
	"bartu": []string{"bar"},
	"basti": []string{"bas"},
	"batci": []string{"bat"},
	"balvi": []string{"bav"},
	"banxa": []string{"bax"},
	"banzu": []string{"baz"},
	"bersa": []string{"be'a", "bes"},
	"bende": []string{"be'e", "bed"},
	"benji": []string{"be'i", "bej"},
	"bemro": []string{"be'o", "bem"},
	"betfu": []string{"be'u", "bef"},
	"bebna": []string{"beb"},
	"bengo": []string{"beg"},
	"bevri": []string{"bei", "bev"},
	"besna": []string{"ben"},
	"berti": []string{"ber"},
	"betri": []string{"bet"},
	"bilma": []string{"bi'a"},
	"brife": []string{"bi'e", "bif"},
	"jbini": []string{"bi'i", "bin"},
	"binxo": []string{"bi'o", "bix"},
	"bitmu": []string{"bi'u", "bim"},
	"bifce": []string{"bic"},
	"bindo": []string{"bid"},
	"bilga": []string{"big"},
	"briju": []string{"bij"},
	"bikla": []string{"bik"},
	"bilni": []string{"bil"},
	"birka": []string{"bir"},
	"bisli": []string{"bis"},
	"birti": []string{"bit"},
	"bi":    []string{"biv"},
	"bi'i":  []string{"biz"},
	"blanu": []string{"bla"},
	"ruble": []string{"ble", "rub"},
	"bliku": []string{"bli"},
	"bloti": []string{"blo", "lo'i", "lot"},
	"ciblu": []string{"blu"},
	"boxna": []string{"bo'a", "bon"},
	"brode": []string{"bo'e"},
	"botpi": []string{"bo'i", "bot"},
	"boxfo": []string{"bo'o", "bof"},
	"bongu": []string{"bo'u", "bog"},
	"bolci": []string{"boi", "bol"},
	"bo":    []string{"bor"},
	"barda": []string{"bra"},
	"bredi": []string{"bre", "red"},
	"bridi": []string{"bri"},
	"xebro": []string{"bro", "xeb"},
	"burcu": []string{"bru"},
	"bruna": []string{"bu'a", "bun"},
	"bunre": []string{"bu'e", "bur"},
	"bu":    []string{"bu'i", "bus"},
	"budjo": []string{"bu'o", "buj"},
	"bukpu": []string{"bu'u", "buk"},
	"bunda": []string{"bud"},
	"bu'a":  []string{"bul"},
	"bumru": []string{"bum"},
	"cabra": []string{"ca'a"},
	"catke": []string{"ca'e"},
	"catni": []string{"ca'i"},
	"canko": []string{"ca'o"},
	"canlu": []string{"ca'u", "cal"},
	"cabna": []string{"cab"},
	"tcaci": []string{"cac"},
	"cando": []string{"cad"},
	"cafne": []string{"caf"},
	"cange": []string{"cag"},
	"carmi": []string{"cai", "cam"},
	"canja": []string{"caj"},
	"calku": []string{"cak"},
	"canre": []string{"can"},
	"ckape": []string{"cap"},
	"carna": []string{"car"},
	"ckasu": []string{"cas"},
	"cartu": []string{"cat"},
	"claxu": []string{"cau"},
	"carvi": []string{"cav"},
	"caxno": []string{"cax"},
	"ca'a":  []string{"caz"},
	"cecla": []string{"ce'a", "cel"},
	"cteki": []string{"ce'i", "tek"},
	"ce'o":  []string{"ce'o"},
	"cecmu": []string{"ce'u", "cem"},
	"ce":    []string{"cec"},
	"cerda": []string{"ced"},
	"cevni": []string{"cei", "cev"},
	"creka": []string{"cek"},
	"centi": []string{"cen"},
	"cerni": []string{"cer"},
	"censa": []string{"ces"},
	"ce'i":  []string{"cez"},
	"cfari": []string{"cfa"},
	"cfila": []string{"cfi"},
	"ricfu": []string{"cfu", "rif"},
	"ciska": []string{"ci'a"},
	"ciste": []string{"ci'e"},
	"cinri": []string{"ci'i"},
	"citno": []string{"ci'o", "cit"},
	"ckilu": []string{"ci'u"},
	"ci":    []string{"cib"},
	"cilce": []string{"cic"},
	"cidni": []string{"cid"},
	"cifnu": []string{"cif"},
	"cigla": []string{"cig"},
	"cinje": []string{"cij"},
	"cikna": []string{"cik"},
	"cilta": []string{"cil"},
	"cilmo": []string{"cim"},
	"cinse": []string{"cin"},
	"cipra": []string{"cip"},
	"citri": []string{"cir"},
	"crisa": []string{"cis"},
	"civla": []string{"civ"},
	"cizra": []string{"ciz"},
	"ckana": []string{"cka"},
	"ckeji": []string{"cke", "kej"},
	"ciksi": []string{"cki"},
	"cokcu": []string{"cko"},
	"cukta": []string{"cku"},
	"clani": []string{"cla"},
	"cilre": []string{"cli"},
	"culno": []string{"clu"},
	"cmalu": []string{"cma"},
	"cmene": []string{"cme", "me'e"},
	"cmima": []string{"cmi", "mim"},
	"cmoni": []string{"cmo", "co'i"},
	"jicmu": []string{"cmu"},
	"canpa": []string{"cna"},
	"cenba": []string{"cne"},
	"cinmo": []string{"cni"},
	"condi": []string{"cno", "coi", "con"},
	"macnu": []string{"cnu"},
	"co'a":  []string{"co'a"},
	"co'e":  []string{"co'e", "com"},
	"co'u":  []string{"co'u"},
	"co":    []string{"col"},
	"cortu": []string{"cor", "cro"},
	"cpacu": []string{"cpa"},
	"cpedu": []string{"cpe"},
	"cipni": []string{"cpi"},
	"lacpu": []string{"cpu", "lap"},
	"crane": []string{"cra"},
	"certu": []string{"cre"},
	"cirko": []string{"cri"},
	"curmi": []string{"cru"},
	"catlu": []string{"cta"},
	"nicte": []string{"cte"},
	"citka": []string{"cti"},
	"xecto": []string{"cto", "xet"},
	"ctuca": []string{"ctu"},
	"cuxna": []string{"cu'a", "cux"},
	"ckule": []string{"cu'e", "kul"},
	"cumki": []string{"cu'i", "cum"},
	"cunso": []string{"cu'o", "cun"},
	"cuntu": []string{"cu'u"},
	"cutci": []string{"cuc"},
	"cukla": []string{"cuk"},
	"cumla": []string{"cul"},
	"clupa": []string{"cup"},
	"curnu": []string{"cur"},
	"cusku": []string{"cus", "sku"},
	"cutne": []string{"cut"},
	"curve": []string{"cuv"},
	"damba": []string{"da'a", "dab"},
	"danre": []string{"da'e"},
	"darxi": []string{"da'i", "dax"},
	"darno": []string{"da'o", "dar"},
	"danlu": []string{"da'u", "dal"},
	"dacru": []string{"dac"},
	"dandu": []string{"dad"},
	"danfu": []string{"daf"},
	"dargu": []string{"dag"},
	"dacti": []string{"dai"},
	"dadjo": []string{"daj"},
	"dakfu": []string{"dak"},
	"danmo": []string{"dam"},
	"danti": []string{"dan"},
	"dapma": []string{"dap"},
	"dasni": []string{"das"},
	"drata": []string{"dat"},
	"darlu": []string{"dau"},
	"da":    []string{"dav", "dza"},
	"da'a":  []string{"daz"},
	"denpa": []string{"de'a", "dep"},
	"denci": []string{"de'i", "den"},
	"delno": []string{"de'o", "del"},
	"dertu": []string{"de'u", "der"},
	"dembi": []string{"deb"},
	"decti": []string{"dec"},
	"degji": []string{"deg"},
	"djedi": []string{"dei", "dje"},
	"dejni": []string{"dej"},
	"dekto": []string{"dek"},
	"denmi": []string{"dem"},
	"desku": []string{"des"},
	"detri": []string{"det"},
	"jdima": []string{"di'a"},
	"dirce": []string{"di'e"},
	"jdini": []string{"di'i", "din"},
	"dinko": []string{"di'o"},
	"dinju": []string{"di'u", "dij"},
	"dirba": []string{"dib"},
	"dikca": []string{"dic"},
	"dirgo": []string{"dig"},
	"dikni": []string{"dik"},
	"dilnu": []string{"dil"},
	"dimna": []string{"dim"},
	"dicra": []string{"dir"},
	"dizlo": []string{"diz", "dzi"},
	"cidja": []string{"dja"},
	"djica": []string{"dji"},
	"sadjo": []string{"djo"},
	"sidju": []string{"dju"},
	"donri": []string{"do'i", "dor"},
	"dotco": []string{"do'o", "dot"},
	"do":    []string{"doi", "don"},
	"drani": []string{"dra"},
	"derxi": []string{"dre"},
	"badri": []string{"dri"},
	"cidro": []string{"dro"},
	"drudi": []string{"dru", "rud"},
	"dunda": []string{"du'a", "dud"},
	"dukse": []string{"du'e", "dus"},
	"dunli": []string{"du'i", "dun"},
	"du":    []string{"du'o", "dub"},
	"dunku": []string{"du'u", "duk"},
	"dugri": []string{"dug"},
	"dunja": []string{"duj"},
	"jduli": []string{"dul", "jdu"},
	"du'u":  []string{"dum"},
	"dunra": []string{"dur"},
	"dukti": []string{"dut"},
	"dzena": []string{"dze"},
	"cadzu": []string{"dzu"},
	"farna": []string{"fa'a", "far"},
	"fatne": []string{"fa'e", "fat"},
	"facki": []string{"fa'i", "fak"},
	"fanmo": []string{"fa'o", "fam"},
	"farlu": []string{"fa'u", "fal"},
	"fatci": []string{"fac"},
	"fadni": []string{"fad"},
	"fagri": []string{"fag"},
	"fatri": []string{"fai"},
	"falnu": []string{"fan"},
	"fapro": []string{"fap", "pro"},
	"fraso": []string{"fas"},
	"fasnu": []string{"fau"},
	"farvi": []string{"fav"},
	"fraxu": []string{"fax"},
	"fanza": []string{"faz"},
	"fenra": []string{"fe'a", "fer"},
	"fetsi": []string{"fe'i", "fet"},
	"fenso": []string{"fe'o", "fen"},
	"fengu": []string{"fe'u", "feg"},
	"febvi": []string{"feb"},
	"fendi": []string{"fed"},
	"fepni": []string{"fei", "fep"},
	"fenki": []string{"fek"},
	"femti": []string{"fem"},
	"festi": []string{"fes"},
	"cfika": []string{"fi'a", "fik"},
	"finpe": []string{"fi'e", "fip"},
	"finti": []string{"fi'i", "fin"},
	"friko": []string{"fi'o"},
	"cfipu": []string{"fi'u"},
	"frica": []string{"fic"},
	"figre": []string{"fig"},
	"frili": []string{"fil"},
	"flira": []string{"fir"},
	"filso": []string{"fis"},
	"friti": []string{"fit"},
	"flalu": []string{"fla"},
	"flecu": []string{"fle"},
	"fliba": []string{"fli"},
	"foldi": []string{"flo", "foi"},
	"fulta": []string{"flu", "ful"},
	"fo'a":  []string{"fo'a"},
	"fo'e":  []string{"fo'e"},
	"fo'i":  []string{"fo'i"},
	"fonmo": []string{"fo'o", "fom"},
	"fonxa": []string{"fon"},
	"frati": []string{"fra"},
	"ferti": []string{"fre"},
	"lifri": []string{"fri", "lif"},
	"forca": []string{"fro"},
	"frumu": []string{"fru"},
	"funca": []string{"fu'a", "fun"},
	"fuzme": []string{"fu'e", "fuz"},
	"fukpi": []string{"fu'i", "fuk"},
	"fusra": []string{"fur"},
	"grana": []string{"ga'a"},
	"ganse": []string{"ga'e", "gas"},
	"galfi": []string{"ga'i", "gaf"},
	"ganlo": []string{"ga'o"},
	"galtu": []string{"ga'u", "gal"},
	"gapci": []string{"gac"},
	"gadri": []string{"gad"},
	"gacri": []string{"gai"},
	"ganra": []string{"gan"},
	"gapru": []string{"gap"},
	"garna": []string{"gar"},
	"gasta": []string{"gat"},
	"gasnu": []string{"gau"},
	"ganxo": []string{"gax"},
	"ganzu": []string{"gaz"},
	"gerna": []string{"ge'a", "gen"},
	"gento": []string{"ge'o", "get"},
	"gerku": []string{"ge'u", "ger"},
	"gleki": []string{"gei", "gek"},
	"genja": []string{"gej"},
	"genxu": []string{"gex"},
	"gidva": []string{"gi'a", "gid"},
	"zgike": []string{"gi'e", "zgi"},
	"gigdo": []string{"gi'o", "gig"},
	"gismu": []string{"gi'u", "gim"},
	"glico": []string{"gic", "gli"},
	"ginka": []string{"gik"},
	"jgina": []string{"gin"},
	"girzu": []string{"gir", "gri"},
	"jgita": []string{"git"},
	"glare": []string{"gla"},
	"gletu": []string{"gle", "let"},
	"gluta": []string{"glu"},
	"gocti": []string{"goc"},
	"gotro": []string{"got"},
	"grake": []string{"gra"},
	"pagre": []string{"gre"},
	"gurni": []string{"gru"},
	"gunka": []string{"gu'a", "gun"},
	"gugde": []string{"gu'e", "gug"},
	"gusni": []string{"gu'i", "gus"},
	"gunro": []string{"gu'o", "gur"},
	"gubni": []string{"gub"},
	"gutci": []string{"guc"},
	"gundi": []string{"gud"},
	"guska": []string{"guk"},
	"gunma": []string{"gum"},
	"gunta": []string{"gut"},
	"guzme": []string{"guz", "zme"},
	"jatna": []string{"ja'a"},
	"jalge": []string{"ja'e", "jag"},
	"jadni": []string{"ja'i", "jad"},
	"jarco": []string{"ja'o"},
	"jgalu": []string{"ja'u"},
	"janbe": []string{"jab"},
	"djacu": []string{"jac", "jau"},
	"jamfu": []string{"jaf", "jma"},
	"jgari": []string{"jai"},
	"jmaji": []string{"jaj"},
	"jarki": []string{"jak"},
	"janli": []string{"jal"},
	"jamna": []string{"jam"},
	"janco": []string{"jan"},
	"jaspu": []string{"jap"},
	"jdari": []string{"jar"},
	"jansu": []string{"jas"},
	"janta": []string{"jat"},
	"ja":    []string{"jav"},
	"jbari": []string{"jba"},
	"jbena": []string{"jbe"},
	"jibni": []string{"jbi"},
	"lojbo": []string{"jbo", "lob"},
	"jubme": []string{"jbu", "jub"},
	"lijda": []string{"jda"},
	"kajde": []string{"jde"},
	"jdice": []string{"jdi"},
	"jecta": []string{"je'a", "jec"},
	"jetce": []string{"je'e"},
	"jersi": []string{"je'i"},
	"jegvo": []string{"je'o", "jeg"},
	"jetnu": []string{"je'u", "jet"},
	"jendu": []string{"jed"},
	"jeftu": []string{"jef"},
	"jenmi": []string{"jei", "jem"},
	"jelca": []string{"jel"},
	"jenca": []string{"jen"},
	"jbera": []string{"jer"},
	"jesni": []string{"jes"},
	"je":    []string{"jev", "jve"},
	"jerxo": []string{"jex"},
	"jei":   []string{"jez"},
	"jganu": []string{"jga"},
	"jgena": []string{"jge"},
	"jgira": []string{"jgi"},
	"jinga": []string{"ji'a", "jig"},
	"jmive": []string{"ji'e", "miv"},
	"jinvi": []string{"ji'i", "jiv"},
	"jipno": []string{"ji'o", "jip"},
	"jvinu": []string{"ji'u", "vin"},
	"jibri": []string{"jib"},
	"jimca": []string{"jic"},
	"jitfa": []string{"jif"},
	"jijnu": []string{"jij"},
	"jikca": []string{"jik"},
	"jilka": []string{"jil"},
	"jinme": []string{"jim"},
	"djine": []string{"jin"},
	"jinru": []string{"jir"},
	"jinsa": []string{"jis"},
	"jimte": []string{"jit"},
	"jinzi": []string{"jiz"},
	"jemna": []string{"jme"},
	"jimpe": []string{"jmi"},
	"jorne": []string{"jo'e", "jon"},
	"jordo": []string{"jo'o", "jor"},
	"jo'u":  []string{"jo'u"},
	"joi":   []string{"joi", "jol"},
	"jo'e":  []string{"jom"},
	"jo":    []string{"jov"},
	"jufra": []string{"ju'a", "juf"},
	"julne": []string{"ju'e"},
	"jundi": []string{"ju'i", "jud"},
	"djuno": []string{"ju'o", "jun"},
	"jungo": []string{"jug"},
	"jukni": []string{"juk"},
	"junla": []string{"jul"},
	"jurme": []string{"jum"},
	"jukpa": []string{"jup"},
	"junri": []string{"jur"},
	"jursa": []string{"jus"},
	"jutsi": []string{"jut"},
	"ju":    []string{"juv"},
	"juxre": []string{"jux"},
	"javni": []string{"jva"},
	"jivna": []string{"jvi"},
	"lujvo": []string{"jvo", "luv"},
	"katna": []string{"ka'a"},
	"kakne": []string{"ka'e"},
	"krati": []string{"ka'i"},
	"kanro": []string{"ka'o"},
	"kantu": []string{"ka'u"},
	"karbi": []string{"kab"},
	"kancu": []string{"kac"},
	"kandi": []string{"kad"},
	"ckafi": []string{"kaf"},
	"kagni": []string{"kag"},
	"ckaji": []string{"kai"},
	"kanji": []string{"kaj"},
	"klaku": []string{"kak"},
	"kanla": []string{"kal"},
	"ka":    []string{"kam"},
	"kansa": []string{"kan"},
	"skapi": []string{"kap"},
	"kalri": []string{"kar"},
	"kalsa": []string{"kas"},
	"kalte": []string{"kat"},
	"kampu": []string{"kau"},
	"kavbu": []string{"kav"},
	"kanxe": []string{"kax"},
	"kevna": []string{"ke'a", "kev"},
	"ke'e":  []string{"ke'e", "kep"},
	"kecti": []string{"ke'i", "kec"},
	"kelvo": []string{"ke'o"},
	"krefu": []string{"ke'u", "ref"},
	"kelci": []string{"kei", "kel"},
	"ke":    []string{"kem"},
	"kenra": []string{"ken"},
	"kerlo": []string{"ker"},
	"kensa": []string{"kes"},
	"ketco": []string{"ket", "tco"},
	"kei":   []string{"kez"},
	"krixa": []string{"ki'a", "kix"},
	"kicne": []string{"ki'e", "kic"},
	"ckini": []string{"ki'i"},
	"kilto": []string{"ki'o"},
	"krinu": []string{"ki'u", "rin"},
	"kijno": []string{"kij"},
	"ckiku": []string{"kik"},
	"kinli": []string{"kil"},
	"skina": []string{"kin"},
	"ckire": []string{"kir"},
	"kisto": []string{"kis"},
	"kliti": []string{"kit"},
	"klama": []string{"kla"},
	"klesi": []string{"kle", "lei"},
	"klina": []string{"kli"},
	"diklo": []string{"klo"},
	"kulnu": []string{"klu"},
	"kojna": []string{"ko'a", "koj"},
	"kolme": []string{"ko'e", "kol"},
	"kobli": []string{"ko'i", "kob"},
	"skoto": []string{"ko'o", "kot"},
	"konju": []string{"ko'u", "kon"},
	"korbi": []string{"koi", "kor"},
	"korka": []string{"kok"},
	"komcu": []string{"kom"},
	"kosta": []string{"kos"},
	"krasi": []string{"kra"},
	"kerfa": []string{"kre"},
	"krici": []string{"kri"},
	"korcu": []string{"kro"},
	"kruvi": []string{"kru", "ruv"},
	"kumfa": []string{"ku'a", "kum"},
	"kuspe": []string{"ku'e", "kup"},
	"kurji": []string{"ku'i", "kuj"},
	"skuro": []string{"ku'o"},
	"ckunu": []string{"ku'u"},
	"kubli": []string{"kub"},
	"kruca": []string{"kuc"},
	"kufra": []string{"kuf"},
	"kukte": []string{"kuk"},
	"kunra": []string{"kun"},
	"kurfa": []string{"kur"},
	"kusru": []string{"kus"},
	"kunti": []string{"kut"},
	"ku'a":  []string{"kuz"},
	"lasna": []string{"la'a"},
	"lakne": []string{"la'e"},
	"lamji": []string{"la'i", "lam"},
	"latmo": []string{"la'o"},
	"lalxu": []string{"la'u"},
	"blabi": []string{"lab"},
	"lacri": []string{"lac"},
	"ladru": []string{"lad"},
	"lafti": []string{"laf"},
	"vlagi": []string{"lag"},
	"klani": []string{"lai"},
	"klaji": []string{"laj"},
	"lakse": []string{"lak"},
	"lanli": []string{"lal"},
	"lanme": []string{"lan"},
	"larcu": []string{"lar"},
	"slasi": []string{"las"},
	"mlatu": []string{"lat"},
	"cladu": []string{"lau"},
	"lanxe": []string{"lax"},
	"lanzu": []string{"laz"},
	"lebna": []string{"le'a", "leb"},
	"pleji": []string{"le'i", "lej"},
	"lenjo": []string{"le'o", "len"},
	"lerfu": []string{"le'u", "ler"},
	"lerci": []string{"lec"},
	"mledi": []string{"led"},
	"lenku": []string{"lek"},
	"le'e":  []string{"lem"},
	"cliva": []string{"li'a", "liv"},
	"lidne": []string{"li'e"},
	"linji": []string{"li'i", "lij"},
	"linto": []string{"li'o"},
	"litru": []string{"li'u"},
	"libjo": []string{"lib"},
	"litce": []string{"lic"},
	"lindi": []string{"lid"},
	"sligu": []string{"lig"},
	"litki": []string{"lik"},
	"livla": []string{"lil"},
	"limna": []string{"lim"},
	"linsi": []string{"lin"},
	"clira": []string{"lir"},
	"lisri": []string{"lis"},
	"clite": []string{"lit"},
	"plixa": []string{"lix"},
	"li'i":  []string{"liz"},
	"slovo": []string{"lo'o", "lov"},
	"lorxu": []string{"lo'u", "lor"},
	"loldi": []string{"loi", "lol"},
	"logji": []string{"loj"},
	"lo'e":  []string{"lom"},
	"pluta": []string{"lu'a", "lut"},
	"klupe": []string{"lu'e", "lup"},
	"lumci": []string{"lu'i", "lum"},
	"lubno": []string{"lu'o"},
	"lunbe": []string{"lub"},
	"pluja": []string{"luj"},
	"mluni": []string{"lun"},
	"lunra": []string{"lur"},
	"lunsa": []string{"lus"},
	"kluza": []string{"luz"},
	"cmana": []string{"ma'a"},
	"marce": []string{"ma'e"},
	"masti": []string{"ma'i"},
	"cmavo": []string{"ma'o"},
	"makcu": []string{"ma'u"},
	"mabru": []string{"mab"},
	"manci": []string{"mac"},
	"marde": []string{"mad"},
	"makfa": []string{"maf"},
	"margu": []string{"mag"},
	"marji": []string{"mai", "maj"},
	"maksi": []string{"mak"},
	"mabla": []string{"mal"},
	"mamta": []string{"mam"},
	"manku": []string{"man"},
	"mapku": []string{"map"},
	"manri": []string{"mar"},
	"malsi": []string{"mas"},
	"mapti": []string{"mat"},
	"zmadu": []string{"mau", "zma"},
	"mavji": []string{"mav"},
	"marxa": []string{"max"},
	"mleca": []string{"me'a", "mec"},
	"mensi": []string{"me'i", "mes"},
	"mekso": []string{"me'o", "mek"},
	"mentu": []string{"me'u", "met"},
	"mebri": []string{"meb"},
	"megdo": []string{"meg"},
	"mei":   []string{"mei", "mem"},
	"meljo": []string{"mej"},
	"melbi": []string{"mel", "mle"},
	"menli": []string{"men"},
	"merko": []string{"mer"},
	"mexno": []string{"mex"},
	"cmila": []string{"mi'a"},
	"minde": []string{"mi'e", "mid"},
	"minji": []string{"mi'i"},
	"misno": []string{"mi'o", "mis"},
	"mintu": []string{"mi'u", "mit"},
	"mi":    []string{"mib"},
	"mikce": []string{"mic"},
	"mifra": []string{"mif"},
	"midju": []string{"mij"},
	"mikri": []string{"mik"},
	"milti": []string{"mil"},
	"jmina": []string{"min"},
	"mipri": []string{"mip"},
	"minra": []string{"mir"},
	"mixre": []string{"mix", "xre"},
	"mlana": []string{"mla"},
	"milxe": []string{"mli"},
	"molki": []string{"mlo"},
	"simlu": []string{"mlu"},
	"morna": []string{"mo'a", "mon"},
	"morji": []string{"mo'i", "moj"},
	"molro": []string{"mo'o"},
	"moklu": []string{"mo'u", "mol"},
	"mo'a":  []string{"mob"},
	"mokca": []string{"moc"},
	"moi":   []string{"moi", "mom"},
	"morko": []string{"mor"},
	"mosra": []string{"mos"},
	"mo'i":  []string{"mov"},
	"marbi": []string{"mra"},
	"merli": []string{"mre"},
	"mrilu": []string{"mri"},
	"morsi": []string{"mro"},
	"mruli": []string{"mru"},
	"murta": []string{"mu'a", "mur"},
	"munje": []string{"mu'e", "muj"},
	"mukti": []string{"mu'i", "muk"},
	"mulno": []string{"mu'o", "mul"},
	"muvdu": []string{"mu'u", "muv"},
	"smuci": []string{"muc"},
	"mudri": []string{"mud"},
	"mu'e":  []string{"muf"},
	"mu":    []string{"mum"},
	"smuni": []string{"mun", "smu"},
	"mupli": []string{"mup"},
	"muslo": []string{"mus"},
	"mucti": []string{"mut"},
	"muzga": []string{"muz"},
	"nanca": []string{"na'a"},
	"natfe": []string{"na'e", "naf"},
	"nalci": []string{"na'i"},
	"cnano": []string{"na'o"},
	"namcu": []string{"na'u", "nac"},
	"nanba": []string{"nab"},
	"nandu": []string{"nad"},
	"narge": []string{"nag"},
	"natmi": []string{"nai", "nat"},
	"narju": []string{"naj"},
	"nakni": []string{"nak"},
	"na'e":  []string{"nal"},
	"nabmi": []string{"nam"},
	"snanu": []string{"nan"},
	"na":    []string{"nar"},
	"nanmu": []string{"nau"},
	"nanvi": []string{"nav"},
	"naxle": []string{"nax", "xle"},
	"nazbi": []string{"naz", "zbi"},
	"nenri": []string{"ne'i", "ner"},
	"cnebo": []string{"ne'o", "neb"},
	"cnemu": []string{"ne'u", "nem"},
	"nelci": []string{"nei", "nel"},
	"nejni": []string{"nen"},
	"cnita": []string{"ni'a", "nit"},
	"nilce": []string{"ni'e"},
	"nibli": []string{"ni'i", "nib"},
	"cnino": []string{"ni'o", "nin"},
	"ninmu": []string{"ni'u", "nim"},
	"cnici": []string{"nic"},
	"snidu": []string{"nid"},
	"nikle": []string{"nik"},
	"ni":    []string{"nil"},
	"snipa": []string{"nip"},
	"nirna": []string{"nir"},
	"cnisa": []string{"nis"},
	"nivji": []string{"niv"},
	"nixli": []string{"nix", "xli"},
	"no'e":  []string{"no'e", "nor"},
	"nobli": []string{"no'i", "nol"},
	"notci": []string{"noi", "not"},
	"no":    []string{"non"},
	"snura": []string{"nu'a", "nur"},
	"nupre": []string{"nu'e", "nup"},
	"nutli": []string{"nu'i", "nul"},
	"nu'o":  []string{"nu'o"},
	"snuji": []string{"nuj"},
	"nukni": []string{"nuk"},
	"nurma": []string{"num"},
	"nu":    []string{"nun"},
	"snuti": []string{"nut"},
	"nuzba": []string{"nuz"},
	"pacna": []string{"pa'a"},
	"prane": []string{"pa'e"},
	"prami": []string{"pa'i", "pam"},
	"panlo": []string{"pa'o"},
	"patfu": []string{"pa'u", "paf"},
	"parbi": []string{"pab"},
	"palci": []string{"pac"},
	"pandi": []string{"pad"},
	"pagbu": []string{"pag", "pau"},
	"pajni": []string{"pai"},
	"spaji": []string{"paj"},
	"palku": []string{"pak"},
	"prali": []string{"pal"},
	"panci": []string{"pan"},
	"panpi": []string{"pap"},
	"cpare": []string{"par"},
	"pastu": []string{"pas"},
	"pante": []string{"pat"},
	"pa":    []string{"pav"},
	"patxu": []string{"pax"},
	"panzi": []string{"paz"},
	"preja": []string{"pe'a", "pej"},
	"penmi": []string{"pe'i", "pen"},
	"pendo": []string{"pe'o", "ped"},
	"pencu": []string{"pe'u", "pec"},
	"penbi": []string{"peb"},
	"pensi": []string{"pei", "pes"},
	"pelxu": []string{"pel"},
	"pemci": []string{"pem"},
	"perli": []string{"per"},
	"petso": []string{"pet"},
	"pe'a":  []string{"pev"},
	"pesxu": []string{"pex"},
	"pezli": []string{"pez"},
	"pilka": []string{"pi'a", "pil"},
	"plipe": []string{"pi'e", "pip"},
	"pilji": []string{"pi'i"},
	"pipno": []string{"pi'o"},
	"pimlu": []string{"pi'u", "pim"},
	"plibu": []string{"pib"},
	"picti": []string{"pic"},
	"pindi": []string{"pid"},
	"pinfu": []string{"pif"},
	"prije": []string{"pij"},
	"pinka": []string{"pik"},
	"pinta": []string{"pin"},
	"pixra": []string{"pir", "xra"},
	"pinsi": []string{"pis"},
	"plita": []string{"pit"},
	"pi'u":  []string{"piv"},
	"pinxe": []string{"pix"},
	"pi":    []string{"piz"},
	"platu": []string{"pla"},
	"pelji": []string{"ple"},
	"pilno": []string{"pli"},
	"polje": []string{"plo"},
	"daplu": []string{"plu"},
	"spoja": []string{"po'a", "poj"},
	"ponse": []string{"po'e", "pos"},
	"porpi": []string{"po'i", "pop"},
	"ponjo": []string{"po'o", "pon"},
	"spofu": []string{"po'u", "pof"},
	"porsi": []string{"poi", "por"},
	"polno": []string{"pol"},
	"porto": []string{"pot"},
	"cupra": []string{"pra"},
	"prenu": []string{"pre"},
	"prina": []string{"pri"},
	"purci": []string{"pru", "pur"},
	"pluka": []string{"pu'a", "puk"},
	"pulce": []string{"pu'e", "puc"},
	"punji": []string{"pu'i", "puj"},
	"purmo": []string{"pu'o", "pum"},
	"sputu": []string{"pu'u", "put"},
	"purdi": []string{"pud"},
	"punli": []string{"pul"},
	"pruni": []string{"pun"},
	"pu'i":  []string{"pus"},
	"pu'u":  []string{"puv"},
	"srana": []string{"ra'a"},
	"ralte": []string{"ra'e"},
	"ranji": []string{"ra'i"},
	"radno": []string{"ra'o"},
	"raktu": []string{"ra'u"},
	"xrabo": []string{"rab"},
	"ralci": []string{"rac"},
	"randa": []string{"rad"},
	"rafsi": []string{"raf"},
	"rango": []string{"rag"},
	"traji": []string{"rai"},
	"sraji": []string{"raj"},
	"sraku": []string{"rak"},
	"ralju": []string{"ral"},
	"ranmi": []string{"ram"},
	"ranti": []string{"ran"},
	"rapli": []string{"rap"},
	"rarna": []string{"rar"},
	"grasu": []string{"ras"},
	"ratni": []string{"rat"},
	"gradu": []string{"rau"},
	"ragve": []string{"rav"},
	"ranxi": []string{"rax"},
	"brazo": []string{"raz"},
	"remna": []string{"re'a", "rem"},
	"trene": []string{"re'e", "ren"},
	"renvi": []string{"re'i", "rev"},
	"renro": []string{"re'o", "rer"},
	"rectu": []string{"re'u", "rec"},
	"rebla": []string{"reb"},
	"preti": []string{"rei", "ret"},
	"vreji": []string{"rej", "vei"},
	"greku": []string{"rek"},
	"re":    []string{"rel"},
	"crepu": []string{"rep"},
	"respa": []string{"res"},
	"rinka": []string{"ri'a", "rik"},
	"rirxe": []string{"ri'e"},
	"ritli": []string{"ri'i"},
	"crino": []string{"ri'o"},
	"rinju": []string{"ri'u"},
	"cribe": []string{"rib"},
	"tricu": []string{"ric"},
	"crida": []string{"rid"},
	"rigni": []string{"rig"},
	"rijno": []string{"rij"},
	"rilti": []string{"ril"},
	"rimni": []string{"rim"},
	"cripu": []string{"rip"},
	"rirni": []string{"rir"},
	"rismi": []string{"ris"},
	"brito": []string{"rit"},
	"rivbi": []string{"riv"},
	"trixe": []string{"rix", "ti'e"},
	"prosa": []string{"ro'a", "ros"},
	"rokci": []string{"ro'i", "rok"},
	"ropno": []string{"ro'o", "ron"},
	"rotsu": []string{"ro'u", "rot", "tsu"},
	"broda": []string{"rod"},
	"romge": []string{"rog"},
	"roi":   []string{"roi", "rom"},
	"ro":    []string{"rol"},
	"rorci": []string{"ror"},
	"rozgu": []string{"roz", "zgu"},
	"sruma": []string{"ru'a"},
	"pruce": []string{"ru'e", "ruc"},
	"pruxi": []string{"ru'i", "rux"},
	"rusko": []string{"ru'o", "ruk"},
	"rupnu": []string{"ru'u", "rup"},
	"rufsu": []string{"ruf"},
	"kruji": []string{"ruj"},
	"xrula": []string{"rul"},
	"runme": []string{"rum"},
	"rutni": []string{"run"},
	"sruri": []string{"rur", "sru"},
	"grusi": []string{"rus"},
	"grute": []string{"rut"},
	"sanga": []string{"sa'a", "sag"},
	"satre": []string{"sa'e"},
	"sanli": []string{"sa'i"},
	"salpo": []string{"sa'o"},
	"sarcu": []string{"sa'u"},
	"sabji": []string{"sab"},
	"stace": []string{"sac"},
	"snada": []string{"sad"},
	"sanmi": []string{"sai"},
	"sanji": []string{"saj"},
	"sakci": []string{"sak"},
	"sakli": []string{"sal"},
	"skami": []string{"sam"},
	"spano": []string{"san"},
	"sampu": []string{"sap"},
	"slari": []string{"sar"},
	"srasu": []string{"sas"},
	"sakta": []string{"sat"},
	"slabu": []string{"sau"},
	"savru": []string{"sav", "vru"},
	"sarxe": []string{"sax"},
	"sazri": []string{"saz"},
	"setca": []string{"se'a"},
	"sevzi": []string{"se'i", "sez"},
	"selfu": []string{"se'u", "sef"},
	"steba": []string{"seb"},
	"senci": []string{"sec"},
	"stedu": []string{"sed"},
	"sepli": []string{"sei", "sep"},
	"se":    []string{"sel"},
	"senpi": []string{"sen"},
	"serti": []string{"ser"},
	"senta": []string{"set"},
	"senva": []string{"sev", "sne"},
	"sfasa": []string{"sfa"},
	"sefta": []string{"sfe"},
	"sfofa": []string{"sfo"},
	"sufti": []string{"sfu"},
	"sinma": []string{"si'a"},
	"snime": []string{"si'e"},
	"sicni": []string{"si'i"},
	"sidbo": []string{"si'o", "sib"},
	"simxu": []string{"si'u", "sim"},
	"stici": []string{"sic"},
	"stidi": []string{"sid", "ti'i"},
	"sigja": []string{"sig"},
	"skiji": []string{"sij"},
	"silka": []string{"sik"},
	"siclu": []string{"sil"},
	"tsina": []string{"sin"},
	"sipna": []string{"sip"},
	"sirji": []string{"sir"},
	"sisku": []string{"sis"},
	"sitna": []string{"sit"},
	"sivni": []string{"siv"},
	"sirxo": []string{"six"},
	"si'o":  []string{"siz"},
	"skari": []string{"ska"},
	"saske": []string{"ske"},
	"skicu": []string{"ski"},
	"skori": []string{"sko"},
	"salci": []string{"sla"},
	"selci": []string{"sle"},
	"slilu": []string{"sli"},
	"solji": []string{"slo"},
	"sluji": []string{"slu"},
	"smaji": []string{"sma"},
	"semto": []string{"sme"},
	"simsa": []string{"smi"},
	"smoka": []string{"smo"},
	"sance": []string{"sna"},
	"sinxa": []string{"sni"},
	"masno": []string{"sno"},
	"casnu": []string{"snu"},
	"sovda": []string{"so'a", "sov"},
	"sobde": []string{"so'e", "sob"},
	"so'i":  []string{"so'i", "sor"},
	"sombo": []string{"so'o", "som"},
	"sorcu": []string{"soc", "sro"},
	"sodva": []string{"sod"},
	"softo": []string{"sof"},
	"sorgu": []string{"sog"},
	"sonci": []string{"soi", "son"},
	"so'a":  []string{"soj"},
	"solri": []string{"sol"},
	"so'e":  []string{"sop"},
	"so'o":  []string{"sos"},
	"so'u":  []string{"sot"},
	"so":    []string{"soz"},
	"spati": []string{"spa"},
	"speni": []string{"spe"},
	"spisa": []string{"spi"},
	"daspo": []string{"spo"},
	"spuda": []string{"spu"},
	"sarji": []string{"sra"},
	"srera": []string{"sre"},
	"dasri": []string{"sri"},
	"stali": []string{"sta"},
	"liste": []string{"ste"},
	"sisti": []string{"sti"},
	"stodi": []string{"sto"},
	"stuzi": []string{"stu", "tuz"},
	"stura": []string{"su'a", "tur"},
	"su'e":  []string{"su'e", "sup"},
	"sumti": []string{"su'i", "sum"},
	"su'o":  []string{"su'o", "suz"},
	"sfubu": []string{"su'u", "sub"},
	"sucta": []string{"suc"},
	"sudga": []string{"sud"},
	"sunga": []string{"sug"},
	"sumji": []string{"suj"},
	"suksa": []string{"suk"},
	"sunla": []string{"sul"},
	"stuna": []string{"sun"},
	"surla": []string{"sur"},
	"sutra": []string{"sut"},
	"su'u":  []string{"suv"},
	"tavla": []string{"ta'a", "tav"},
	"tanxe": []string{"ta'e", "tax"},
	"tatpi": []string{"ta'i"},
	"tanbo": []string{"ta'o"},
	"taxfu": []string{"ta'u", "taf"},
	"tabno": []string{"tab"},
	"tance": []string{"tac"},
	"tadni": []string{"tad"},
	"tagji": []string{"tag"},
	"tarmi": []string{"tai", "tam"},
	"tamji": []string{"taj"},
	"staku": []string{"tak"},
	"talsa": []string{"tal"},
	"tsani": []string{"tan"},
	"stapa": []string{"tap"},
	"tarci": []string{"tar"},
	"tansi": []string{"tas"},
	"tatru": []string{"tat"},
	"tanru": []string{"tau"},
	"ta":    []string{"taz"},
	"tcadu": []string{"tca"},
	"mutce": []string{"tce"},
	"tutci": []string{"tci"},
	"nitcu": []string{"tcu"},
	"terpa": []string{"te'a", "tep"},
	"steci": []string{"te'i", "tec"},
	"stero": []string{"te'o"},
	"tengu": []string{"te'u", "teg"},
	"ctebi": []string{"teb"},
	"terdi": []string{"ted"},
	"tenfa": []string{"tef"},
	"temci": []string{"tei", "tem"},
	"stela": []string{"tel"},
	"tcena": []string{"ten"},
	"te":    []string{"ter"},
	"terto": []string{"tet"},
	"tcima": []string{"ti'a", "tim"},
	"ctino": []string{"ti'o"},
	"tixnu": []string{"ti'u", "tix"},
	"tinbe": []string{"tib"},
	"tcica": []string{"tic"},
	"tcidu": []string{"tid"},
	"ti":    []string{"tif"},
	"tigni": []string{"tig"},
	"tilju": []string{"tij"},
	"stika": []string{"tik"},
	"tcila": []string{"til"},
	"tirna": []string{"tin"},
	"tikpa": []string{"tip"},
	"tirse": []string{"tir"},
	"tisna": []string{"tis"},
	"titla": []string{"tit"},
	"tivni": []string{"tiv"},
	"stizu": []string{"tiz"},
	"tonga": []string{"to'a", "tog"},
	"to'e":  []string{"to'e", "tol"},
	"torni": []string{"to'i", "ton"},
	"tordu": []string{"to'u", "tor"},
	"troci": []string{"toc", "toi"},
	"toldi": []string{"tod"},
	"toknu": []string{"tok"},
	"tarti": []string{"tra"},
	"mitre": []string{"tre"},
	"trina": []string{"tri"},
	"jitro": []string{"tro"},
	"turni": []string{"tru"},
	"tsali": []string{"tsa"},
	"zutse": []string{"tse", "zut"},
	"tsiju": []string{"tsi"},
	"tumla": []string{"tu'a", "tum"},
	"tuple": []string{"tu'e", "tup"},
	"tugni": []string{"tu'i", "tug"},
	"tunlo": []string{"tu'o", "tul"},
	"tubnu": []string{"tu'u"},
	"tunba": []string{"tub"},
	"tu":    []string{"tuf"},
	"tujli": []string{"tuj"},
	"tunka": []string{"tuk"},
	"tunta": []string{"tun"},
	"tutra": []string{"tut"},
	"vamji": []string{"va'i", "vam"},
	"vasxu": []string{"va'u", "vax"},
	"vanbi": []string{"vab"},
	"vanci": []string{"vac"},
	"vajni": []string{"vai", "vaj"},
	"valsi": []string{"val", "vla"},
	"vanju": []string{"van"},
	"vacri": []string{"var"},
	"vasru": []string{"vas", "vau"},
	"vamtu": []string{"vat"},
	"va":    []string{"vaz"},
	"verba": []string{"ve'a", "ver"},
	"ve'e":  []string{"ve'e"},
	"vecnu": []string{"ve'u", "ven"},
	"venfu": []string{"vef"},
	"ve":    []string{"vel"},
	"vensa": []string{"ves"},
	"viska": []string{"vi'a", "vis"},
	"vitke": []string{"vi'e"},
	"vikmi": []string{"vi'i", "vim"},
	"vitno": []string{"vi'o"},
	"vimcu": []string{"vi'u", "vic"},
	"vibna": []string{"vib"},
	"vindu": []string{"vid"},
	"vifne": []string{"vif"},
	"vinji": []string{"vij"},
	"viknu": []string{"vik"},
	"vlile": []string{"vil"},
	"vipsi": []string{"vip"},
	"vidru": []string{"vir"},
	"vitci": []string{"vit"},
	"vi":    []string{"viz"},
	"zivle": []string{"vle", "ziv"},
	"vlipa": []string{"vli"},
	"voksa": []string{"vo'a", "vok"},
	"vofli": []string{"voi", "vol"},
	"vo":    []string{"von"},
	"vorme": []string{"vor", "vro"},
	"vraga": []string{"vra"},
	"vreta": []string{"vre"},
	"virnu": []string{"vri"},
	"vrude": []string{"vu'e", "vud"},
	"vrusi": []string{"vu'i", "vus"},
	"vukro": []string{"vu'o", "vur"},
	"vu":    []string{"vuz"},
	"xatra": []string{"xa'a"},
	"xance": []string{"xa'e", "xan"},
	"xarci": []string{"xa'i", "xac"},
	"xampo": []string{"xa'o", "xap"},
	"xabju": []string{"xa'u"},
	"xadba": []string{"xab"},
	"xadni": []string{"xad"},
	"xamgu": []string{"xag", "xau"},
	"xrani": []string{"xai"},
	"xarju": []string{"xaj"},
	"xaksu": []string{"xak"},
	"xalka": []string{"xal"},
	"xajmi": []string{"xam"},
	"xanri": []string{"xar"},
	"xamsi": []string{"xas"},
	"xatsi": []string{"xat"},
	"xa":    []string{"xav"},
	"xazdo": []string{"xaz", "zdo"},
	"xedja": []string{"xe'a", "xej"},
	"xekri": []string{"xe'i", "xek"},
	"xendo": []string{"xe'o", "xed"},
	"xenru": []string{"xe'u", "xer"},
	"xebni": []string{"xei", "xen"},
	"xe":    []string{"xel"},
	"xelso": []string{"xes"},
	"xexso": []string{"xex"},
	"xirma": []string{"xi'a", "xir"},
	"xriso": []string{"xi'o", "xis"},
	"xislu": []string{"xi'u", "xil"},
	"xinmo": []string{"xim"},
	"xindo": []string{"xin"},
	"xispo": []string{"xip"},
	"xlali": []string{"xla"},
	"xlura": []string{"xlu"},
	"xotli": []string{"xoi", "xol"},
	"maxri": []string{"xri"},
	"xruti": []string{"xru"},
	"xusra": []string{"xu'a", "xus"},
	"xunre": []string{"xu'e", "xun"},
	"xukmi": []string{"xu'i", "xum"},
	"xurdo": []string{"xu'o", "xur"},
	"xruba": []string{"xub"},
	"xruki": []string{"xuk"},
	"xutla": []string{"xul"},
	"zabna": []string{"za'a", "zan"},
	"zasti": []string{"za'i", "zat"},
	"za'o":  []string{"za'o"},
	"zargu": []string{"za'u", "zag"},
	"zarci": []string{"zac", "zai"},
	"zajba": []string{"zaj"},
	"zalvi": []string{"zal"},
	"zanru": []string{"zar", "zau"},
	"zasni": []string{"zas"},
	"za'i":  []string{"zaz"},
	"zbasu": []string{"zba"},
	"zbepi": []string{"zbe"},
	"zdani": []string{"zda"},
	"zdile": []string{"zdi"},
	"zenba": []string{"ze'a", "zen"},
	"ze'e":  []string{"ze'e"},
	"ze'o":  []string{"ze'o", "zev"},
	"zekri": []string{"zei", "zer"},
	"ze":    []string{"zel"},
	"zepti": []string{"zep"},
	"zetro": []string{"zet"},
	"zgana": []string{"zga"},
	"zifre": []string{"zi'e", "zif"},
	"zinki": []string{"zi'i", "zin"},
	"dzipo": []string{"zi'o", "zip"},
	"zirpu": []string{"zi'u", "zir"},
	"zmiku": []string{"zmi"},
	"zumri": []string{"zmu"},
	"zo'a":  []string{"zo'a", "zon"},
	"zo'i":  []string{"zo'i", "zor"},
	"zukte": []string{"zu'e", "zuk"},
	"zunti": []string{"zu'i", "zun"},
	"zungi": []string{"zug"},
	"zunle": []string{"zul"},
	"zu'o":  []string{"zum"},
	"zvati": []string{"zva"},
}
