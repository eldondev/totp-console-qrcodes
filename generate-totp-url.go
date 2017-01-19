package main
import (
	"github.com/mdp/qrterminal"
	"os"
	"os/exec"
	"strings"
	"log"
	"flag"
	"fmt"
)

func main() {
	will_test := flag.Bool("iwilltestmycodes", false, "You promise to test your codes")
	database := flag.String("database", "databases", "name of the database file")
	flag.Parse()
	if !*will_test {
		log.Fatalf("You didn't promise to test your codes. This is alpha software, and involves hard-to-recover things. Promise by passing  -iwilltestmycodes")
	}
	urls, err := exec.Command("sh", "-c", fmt.Sprintf("echo 'select printf(\"otpauth://totp/%%s?secret=%%s&issuer=%%s\", email, secret, issuer) from accounts;'  |sqlite3 %s", *database)).Output()
  if err != nil {
		log.Fatalf("%s", err)
	}
  for _, url := range(strings.Split(string(urls), "\n")) {
	  if url != ""{
		fmt.Printf("\n\n\n")
		qrterminal.Generate(url, qrterminal.M, os.Stdout)
		fmt.Printf("\n\n\n\n\n\n\n\n")
		os.Stdin.Read([]byte{0x00})
		}
	}
}
