package main

import (
	"github.com/amolabs/amoconsole/cmd"
)

/* Commands (expected hierarchy)
 *
 * amocli |- version
 *		  |- status
 * 		  |- key |- list
 *		  		 |- generate <nickname>
 *				 |- remove <nickname>
 *
 *		  |- tx |- transfer --from <address> --to <address> --amount <number>
 *		  		|- purchase --from <address> --file <hash>
 *
 *		  |- query
 */

func main() {
	cmd.Execute()
}
