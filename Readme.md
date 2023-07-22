Radi liobitelyskovo interesa mne byilo liobopyitno predstavity kak dubet vyigliadety tekst na russkom yazyike
yesli byi on iznacalyno (ili po krayney mere mnogo vekov) ispolyzoval latinskiye bukvyi a ne cyrillic̹u.

## Sootvetstviye s cyrillic̹ey

| latinic̹a | cyrillic̹a |
| - | - |
| _c_  | _ч_ |
| _c̹_  | _ц_ |
| _x_  | _ш_ |
| _x̹_  | _щ_ |
| _y̆_  | _й_ |
| _y_  | _ь_ |
| _yı_ | _ы_ |
| _ıo_ | _ю_ |
| _ıa_ | _я_ |
| _ё_  | _ё_ |
| _j_  | _ж_ |
| _h_  | _х_ |

V ostalynom sootvetstviye standardnoye.

Glasnyiye _ë_, _ıo_, _ıa_, _yı_ upotrebliayutsa tolyko posle soglasnyih.
V otliciye ot ih cyrilliceskih ekvivalentov bukvyi _ë_, _ıo_, _ıa_ ispolyzuyutsa tolyko dlia oboznaceniya miagkosti predidux̹ih soglasnyih.
V nacale slov i posle glasnyih pixem _y̆o_, _y̆u_, _y̆a_.
Takje i bukva _e_ ne upotrebliayetsa v nacale slov i posle glasnyih za redkimi isklioceniyami
(_eto_, _poetomu_), vmesto etovo pixem _y̆e_.
Bukva _ъ_ ne imeyet ekvivalenta vovse tak kak yeyo funkc̹iya ne trebuyetsa.
Naprimer: _y̆abloko_, _priy̆atnyıy̆_, _y̆olka_, _y̆ujnyıy̆_, _y̆esli_, _hozıay̆in_, _v zdaniy̆i_, _oty̆ehaty_, _oby̆om_, _vy̆uga_. Boleye obyomnyiy primer ![1](/starobinets.pdf).

## Uprox̹onnyiye formyi

Dopuskayutsa sleduyux̹iye uprox̹onnyiye formyi:
| ishodnaya | uprox̹onnaya  |
| -  | -  |
| y̆  | y  |
| ë  | e  |
| ıa | ia |
| ıo | io |
| yı | yi |

Naprimer, sravnite

> Pod kopyıta moy̆ey̆ loxadi brosilsa nix̹iy̆, vopıa, cto grıadët konec̹ sveta i y̆a doljen pokay̆atsa i otdaty y̆emu vse denygi. Odnako, ponıav, cto y̆a ne otlicay̆usy osoboy̆ nabojnosty̆u, on tut je zabyıl obo mne i pristal k dvum dorodnyım kupc̹am, kotoryıy̆e byıli neskolyko perepuganyı tem bezumiy̆em, cto proishodilo vokrug.

s

> Pod kopyita moyey loxadi brosilsa nix̹iy, vopia, cto griadet konec̹ sveta i ya doljen pokayatsa i otdaty yemu vse denygi. Odnako, poniav, cto ya ne otlicayusy osoboy nabojnostyu, on tut je zabyil obo mne i pristal k dvum dorodnyim kupc̹am, kotoryiye byili neskolyko perepuganyi tem bezumiyem, cto proishodilo vokrug.

Boleye obyomnyiy primer ![2](/pehov.pdf).

## Istocniki primerov

- 1: otryivok iz _Anna Starobinets ``X̹ipac''_
- 2: otryivok iz _Aleksey Pehov ``Zolotyiye kostryi''_

## Techniceskiye detali

Oba primera sdelanyi s ispolyzovaniyem Golang 1.18 and Tex Live 2023.

```
go run . -i starobinets > starobinets.tex && xelatex starobinets.tex
go run . -i pehov -f 10 -s > pehov.tex && xelatex pehov.tex
```
