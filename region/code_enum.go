// Code generated by "enumer --type=Code --linecomment --json --extramethod --output=code_enum.go"; DO NOT EDIT.

package region

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeEC-1]
	_ = x[CodeBR-2]
	_ = x[CodeMX-3]
	_ = x[CodeCO-4]
	_ = x[CodeID-5]
	_ = x[CodePH-6]
	_ = x[CodePE-7]
	_ = x[CodeEG-8]
	_ = x[CodeTR-9]
	_ = x[CodeAR-10]
	_ = x[CodeCL-11]
	_ = x[CodeVE-12]
	_ = x[CodeBO-13]
	_ = x[CodeIN-14]
	_ = x[CodeJP-15]
	_ = x[CodeTH-16]
	_ = x[CodeMY-17]
	_ = x[CodeVN-18]
	_ = x[CodeRU-19]
	_ = x[CodeSG-20]
	_ = x[CodeBD-21]
	_ = x[CodeKR-22]
	_ = x[CodeAZ-23]
	_ = x[CodePA-24]
	_ = x[CodeJO-25]
	_ = x[CodeOM-26]
	_ = x[CodeKZ-27]
	_ = x[CodeSA-28]
	_ = x[CodeUK-29]
	_ = x[CodeDO-30]
	_ = x[CodeFK-31]
	_ = x[CodeGF-32]
	_ = x[CodeGY-33]
	_ = x[CodePY-34]
	_ = x[CodeSR-35]
	_ = x[CodeUY-36]
	_ = x[CodeCA-37]
	_ = x[CodeUS-38]
	_ = x[CodeBZ-39]
	_ = x[CodeCR-40]
	_ = x[CodeCU-41]
	_ = x[CodeSV-42]
	_ = x[CodeGT-43]
	_ = x[CodeHT-44]
	_ = x[CodeHN-45]
	_ = x[CodeJM-46]
	_ = x[CodeNI-47]
	_ = x[CodePR-48]
	_ = x[CodeTT-49]
	_ = x[CodeAL-50]
	_ = x[CodeAD-51]
	_ = x[CodeAT-52]
	_ = x[CodeBY-53]
	_ = x[CodeBE-54]
	_ = x[CodeBA-55]
	_ = x[CodeBG-56]
	_ = x[CodeHR-57]
	_ = x[CodeCY-58]
	_ = x[CodeCZ-59]
	_ = x[CodeDK-60]
	_ = x[CodeEE-61]
	_ = x[CodeFO-62]
	_ = x[CodeFI-63]
	_ = x[CodeFR-64]
	_ = x[CodeDE-65]
	_ = x[CodeGR-66]
	_ = x[CodeTW-67]
	_ = x[CodeAF-68]
	_ = x[CodeIR-69]
	_ = x[CodeIQ-70]
	_ = x[CodeKW-71]
	_ = x[CodeKG-72]
	_ = x[CodeLB-73]
	_ = x[CodeMV-74]
	_ = x[CodeMN-75]
	_ = x[CodeNP-76]
	_ = x[CodeUZ-77]
	_ = x[CodeYE-78]
	_ = x[CodeZA-79]
	_ = x[CodeIL-80]
	_ = x[CodeAU-81]
	_ = x[CodeNZ-82]
	_ = x[CodeIT-83]
	_ = x[CodeNG-84]
	_ = x[CodeMA-85]
	_ = x[CodeLK-86]
	_ = x[CodeUSV-87]
	_ = x[CodeGN-88]
	_ = x[CodeHK-89]
	_ = x[CodeAE-90]
	_ = x[CodePL-91]
	_ = x[CodeMC-92]
	_ = x[CodeIE-93]
	_ = x[CodeLA-94]
	_ = x[CodeSY-95]
	_ = x[CodeLV-96]
	_ = x[CodeNL-97]
	_ = x[CodePT-98]
	_ = x[CodeAM-99]
	_ = x[CodeES-100]
	_ = x[CodeHU-101]
	_ = x[CodeSK-102]
	_ = x[CodeSI-103]
	_ = x[CodeLT-104]
	_ = x[CodeGA-105]
	_ = x[CodeTL-106]
	_ = x[CodeRS-107]
	_ = x[CodeNO-108]
	_ = x[CodeIS-109]
	_ = x[CodeLU-110]
	_ = x[CodeME-111]
	_ = x[CodeTM-112]
	_ = x[CodeDZ-113]
	_ = x[CodeSE-114]
	_ = x[CodeRO-115]
	_ = x[CodeMO-116]
	_ = x[CodeML-117]
	_ = x[CodeBH-118]
	_ = x[CodeUA-119]
	_ = x[CodeUSP-120]
	_ = x[CodeAO-121]
	_ = x[CodeKE-122]
	_ = x[CodeTD-123]
	_ = x[CodeMS-124]
	_ = x[CodeCG-125]
	_ = x[CodeTN-126]
	_ = x[CodeMZ-127]
	_ = x[CodeSN-128]
	_ = x[CodePG-129]
	_ = x[CodeNA-130]
	_ = x[CodeET-131]
	_ = x[CodeLY-132]
	_ = x[CodeNE-133]
	_ = x[CodePK-134]
	_ = x[CodeLS-135]
	_ = x[CodeUG-136]
	_ = x[CodeZM-137]
	_ = x[CodeMG-138]
	_ = x[CodeMM-139]
	_ = x[CodeTG-140]
	_ = x[CodeGH-141]
	_ = x[CodeBW-142]
	_ = x[CodeGM-143]
	_ = x[CodeBS-144]
	_ = x[CodeFJ-145]
	_ = x[CodeSO-146]
	_ = x[CodeBJ-147]
	_ = x[CodeCI-148]
	_ = x[CodeKH-149]
	_ = x[CodeCM-150]
	_ = x[CodeBI-151]
	_ = x[CodeLR-152]
	_ = x[CodeGW-153]
	_ = x[CodeMK-154]
	_ = x[CodeAW-155]
	_ = x[CodeKM-156]
	_ = x[CodeZW-157]
	_ = x[CodeRW-158]
	_ = x[CodeLC-159]
	_ = x[CodeMW-160]
	_ = x[CodeMU-161]
	_ = x[CodeMR-162]
	_ = x[CodeSS-163]
	_ = x[CodeGQ-164]
	_ = x[CodeCF-165]
	_ = x[CodeSL-166]
	_ = x[CodeSZ-167]
	_ = x[CodeAI-168]
	_ = x[CodeCH-169]
	_ = x[CodeST-170]
	_ = x[CodeBB-171]
	_ = x[CodeSC-172]
	_ = x[CodeTZ-173]
	_ = x[CodeKN-174]
	_ = x[CodeNC-175]
	_ = x[CodeTJ-176]
	_ = x[CodeMD-177]
	_ = x[CodeGE-178]
	_ = x[CodeBF-179]
	_ = x[CodeCN-180]
	_ = x[CodeQA-181]
	_ = x[CodeBN-182]
	_ = x[CodeGP-183]
	_ = x[CodeER-184]
	_ = x[CodeAG-185]
	_ = x[CodeDJ-186]
	_ = x[CodeKY-187]
	_ = x[CodeGD-188]
	_ = x[CodeCV-189]
	_ = x[CodeVC-190]
	_ = x[CodeDM-191]
	_ = x[CodeSB-192]
	_ = x[CodeTO-193]
	_ = x[CodeBT-194]
	_ = x[CodeCD-195]
	_ = x[CodeLI-196]
	_ = x[CodeRE-197]
	_ = x[CodeWS-198]
	_ = x[CodeSX-199]
}

const _Code_name = "ECBRMXCOIDPHPEEGTRARCLVEBOINJPTHMYVNRUSGBDKRAZPAJOOMKZSAUKDOFKGFGYPYSRUYCAUSBZCRCUSVGTHTHNJMNIPRTTALADATBYBEBABGHRCYCZDKEEFOFIFRDEGRTWAFIRIQKWKGLBMVMNNPUZYEZAILAUNZITNGMALKUSVGNHKAEPLMCIELASYLVNLPTAMESHUSKSILTGATLRSNOISLUMETMDZSEROMOMLBHUAUSPAOKETDMSCGTNMZSNPGNAETLYNEPKLSUGZMMGMMTGGHBWGMBSFJSOBJCIKHCMBILRGWMKAWKMZWRWLCMWMUMRSSGQCFSLSZAICHSTBBSCTZKNNCTJMDGEBFCNQABNGPERAGDJKYGDCVVCDMSBTOBTCDLIREWSSX"

var _Code_index = [...]uint16{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 175, 177, 179, 181, 183, 185, 187, 189, 191, 193, 195, 197, 199, 201, 203, 205, 207, 209, 211, 213, 215, 217, 219, 221, 223, 225, 227, 229, 231, 233, 235, 237, 239, 242, 244, 246, 248, 250, 252, 254, 256, 258, 260, 262, 264, 266, 268, 270, 272, 274, 276, 278, 280, 282, 284, 286, 288, 290, 292, 294, 296, 298, 300, 302, 304, 306, 308, 310, 312, 314, 316, 318, 320, 322, 324, 326, 328, 330, 332, 334, 336, 338, 340, 342, 344, 346, 348, 350, 352, 354, 356, 358, 360, 362, 364, 366, 368, 370, 372, 374, 376, 378, 380, 382, 384, 386, 388, 390, 392, 394, 396, 398, 400}

func (i Code) String() string {
	i -= 1
	if i < 0 || i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}

var _Code_values = []Code{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199}

var _Code_name_to_value_map = map[string]Code{
	_Code_name[0:2]:     1,
	_Code_name[2:4]:     2,
	_Code_name[4:6]:     3,
	_Code_name[6:8]:     4,
	_Code_name[8:10]:    5,
	_Code_name[10:12]:   6,
	_Code_name[12:14]:   7,
	_Code_name[14:16]:   8,
	_Code_name[16:18]:   9,
	_Code_name[18:20]:   10,
	_Code_name[20:22]:   11,
	_Code_name[22:24]:   12,
	_Code_name[24:26]:   13,
	_Code_name[26:28]:   14,
	_Code_name[28:30]:   15,
	_Code_name[30:32]:   16,
	_Code_name[32:34]:   17,
	_Code_name[34:36]:   18,
	_Code_name[36:38]:   19,
	_Code_name[38:40]:   20,
	_Code_name[40:42]:   21,
	_Code_name[42:44]:   22,
	_Code_name[44:46]:   23,
	_Code_name[46:48]:   24,
	_Code_name[48:50]:   25,
	_Code_name[50:52]:   26,
	_Code_name[52:54]:   27,
	_Code_name[54:56]:   28,
	_Code_name[56:58]:   29,
	_Code_name[58:60]:   30,
	_Code_name[60:62]:   31,
	_Code_name[62:64]:   32,
	_Code_name[64:66]:   33,
	_Code_name[66:68]:   34,
	_Code_name[68:70]:   35,
	_Code_name[70:72]:   36,
	_Code_name[72:74]:   37,
	_Code_name[74:76]:   38,
	_Code_name[76:78]:   39,
	_Code_name[78:80]:   40,
	_Code_name[80:82]:   41,
	_Code_name[82:84]:   42,
	_Code_name[84:86]:   43,
	_Code_name[86:88]:   44,
	_Code_name[88:90]:   45,
	_Code_name[90:92]:   46,
	_Code_name[92:94]:   47,
	_Code_name[94:96]:   48,
	_Code_name[96:98]:   49,
	_Code_name[98:100]:  50,
	_Code_name[100:102]: 51,
	_Code_name[102:104]: 52,
	_Code_name[104:106]: 53,
	_Code_name[106:108]: 54,
	_Code_name[108:110]: 55,
	_Code_name[110:112]: 56,
	_Code_name[112:114]: 57,
	_Code_name[114:116]: 58,
	_Code_name[116:118]: 59,
	_Code_name[118:120]: 60,
	_Code_name[120:122]: 61,
	_Code_name[122:124]: 62,
	_Code_name[124:126]: 63,
	_Code_name[126:128]: 64,
	_Code_name[128:130]: 65,
	_Code_name[130:132]: 66,
	_Code_name[132:134]: 67,
	_Code_name[134:136]: 68,
	_Code_name[136:138]: 69,
	_Code_name[138:140]: 70,
	_Code_name[140:142]: 71,
	_Code_name[142:144]: 72,
	_Code_name[144:146]: 73,
	_Code_name[146:148]: 74,
	_Code_name[148:150]: 75,
	_Code_name[150:152]: 76,
	_Code_name[152:154]: 77,
	_Code_name[154:156]: 78,
	_Code_name[156:158]: 79,
	_Code_name[158:160]: 80,
	_Code_name[160:162]: 81,
	_Code_name[162:164]: 82,
	_Code_name[164:166]: 83,
	_Code_name[166:168]: 84,
	_Code_name[168:170]: 85,
	_Code_name[170:172]: 86,
	_Code_name[172:175]: 87,
	_Code_name[175:177]: 88,
	_Code_name[177:179]: 89,
	_Code_name[179:181]: 90,
	_Code_name[181:183]: 91,
	_Code_name[183:185]: 92,
	_Code_name[185:187]: 93,
	_Code_name[187:189]: 94,
	_Code_name[189:191]: 95,
	_Code_name[191:193]: 96,
	_Code_name[193:195]: 97,
	_Code_name[195:197]: 98,
	_Code_name[197:199]: 99,
	_Code_name[199:201]: 100,
	_Code_name[201:203]: 101,
	_Code_name[203:205]: 102,
	_Code_name[205:207]: 103,
	_Code_name[207:209]: 104,
	_Code_name[209:211]: 105,
	_Code_name[211:213]: 106,
	_Code_name[213:215]: 107,
	_Code_name[215:217]: 108,
	_Code_name[217:219]: 109,
	_Code_name[219:221]: 110,
	_Code_name[221:223]: 111,
	_Code_name[223:225]: 112,
	_Code_name[225:227]: 113,
	_Code_name[227:229]: 114,
	_Code_name[229:231]: 115,
	_Code_name[231:233]: 116,
	_Code_name[233:235]: 117,
	_Code_name[235:237]: 118,
	_Code_name[237:239]: 119,
	_Code_name[239:242]: 120,
	_Code_name[242:244]: 121,
	_Code_name[244:246]: 122,
	_Code_name[246:248]: 123,
	_Code_name[248:250]: 124,
	_Code_name[250:252]: 125,
	_Code_name[252:254]: 126,
	_Code_name[254:256]: 127,
	_Code_name[256:258]: 128,
	_Code_name[258:260]: 129,
	_Code_name[260:262]: 130,
	_Code_name[262:264]: 131,
	_Code_name[264:266]: 132,
	_Code_name[266:268]: 133,
	_Code_name[268:270]: 134,
	_Code_name[270:272]: 135,
	_Code_name[272:274]: 136,
	_Code_name[274:276]: 137,
	_Code_name[276:278]: 138,
	_Code_name[278:280]: 139,
	_Code_name[280:282]: 140,
	_Code_name[282:284]: 141,
	_Code_name[284:286]: 142,
	_Code_name[286:288]: 143,
	_Code_name[288:290]: 144,
	_Code_name[290:292]: 145,
	_Code_name[292:294]: 146,
	_Code_name[294:296]: 147,
	_Code_name[296:298]: 148,
	_Code_name[298:300]: 149,
	_Code_name[300:302]: 150,
	_Code_name[302:304]: 151,
	_Code_name[304:306]: 152,
	_Code_name[306:308]: 153,
	_Code_name[308:310]: 154,
	_Code_name[310:312]: 155,
	_Code_name[312:314]: 156,
	_Code_name[314:316]: 157,
	_Code_name[316:318]: 158,
	_Code_name[318:320]: 159,
	_Code_name[320:322]: 160,
	_Code_name[322:324]: 161,
	_Code_name[324:326]: 162,
	_Code_name[326:328]: 163,
	_Code_name[328:330]: 164,
	_Code_name[330:332]: 165,
	_Code_name[332:334]: 166,
	_Code_name[334:336]: 167,
	_Code_name[336:338]: 168,
	_Code_name[338:340]: 169,
	_Code_name[340:342]: 170,
	_Code_name[342:344]: 171,
	_Code_name[344:346]: 172,
	_Code_name[346:348]: 173,
	_Code_name[348:350]: 174,
	_Code_name[350:352]: 175,
	_Code_name[352:354]: 176,
	_Code_name[354:356]: 177,
	_Code_name[356:358]: 178,
	_Code_name[358:360]: 179,
	_Code_name[360:362]: 180,
	_Code_name[362:364]: 181,
	_Code_name[364:366]: 182,
	_Code_name[366:368]: 183,
	_Code_name[368:370]: 184,
	_Code_name[370:372]: 185,
	_Code_name[372:374]: 186,
	_Code_name[374:376]: 187,
	_Code_name[376:378]: 188,
	_Code_name[378:380]: 189,
	_Code_name[380:382]: 190,
	_Code_name[382:384]: 191,
	_Code_name[384:386]: 192,
	_Code_name[386:388]: 193,
	_Code_name[388:390]: 194,
	_Code_name[390:392]: 195,
	_Code_name[392:394]: 196,
	_Code_name[394:396]: 197,
	_Code_name[396:398]: 198,
	_Code_name[398:400]: 199,
}

// CodeFromStr 将字符串转换成 Code 枚举类型；如果传入的不是有效的枚举值，返回 0 值。
func CodeFromStr(s string) Code {
	if val, ok := _Code_name_to_value_map[s]; ok {
		return val
	}
	return 0
}

// CodeFromStrE 将字符串转换成 Code 枚举类型；如果传入的不是有效的枚举值，抛出错误。
func CodeFromStrE(s string) (Code, error) {
	if val, ok := _Code_name_to_value_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("not valid enum: %s", s)
}

// CodeFromInt 将int值转换成 Code 枚举类型；如果传入的不是有效的枚举值，返回 0 值。
func CodeFromInt(i int) Code {
	if _CodeIsValid(i) {
		return Code(i)
	}
	return 0
}

// CodeFromIntE 将int值转换成 Code 枚举类型；如果传入的不是有效的枚举值，抛出错误。
func CodeFromIntE(i int) (Code, error) {
	if _CodeIsValid(i) {
		return Code(i), nil
	}
	return 0, fmt.Errorf("not valid enum: %d", i)
}

// CodeFromInt64 将int64转换成 Code 枚举类型；如果传入的不是有效的枚举值，返回 0 值。
func CodeFromInt64(i int64) Code {
	return CodeFromInt(int(i))
}

// CodeFromInt64E 将int64转换成 Code 枚举类型；如果传入的不是有效的枚举值，抛出错误。
func CodeFromInt64E(i int64) (Code, error) {
	return CodeFromIntE(int(i))
}

// IsValid 判断该枚举值是否有效。
func (i Code) IsValid() bool {
	return _CodeIsValid(int(i))
}

// Check 校验该枚举值是否有效，如果无效，抛出错误。
func (i Code) Check() (err error) {
	if is := i.IsValid(); !is {
		err = fmt.Errorf("not valid enum: %v", i)
	}
	return err
}

type CodeSlice []Code

// Has 判断给定枚举值是否在当前数组中。
func (inst CodeSlice) Has(i Code) bool {
	for _, v := range inst {
		if i == v {
			return true
		}
	}
	return false
}

// Strings 将枚举数组转换为字符串数组。
func (inst CodeSlice) Strings() []string {
	ss := make([]string, len(inst))
	for i, v := range inst {
		ss[i] = v.String()
	}
	return ss
}

// CodeValues 返回 Code 类型的所有的枚举值。
func CodeValues() CodeSlice {
	return _Code_values
}

// CodeFromStrs 将字符串数组转换成 Code 枚举数组；如果传入的字符串不是有效的枚举值，抛出错误。
func CodeFromStrs(ss []string) (CodeSlice, error) {
	es := make(CodeSlice, 0, len(ss))
	for _, s := range ss {
		v, err := CodeFromStrE(s)
		if err != nil {
			return nil, err
		}
		es = append(es, v)
	}
	return es, nil
}

// CodeFromInts 将int数组转换成 Code 枚举数组；如果传入的不是有效的枚举值，抛出错误。
func CodeFromInts(ints []int) (CodeSlice, error) {
	es := make(CodeSlice, 0, len(ints))
	for _, i := range ints {
		v, err := CodeFromIntE(i)
		if err != nil {
			return nil, err
		}
		es = append(es, v)
	}
	return es, nil
}

// CodeFromInt64s 将int64数组转换成 Code 枚举数组；如果传入的不是有效的枚举值，抛出错误。
func CodeFromInt64s(int64s []int64) (CodeSlice, error) {
	ints := make([]int, 0, len(int64s))
	for _, i := range int64s {
		ints = append(ints, int(i))
	}
	return CodeFromInts(ints)
}

// _CodeIsValid 判断 int 是否是有效的枚举值
func _CodeIsValid(i int) bool {
	for _, v := range _Code_values {
		if i == int(v) {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Code
func (i Code) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Code
func (i *Code) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Code should be a string, got %v", data)
	}

	v, ok := _Code_name_to_value_map[s]
	if !ok {
		return fmt.Errorf("not valid enum, value is %v", s)
	}

	*i = v
	return nil
}
