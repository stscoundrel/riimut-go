# Riimut

Transform latin letters to runes &amp; vice versa. Go version.

Includes transformers for four main runic alphabets:

- Elder Futhark
- Younger Futhark
- Medieval Futhork
- Futhorc (Anglo-Frisian runes)


## Usage

Text to runes:
```go
package main

// Ships four runic dialects under riimut module.
import (
    "fmt"

    "github.com/stscoundrel/riimut-go/elderfuthark"
    "github.com/stscoundrel/riimut-go/futhorc"
    "github.com/stscoundrel/riimut-go/medievalfuthork"
    "github.com/stscoundrel/riimut-go/youngerfuthark"
)

func main() {
    // From Old Groms runestone.
    const content = "auk tani karþi kristna"
    youngerFuthark = youngerfuthark.LettersToRunes(content)
    fmt.Println(youngerFuthark) // ᛅᚢᚴ:ᛏᛅᚾᛁ:ᚴᛅᚱᚦᛁ:ᚴᚱᛁᛋᛏᚾᛅ

    // From 4th century axe in Jutland
    const content = "wagagastiz alu wihgu sikijaz aiþalataz"
    elderFuthark = elderfuthark.LettersToRunes(content)
    fmt.Println(elderFuthark) // ᚹᚨᚷᚨᚷᚨᛋᛏᛁᛉ:ᚨᛚᚢ:ᚹᛁᚻᚷᚢ:ᛋᛁᚲᛁᛃᚨᛉ:ᚨᛁᚦᚨᛚᚨᛏᚨᛉ

    // From Lord's Prayer, in Old Norse.
    const content = "Faðer uor som ast i himlüm, halgað warðe þit nama"
    medievalFuthork = medievalfuthork.LettersToRunes(content)
    fmt.Println(medievalFuthork) // ᚠᛆᚦᚽᚱ:ᚢᚮᚱ:ᛋᚮᛘ:ᛆᛋᛏ:ᛁ:ᚼᛁᛘᛚᚢᛘ,:ᚼᛆᛚᚵᛆᚦ:ᚠᛆᚱᚦᚽ:ᚦᛁᛏ:ᚿᛆᛘᛆ

    // From 8th century Franks Casket, in late West Saxon.
    const content = "fisc.flodu.ahofonferg | enberig |"
    futhorc = futhorc.LettersToRunes(content)
    fmt.Println(futhorc) // ᚠᛁᛋᚳ.ᚠᛚᚩᛞᚢ.ᚪᚻᚩᚠᚩᚾᚠᛖᚱᚷ:|:ᛖᚾᛒᛖᚱᛁᚷ:|
}
```

Runes to text:
```go
package main

// All four dialects contain RunesToLetters method.
import (
    "fmt",

    "github.com/stscoundrel/riimut-go/youngerfuthark"
)

func main() {
    const runicText = "ᛅᚢᚴ:ᛏᛅᚾᛁ:ᚴᛅᚱᚦᛁ:ᚴᚱᛁᛋᛏᚾᛅ"
    latinText = youngerfuthark.RunesToLetters(runicText)

    fmt.Println(latinText) // "auk tani karþi kristna"
}

```


#### What's in the name?

"Riimut" is the Finnish word for "runes".
