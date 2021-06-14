package main

import (
	"os"
	"fmt"
	secretbox "github.com/0187773933/SecretBox/v1/secretbox"
	cli "github.com/urfave/cli/v2"
)

func Test() {
	// box := secretbox.New()
	// fmt.Println( box.HumanString )
	// fmt.Println( box.SealMessage( "waduwaduwadu" ) )
	// box.SealFile( "test.txt" , "waduwaduwadu" )
	// box.OpenFile( "/Users/morpheous/WORKSPACE/GO/SecretBox/test" )
	// fmt.Println( box.OpenMessage( "ZH5ieSI5gYq5MYB+/1RKanrkOxcbZglHEbrMGPsaeEPqrwSQznVvf4YzUhkzRE459lNrzg==" ) )
	// fmt.Println( box.OpenFile( "/Users/morpheous/WORKSPACE/GO/SecretBox/test.txt" ) )

	// box := secretbox.Load( "jDDO/Ew7GdQQPKwFuz3PTZ3I6atG5mjaXlJ8GBV4LOVAQEA9PT1AQEBkfmJ5IjmBirkxgH7/VEpqeuQ7FxtmCUc=" )
	// fmt.Println( box.HumanString )
}

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md
// https://pkg.go.dev/github.com/urfave/cli/v2#Context.FlagNames
func main() {
	var imported_key string
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "key",
				Aliases: []string{ "k" } ,
				Usage: "Creates New SecretBox" ,
				Destination: &imported_key ,
			} ,
	  	} ,
		Commands: []*cli.Command {
			{
				Name: "new" ,
				Aliases: []string{ "n" } ,
				Usage: "Creates New SecretBox" ,
				Action: func( c *cli.Context ) error {
					box := secretbox.New()
					fmt.Println( box.HumanString )
					return nil
				} ,
			} ,
			{
				Name: "seal" ,
				Aliases: []string{ "s" } ,
				Usage: "seals a plain text message or file" ,
				Subcommands: []*cli.Command{
					{
						Name: "message",
						Aliases: []string{ "m" } ,
						Usage: "seals a plain text message" ,
						Action: func( c *cli.Context ) error {
							// flag_names := c.FlagNames()
							// arg_list := c.Args().Slice()
							box := secretbox.Load( imported_key )
							fmt.Println( box.SealMessage( c.Args().First() ) )
							return nil
						} ,
					} ,
					{
						Name: "file",
						Aliases: []string{ "f" } ,
						Usage: "seals a plain text file" ,
						Action: func( c *cli.Context ) error {
							arg_list := c.Args().Slice()
							box := secretbox.Load( imported_key )
							fmt.Println( box.SealFile( arg_list[ 0 ] , arg_list[ 1 ] ) )
							return nil
						} ,
					} ,
				} ,
			} ,
			{
				Name: "open" ,
				Aliases: []string{ "o" } ,
				Usage: "opens an encrypted text message or file" ,
				Subcommands: []*cli.Command{
					{
						Name: "message",
						Aliases: []string{ "m" } ,
						Usage: "opens an encrypted text message" ,
						Action: func( c *cli.Context ) error {
							box := secretbox.Load( imported_key )
							fmt.Println( box.OpenMessage( c.Args().First() ) )
							return nil
						} ,
					} ,
					{
						Name: "file",
						Aliases: []string{ "f" } ,
						Usage: "opens an encrypted text file" ,
						Action: func( c *cli.Context ) error {
							box := secretbox.Load( imported_key )
							fmt.Println( box.OpenFile( c.Args().First() ) )
							return nil
						} ,
					} ,
				} ,
			} ,
		} ,
	}
	app.Run( os.Args )
}