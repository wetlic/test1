package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := 0
	for caseFour := true; caseFour; i++ {
		//	i += 1
		rand.Seed(time.Now().UnixNano())
		var isHeistOn bool = true
		var eludedGuards int = rand.Intn(100)

		if eludedGuards >= 50 {
			isHeistOn = true
			fmt.Println(i, ": Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
		} else {
			fmt.Println(i, ": Plan a better disguise next time?")
			isHeistOn = false
		}
		fmt.Println(i, ": Status of isHeistOn is:", isHeistOn)

		var openedVault int = rand.Intn(100)

		if isHeistOn && openedVault >= 70 {
			fmt.Println(i, ": Grab and GO!")
		} else if isHeistOn {
			isHeistOn = false
			fmt.Println(i, ": vault can’t be opened")
		}

		var leftSafely int = rand.Intn(5)
		if isHeistOn {
			switch leftSafely {
			case 0:
				isHeistOn = false
				fmt.Println(i, ": Case ", leftSafely, " failed")
			case 1:
				isHeistOn = false
				fmt.Println(i, ": Turns out vault doors don't open from the inside... for case ", leftSafely)
			case 2:
				isHeistOn = false
				fmt.Println(i, ": Turns out vault doors don't open from the inside... for case ", leftSafely)
			case 3:
				isHeistOn = false
				fmt.Println(i, ": Turns out vault doors don't open from the inside... for case ", leftSafely)
			default:
				fmt.Println(i, ": Start the getaway car! for case ", leftSafely)

			}
			if isHeistOn {
				amtStolen := 10000 + rand.Intn(1000000)
				fmt.Println(i, ": amtStolen is :", amtStolen, " nach ", i, "Runden")
				caseFour = false
			}
		}
	}

}
