package main
import (
	"fmt"
	"log"
)

// calculate average grade 

func calculateAverage(grades []float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade;
	}

	return total / float64(len(grades))
}


func main()  {
	var name string
	var numSubjects int 

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your number of subjects: ")
	fmt.Scanln(&numSubjects)

	if numSubjects <= 0 {
		fmt.Println("The number of subjects must be greater than zero")
	}

	subjects := make([] string, numSubjects)
	grades := make([] float64, numSubjects)

	for i := 0; i < numSubjects; i ++{
		fmt.Println("Enter the name of the subject %d: ",  i + 1)
		fmt.Scanln(&subjects[i])

		fmt.Println("Enter the grade for %s: ", subjects[i])
		fmt.Scanln(&grades[i])

		if grades[i] < 0 || grades[i] > 100 {
			log.Fatalf("Grade for %s must be between 0 and 100.", subjects[i])
		}

	}
	
 averageGrade := calculateAverage(grades)
 fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Grades:")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.2f\n", subjects[i], grades[i])
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)

	
}