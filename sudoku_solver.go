package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "os"
    "flag"
    )

type Sudoku struct{
     //The Sudoku-field. 0 means missing number.
     field        [9][9]int
}

func (s *Sudoku) Fill(filename string) error {

    //Read file
    b, err := ioutil.ReadFile(filename)
    if err != nil { return err }

    //fill Soduku field
    lines := strings.Split(string(b), "\n")

    for i := 0 ; i<9 ; i++ {
    	for j := 0 ; j<9 ; j++ {
	    str_field    := string(lines[i][j])
	    if str_field == " " { continue }
            int_field, err := strconv.ParseInt(str_field,10,0)
	    if err != nil { return err }
	    s.field[i][j] = int(int_field)
	}
    }
    return nil
}

func (s *Sudoku) check_consistency() bool {

    //check consistency in rows and cols
    for dim1 := 0 ; dim1<9 ; dim1++ {

        var found_dim1 [9]bool
	var found_dim2 [9]bool
	
        for dim2 := 0 ; dim2<9 ; dim2++ {
	    if s.field[dim1][dim2] > 0 {

	        if found_dim1[s.field[dim1][dim2]-1] {
                    return false
		} else {
		    found_dim1[s.field[dim1][dim2]-1] = true
		}
            }

            if s.field[dim2][dim1] > 0 {

	        if found_dim2[s.field[dim2][dim1]-1] {
                    return false
		} else {
		    found_dim2[s.field[dim2][dim1]-1] = true
		}	
	    }

        }
    }

    //check consistency in blocks
    for block_row := 1 ; block_row < 8 ; block_row += 3 {
        for block_col := 1 ; block_col < 8 ; block_col += 3 {
            var found [9]bool

            for i := -1 ; i < 2 ; i++ {
                for j := -1 ; j < 2 ; j++ {

                    if s.field[block_row+i][block_col+j] > 0 {

                        if found[s.field[block_row+i][block_col+j]-1] {
		            return false
		        } else {
		            found[s.field[block_row+i][block_col+j]-1] = true
                        }
                    }
                }
            }
        }
    }
    return true
}

func (s *Sudoku) add_number(row, col, value int) bool {

    for i := 0 ; i < 9 ; i++ {

        if s.field[row][i] == value { return false }
	if s.field[i][col] == value { return false }
    }

    block_row := int(row/3)*3
    block_col := int(col/3)*3
    for r := 0 ; r < 3 ; r++ {
        for c := 0 ; c < 3 ; c++ {

            if s.field[block_row+r][block_col+c] == value { return false }
        }
    }
    s.field[row][col] = value
    return true
}

func (s *Sudoku) next(row, col int) bool {
    if col < 8 {
        if s.solve(row, col +1) { return true }
    } else  {
        if s.solve(row+1,0) { return true }
    }
    return false
}

func (s *Sudoku) solve(row, col int) bool {
    if row > 8 { return true }

    if s.field[row][col] != 0 {
         if s.next(row, col) { return true }
    } else {
        for number := 1 ; number < 10 ; number++ {
            if s.add_number(row, col, number) {
                if s.next(row, col) { return true }
            }
        }
	s.field[row][col] = 0
    }
   return false
}

func (s *Sudoku) ToString() string {

    var str string

    str += "-------------\n"

    for row := 0 ; row<9 ; row++ {
        str += "|"
	for col := 0 ; col<9 ; col++ {
	    if s.field[row][col] == 0 {
                str += " "
            } else {
                str += strconv.Itoa(s.field[row][col])
            }

            if (col+1)%3 == 0 {
                str += "|"
            }

        }
	str += "\n"
        if (row+1)%3 == 0 {
            str += "-------------\n"
        }

    }
    
    return str
}

func main() {

    //parse flags
    consistency_only := flag.Bool("c", false, "Only check consistency of the entered sudoku.")

    flag.Parse()
    flag_args := flag.Args()    

    var filename string

    switch {
    case len(flag_args) == 1:
        filename = flag_args[0]
    case len(flag_args) == 0:
        fmt.Println("Please specify a path to the sudoku file.")
	os.Exit(1)
    default:
        fmt.Println("To many flag arguments. Filepath must be the last argument.")
	os.Exit(1)
    }

    //load sudoku
    sudoku := Sudoku{}
    err := sudoku.Fill(filename)
    if err != nil {
        fmt.Println("File not readable.")
	os.Exit(1)
    }

    if !sudoku.check_consistency() {
        fmt.Print("Sudoku is not consistent and thus not solvable.\n")
	fmt.Println(sudoku.ToString())
	os.Exit(1)
    }

    if *consistency_only {
        fmt.Println("Sudoku is consistent.")
        fmt.Println(sudoku.ToString())
        os.Exit(0)
    }

    fmt.Println("Sudoku to solve:")
    fmt.Println(sudoku.ToString())

    sudoku.solve(0,0)

    fmt.Println("Solved Sudoku:")
    fmt.Println(sudoku.ToString())

}
