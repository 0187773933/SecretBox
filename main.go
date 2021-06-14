package main

import (
	"fmt"
	secretbox "github.com/0187773933/SecretBox/v1/secretbox"
)

func main() {

	// box := secretbox.New()
	// fmt.Println( box.HumanString )
	// fmt.Println( box.SealMessage( "waduwaduwadu" ) )
	// box.OpenFile( "/Users/morpheous/WORKSPACE/GO/SecretBox/test" )

	box := secretbox.Load( "jDDO/Ew7GdQQPKwFuz3PTZ3I6atG5mjaXlJ8GBV4LOVAQEA9PT1AQEBkfmJ5IjmBirkxgH7/VEpqeuQ7FxtmCUc=" )
	fmt.Println( box.OpenMessage( "ZH5ieSI5gYq5MYB+/1RKanrkOxcbZglHEbrMGPsaeEPqrwSQznVvf4YzUhkzRE459lNrzg==" ) )
}