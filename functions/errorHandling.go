package functions

import(
	"os"
)

func ErrorHandling() (message string){
	f, err := os.Open("README.md")
    if err != nil {
		message = "open READMeE.md: The system cannot find the file specified."
        return
    }
	message =  f.Name()+": Opened successfully"
	return 
}