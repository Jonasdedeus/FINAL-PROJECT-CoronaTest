package main

import (
	"fmt"
	"math/rand"
	"time"
)

const nMax int = 100

type date = struct {
	day   int
	month string
	year  int
}
type register = struct {
	ID          int
	name        string
	dob         date
	nationality string
	height      int
	weight      int
}

type checkup = struct {
	ID, age            int
	temperature, comma int
	name, result       string
}

type cases = struct {
	country       string
	totalpositive int
	totalRecover  int
	totalCases    int
	tP, tR, tC    int
}

type history = struct {
	name, result string
}

type arrayListRegister = struct {
	table        [nMax]register
	checkedUp    [nMax]checkup
	tableCase    [nMax]cases
	tablehistory [nMax]history
	n, m, o, p   int
}

func consultation(name string, ID, age, temperature, comma int, result string, arr *arrayListRegister) {
	//description: This procedure is to input the data when the patient is doing the test.
	var check bool
	check = true
	i := 0
	for check {
		if arr.checkedUp[i].ID != 0 {
			i++
		} else {
			arr.checkedUp[i].ID = ID
			arr.checkedUp[i].name = name
			arr.checkedUp[i].age = age
			arr.checkedUp[i].result = result
			arr.checkedUp[i].temperature = temperature
			arr.checkedUp[i].comma = comma
			arr.m = arr.m + 1
			check = false
		}
	}
}

func registration(name string, arr *arrayListRegister) {
	//This Procedure is to do the registration for new patient
	var found bool
	found = true
	i := 0
	for found {
		if arr.table[i].name != "" {
			i++
		} else {
			arr.table[i].ID = 1001 + i
			arr.table[i].name = name
			fmt.Print("Date of Birth: ")
			fmt.Scanln(&arr.table[i].dob.day, &arr.table[i].dob.month, &arr.table[i].dob.year)
			fmt.Print("Nationality: ")
			fmt.Scanln(&arr.table[i].nationality)
			fmt.Print("Height: ")
			fmt.Scanln(&arr.table[i].height)
			fmt.Print("Weight: ")
			fmt.Scanln(&arr.table[i].weight)
			arr.n = i + 1
			found = false
		}
	}
	fmt.Println("<>---------------------------<>")
	fmt.Println("Notice: You are registered")
	fmt.Println("<>===========================<>")
	fmt.Println(" ")
}

func resultTest(i int, arr *arrayListRegister) {
	//This procedure prints the result of patient who have already done the test.
	fmt.Println("[<>]=======================================[<>]")
	fmt.Println("\t\t RESULT")
	fmt.Println("_______________________________________________")
	fmt.Println("    ID | Name | Age | Temperature | Result")
	fmt.Println("    ", arr.checkedUp[i].ID, "|", arr.checkedUp[i].name, "|", arr.checkedUp[i].age, "|", arr.checkedUp[i].temperature, ",", arr.checkedUp[i].comma, "|", arr.checkedUp[i].result)
	fmt.Println("[<>]=======================================[<>]")
	fmt.Println("\n")
}

func regList(arr *arrayListRegister) {
	//This Procedure Prints the registration list after the new patient do the registration.
	fmt.Println("<>=====================================================<>")
	fmt.Println("\t\tLIST OF REGISTRATION")
	fmt.Println("<>=====================================================<>")
	fmt.Println("   ID | Name | Date of Birth | Height | weight")
	for i := 0; i < arr.n; i++ {
		fmt.Println("  ", arr.table[i].ID, "|", arr.table[i].name, "|", arr.table[i].dob, "|", arr.table[i].nationality, "|", arr.table[i].height, "|", arr.table[i].weight)
	}
	fmt.Println("<>=====================================================<>\n\n")

}

func swapTest(result string, temperature, comma, index int, arr *arrayListRegister) {
	//This Procedure will replace the data from patient who did the test for the next time.
	arr.checkedUp[index].temperature = temperature
	arr.checkedUp[index].comma = comma
	arr.checkedUp[index].result = result
}

func sort(chs int, arr *arrayListRegister) {
	//This Procedure is doing the sortof list corona cases.
	var mindex int
	for i := 0; i < arr.o; i++ {
		mindex = i
		for j := i + 1; j < arr.o; j++ {
			if chs == 1 {
				if arr.tableCase[j].country < arr.tableCase[mindex].country {
					mindex = j
				}
			} else if chs == 2 {
				if arr.tableCase[j].totalRecover > arr.tableCase[mindex].totalRecover {
					mindex = j
				}
			} else if chs == 3 {
				if arr.tableCase[j].totalpositive > arr.tableCase[mindex].totalpositive {
					mindex = j
				}
			} else {
				if arr.tableCase[j].totalCases > arr.tableCase[mindex].totalCases {
					mindex = j
				}
			}
		}

		temp := arr.tableCase[mindex]
		arr.tableCase[mindex] = arr.tableCase[i]
		arr.tableCase[i] = temp
	}
}

func insertHistory(name, result string, arr *arrayListRegister) {
	//This procedure will insert the history data from the test result
	var found bool
	found = true
	i := 0
	for found {
		if arr.tablehistory[i].name != "" {
			i++
		} else {
			arr.tablehistory[i].name = name
			arr.tablehistory[i].result = result
			arr.p = i + 1
			found = false
		}
	}
}

func swapHistory(result string, index int, arr *arrayListRegister) {
	//This procedure will replace the result of the next test in the history.
	arr.tablehistory[index].result = result
}
func printCase(arr *arrayListRegister) {
	//This procedure prints all of the cases from each country.
	var (
		totalCases    int
		totalpositive int
		totalRecover  int
	)
	fmt.Println("[<>]==============================...CORONA CASES...=================================[<>]")
	fmt.Println("   COUNTRY\t   ||\tTotal Cases\t||\tTotal Recover\t||\tTOTAL POSITIVE\t")
	fmt.Println("  -------------------------------------------------------------------------------------")
	for i := 0; i < arr.o; i++ {
		fmt.Println("  ", arr.tableCase[i].country, "\t   ||\t", arr.tableCase[i].totalRecover+arr.tableCase[i].totalpositive, "\t\t||\t", arr.tableCase[i].totalRecover, "\t\t||\t", arr.tableCase[i].totalpositive)
		fmt.Println("  -------------------------------------------------------------------------------------")
		arr.tableCase[i].totalCases = arr.tableCase[i].totalRecover + arr.tableCase[i].totalpositive
		arr.tableCase[i].tC = arr.tableCase[i].totalCases
		arr.tableCase[i].tP = arr.tableCase[i].totalpositive
		arr.tableCase[i].tR = arr.tableCase[i].totalRecover
		if i > 0 {
			arr.tableCase[i].tR = arr.tableCase[i].tR + arr.tableCase[i-1].tR
			arr.tableCase[i].tP = arr.tableCase[i].tP + arr.tableCase[i-1].tP
			arr.tableCase[i].tC = arr.tableCase[i].tC + arr.tableCase[i-1].tC
		}
		totalRecover = arr.tableCase[i].tR
		totalpositive = arr.tableCase[i].tP
		totalCases = arr.tableCase[i].tC
	}
	fmt.Println("  -------------------------------------------------------------------------------------")
	fmt.Println("   WORLD\t   ||\t", totalCases, "\t\t||\t", totalRecover, "\t\t||\t", totalpositive)
	fmt.Println("[]===================================================================================[]")
}

func searchPatient(chs int, name string, arr *arrayListRegister) bool {
	//This function will return true if it finds the patient in every list by name otherwise return false
	var found bool
	found = false
	i := 0
	if chs == 1 {
		for i < nMax-1 && !found {
			if arr.table[i].name == name {
				found = true
			}
			i++
		}
	} else if chs == 2 {
		for i < nMax-1 && !found {
			if arr.checkedUp[i].name == name {
				found = true
			}
			i++
		}
	} else {
		for i < nMax-1 && !found {
			if arr.tablehistory[i].name == name {
				found = true
			}
			i++
		}
	}
	return found
}

func searchIndex(chs int, name string, arr *arrayListRegister) int {
	//This function will return index when it finds the pasient in every list by name.
	var found bool
	var index int
	found = false
	i := 0
	if chs == 1 {
		for i < nMax-1 && !found {
			if arr.table[i].name == name {
				index = i
				found = true
			}
			i++
		}
	} else if chs == 2 {
		for i < nMax-1 && !found {
			if arr.checkedUp[i].name == name {
				index = i
				found = true
			}
			i++
		}
	} else {
		for i < nMax-1 && !found {
			if arr.tablehistory[i].name == name {
				index = i
				found = true
			}
			i++
		}
	}
	return index
}

func payment(name, test string, arr *arrayListRegister) {
	//This procedure will do the payment for the patient who have done registration.
	var (
		change, price     int
		checkPrice, found bool
		discount          float64
	)
	checkPrice, found = true, true
	fmt.Println("\n||_________________________________________________________||")
	fmt.Println("  xxxxxxxxxxxxxxxxxxxxx-- PAYMENT --xxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Println("||_________________________________________________________||")
	if test == "PCR" {
		fmt.Println("  Notice:")
		fmt.Println("  1.PCR/SWAB Test cost Rp 1.500.000")
		fmt.Println("  2.IF you already test in here, have discount 30%")
		fmt.Println("=============================================================")
		for checkPrice {
			fmt.Print("Amount: ")
			fmt.Scanln(&price)
			fmt.Println("-----------------------------")
			if price < 1500000 {
				fmt.Println("Your cash is not enough")
			} else {
				i := 0
				for i < nMax-1 && found {
					if arr.checkedUp[i].name == name {
						discount = 1500000 * 0.7
						change = price - int(discount)
						found = false
					}
					i++
				}
				if found == true {
					change = price - 1500000
				}
				checkPrice = false
			}
		}
	} else {
		fmt.Println("  Notice:")
		fmt.Println("  1. ANTIBODY Test cost Rp 300.000")
		fmt.Println("  2. If you already test in here, have discount 30%")
		fmt.Println("==================================================")
		for checkPrice {
			fmt.Print("Amount: ")
			fmt.Scanln(&price)
			fmt.Println("-------------------------------")
			if price < 300000 {
				fmt.Println("Your cash is not enough")
			} else {
				i := 0
				for i < nMax-1 && found {
					if arr.checkedUp[i].name == name {
						discount = 300000 * 0.7
						change = price - int(discount)
						found = false
					}
					i++
				}
				if found == true {
					change = price - 300000
				}
				checkPrice = false
			}
		}
	}
	fmt.Println("Your change is: ", change)
	fmt.Println("||_________________________________________________________||\n")
}

func TestCorona(checkup1 string, temperature int, test string, arr *arrayListRegister) string {
	//This procedure will return the result of patient after the test
	var (
		result, a, b, c, d, e int
		res                   string
	)
	if test == "PCR" {
		fmt.Println("[<>]==================PCR/SWAB TEST===================[<>]")
		fmt.Println("------------------------------------------------------------")
		fmt.Print("   || 1.Do you have fever?1.yes or 0.No: ")
		fmt.Scanln(&a)
		fmt.Print("   || 2.Do you cough? 1.yes or 0.No: ")
		fmt.Scanln(&b)
		fmt.Print("   || 3.Do you have headache? 1.yes or 0.No: ")
		fmt.Scanln(&c)
		fmt.Print("   || 4.Do you have sore throat? 1.yes or 0.No: ")
		fmt.Scanln(&d)
		fmt.Print("   || 5.Do you watery nose? 1.yes or 0.No: ")
		fmt.Scanln(&e)
		fmt.Println("[<>]==================================================[<>]\n")
		result = a + b + c + d + e
	} else {
		fmt.Println("[<>]==================================================ANTI-BODY TEST================================================[<>]")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("   ||  1.Do you have a history of travel to a city infected in the Indonesian region within 14 days? 1.Yes or 0.No: ")
		fmt.Scanln(&a)
		fmt.Print("   ||  2.Do you have a history of travel to an infected country within 14 days? 1.Yes or 0.No: ")
		fmt.Scanln(&b)
		fmt.Print("   ||  3.Have you ever given a nurse or made close contact with someone with COVID-19 in 14 days? 1.Yes or 0.No: ")
		fmt.Scanln(&c)
		fmt.Print("   ||  4.Have you ever been close contact with someone who berpegian abroad in the last 14 days? 1.Yes or 0.No: ")
		fmt.Scanln(&d)
		fmt.Println("[<>]================================================================================================================= [<>]\n")
		result = a + b + c + d
	}
	if result > 3 {
		res = "REACTIVE"
	} else if result == 3 && temperature > 37 {
		res = "REACTIVE"
	} else if checkup1 == "N" || checkup1 == "n" {
		res = "NEGATIVE"
	} else {
		res = "NON-REACTIVE"
	}
	fmt.Println("Notice: You have done your test, you can check in the result menu\n")
	return res
}

func main() {
	//This function will call every procedure in the above to perform the test and insert data to the array
	var (
		arr                         arrayListRegister
		choose, typeTest            int
		checkup, checkup1, checkup2 string
		name, test                  string
		result                      string
		found                       bool
	)
	rand.Seed(time.Now().UnixNano())
	arr.n, arr.m = 0, 0
	arr.table[0].ID = 1001
	min := 35
	max := 40
	fmt.Println("[<>]==================================[<>]")
	fmt.Println(" ||                 MENU               ||")
	fmt.Println("[<>]==================================[<>]")
	fmt.Println("     1.Registration")
	fmt.Println("     2.Take test")
	fmt.Println("     3.Result ")
	fmt.Println("     4.List")
	fmt.Println("     0.Exit")
	fmt.Println("________________________________________")
	fmt.Print("   Choose: ")
	fmt.Scanln(&choose)
	for choose != 0 {
		switch choose {
		case 1:
			fmt.Println("<>===========================<>")
			fmt.Println("\t   Registration")
			fmt.Println("<>===========================<>")
			fmt.Print("Your name: ")
			fmt.Scanln(&name)
			if searchPatient(1, name, &arr) == false {
				registration(name, &arr)
			} else {
				fmt.Println("This name already exist")
			}
		case 2:
			fmt.Println("<>=========consultation==========<>")
			fmt.Print(" 1.Enter you name: ")
			fmt.Scanln(&name)
			age := 2020 - arr.table[searchIndex(1, name, &arr)].dob.year
			ID := arr.table[searchIndex(1, name, &arr)].ID
			if searchPatient(1, name, &arr) == true {
				comma := rand.Intn(10)
				temperature := rand.Intn(max-min) + min
				fmt.Println(" 2.Your temperature: ", temperature, ",", comma)
				fmt.Print(" 3.Have took test in another place before? Y/N (y/n): ")
				fmt.Scanln(&checkup1)
				if checkup1 == "Y" || checkup1 == "y" {
					fmt.Print(" 4.If yes, What type of test? 1.PCR or 2.Antibody? choose in number: ")
					fmt.Scanln(&typeTest)
					fmt.Print(" 5.The result is (-) or (+)? Give symbol: ")
					fmt.Scanln(&result)
					if typeTest == 1 || result == "+" {
						test = "PCR"
						payment(name, test, &arr)
						if searchPatient(2, name, &arr) == true {
							if arr.checkedUp[searchIndex(2, name, &arr)].result != "NEGATIVE" {
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							} else {
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							}
						} else {
							consultation(name, ID, age, temperature, comma, TestCorona(checkup1, temperature, test, &arr), &arr)
						}
					} else if typeTest == 2 || result == "-" {
						test = "ANTIBODY"
						payment(name, test, &arr)
						if searchPatient(2, name, &arr) == true {
							if arr.checkedUp[searchIndex(2, name, &arr)].result != "NEGATIVE" {
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							} else {
								checkup1 = "N"
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							}
						} else {
							checkup1 = "N"
							consultation(name, ID, age, temperature, comma, TestCorona(checkup1, temperature, test, &arr), &arr)
						}
					}
				} else if checkup1 == "N" || checkup1 == "n" {
					fmt.Print(" 4.Do you have any symptoms? Y/N (y/n): ")
					fmt.Scanln(&checkup)
					fmt.Print(" 5.Is there any person who live with you affected the COVID-19? Y/N (y/n): ")
					fmt.Scanln(&checkup2)
					if (checkup == "N" || checkup == "n") && (checkup2 == "N" || checkup2 == "n") {
						test = "ANTIBODY"
						payment(name, test, &arr)
						if searchPatient(2, name, &arr) == true {
							if arr.checkedUp[searchIndex(2, name, &arr)].result != "NEGATIVE" {
								checkup1 = "Y"
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							} else {
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							}
						} else {
							consultation(name, ID, age, temperature, comma, TestCorona(checkup1, temperature, test, &arr), &arr)
						}
					} else if (checkup == "Y" || checkup == "y") || (checkup2 == "Y" || checkup2 == "y") {
						test = "PCR"
						payment(name, test, &arr)
						if searchPatient(2, name, &arr) == true {
							if arr.checkedUp[searchIndex(2, name, &arr)].result != "NEGATIVE" {
								checkup1 = "Y"
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							} else {
								swapTest(TestCorona(checkup1, temperature, test, &arr), temperature, comma, searchIndex(2, name, &arr), &arr)
							}
						} else {
							consultation(name, ID, age, temperature, comma, TestCorona(checkup1, temperature, test, &arr), &arr)
						}
					}

				} else {
					fmt.Println("ERROR Please fill it again, your input is wrong")
				}
				found = false
				i := 0
				for !found {
					if arr.tableCase[i].country == "" {
						arr.tableCase[i].country = arr.table[searchIndex(1, name, &arr)].nationality
						if result == "+" && arr.checkedUp[searchIndex(2, name, &arr)].result == "NON-REACTIVE" {
							arr.tableCase[i].totalRecover++
						} else if arr.checkedUp[searchIndex(2, name, &arr)].result == "REACTIVE" {
							arr.tableCase[i].totalpositive++
						} else {
							arr.tableCase[i].totalRecover = 0
							arr.tableCase[i].totalpositive = 0
						}
						arr.o = i + 1
						found = true
					} else {
						for i < arr.o && !found {
							if arr.table[searchIndex(1, name, &arr)].nationality == arr.tableCase[i].country {
								if arr.checkedUp[searchIndex(2, name, &arr)].result == "REACTIVE" {
									if searchPatient(3, name, &arr) == true && arr.tablehistory[searchIndex(3, name, &arr)].result != "REACTIVE" {
										if searchPatient(3, name, &arr) == true && arr.tablehistory[searchIndex(3, name, &arr)].result == "NEGATIVE" {
											arr.tableCase[i].totalpositive++
										} else {
											arr.tableCase[i].totalpositive++
											arr.tableCase[i].totalRecover--
										}
									} else if searchPatient(3, name, &arr) == false {
										arr.tableCase[i].totalpositive++
									}
								} else if arr.checkedUp[searchIndex(2, name, &arr)].result == "NON-REACTIVE" {
									if searchPatient(3, name, &arr) == true && arr.tablehistory[searchIndex(3, name, &arr)].result != "NON-REACTIVE" {
										if searchPatient(3, name, &arr) == true && arr.tablehistory[searchIndex(3, name, &arr)].result == "NEGATIVE" {
											arr.tableCase[i].totalRecover++
										} else {
											arr.tableCase[i].totalpositive--
											arr.tableCase[i].totalRecover++
										}
									} else if searchPatient(3, name, &arr) == false {
										arr.tableCase[i].totalRecover++
									}
								}
								found = true
							} else {
								i++
							}
						}
						if found == false {
							arr.tableCase[i].country = arr.table[searchIndex(1, name, &arr)].nationality
							if result == "+" && arr.checkedUp[searchIndex(2, name, &arr)].result == "NON-REACTIVE" {
								arr.tableCase[i].totalRecover++
							} else if arr.checkedUp[searchIndex(2, name, &arr)].result == "REACTIVE" {
								arr.tableCase[i].totalpositive++
							}
							arr.o = i + 1
							found = true
						} else {
							found = true
						}
					}
					var res string = arr.checkedUp[searchIndex(2, name, &arr)].result
					if searchPatient(3, name, &arr) == true {
						swapHistory(res, searchIndex(3, name, &arr), &arr)
					} else {
						insertHistory(name, res, &arr)
					}

				}
			} else {
				fmt.Println("Notice: Your name have not registered yet")
			}
		case 3:
			fmt.Print("   Enter your name: ")
			fmt.Scanln(&name)
			if searchPatient(2, name, &arr) == true {
				resultTest(searchIndex(2, name, &arr), &arr)
			} else {
				fmt.Println("   Notice: You have not done the test yet.")
			}

		case 4:
			fmt.Println("<>---------SUB MENU 1-------------")
			fmt.Println("   1.List of Registration")
			fmt.Println("   2.List of Corona Case")
			fmt.Println("_______________________________")
			fmt.Print("   Choose: ")
			fmt.Scanln(&choose)
			switch choose {
			case 1:
				regList(&arr)
			case 2:
				var choose int
				printCase(&arr)
				fmt.Println("<>----------SUB MENU 2-------------")
				fmt.Println("   1.Sorting by Country")
				fmt.Println("   2.Sorting by Total Recovered")
				fmt.Println("   3.Sorting by Total Positive")
				fmt.Println("   4.Sorting by Total Cases")
				fmt.Println("____________________________________")
				fmt.Print("   Choose: ")
				fmt.Scanln(&choose)
				switch choose {
				case 1:
					sort(choose, &arr)
					printCase(&arr)
				case 2:
					sort(choose, &arr)
					printCase(&arr)
				case 3:
					sort(choose, &arr)
					printCase(&arr)
				case 4:
					sort(choose, &arr)
					printCase(&arr)
				default:
					fmt.Println("Wrong Choose")
				}
			default:
				fmt.Println("Wrong Choose")
			}
		default:
			fmt.Println("Wrong Choose")
		}
		fmt.Println("[<>]==================================[<>]")
		fmt.Println(" ||                 MENU               ||")
		fmt.Println("[<>]==================================[<>]")
		fmt.Println("     1.Registration")
		fmt.Println("     2.Take test")
		fmt.Println("     3.Result ")
		fmt.Println("     4.List")
		fmt.Println("     0.Exit")
		fmt.Println("________________________________________")
		fmt.Print("   Choose: ")
		fmt.Scanln(&choose)
	}
	fmt.Println("[<>]==================Thank you===================[<>]")
}
