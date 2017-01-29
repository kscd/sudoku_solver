# sudoku_solver

This is a sudoku solver for standard 9x9 sudokus written in Go using the backtracking method.

To solve a sudoku simply type

./sudokusolver /path/to/sudoku

in which "/path/to/sudoku" should be replaced by the actual path to the sudoku. Before the program starts to solve the sudoku, the consistency of the given sudoku is checked. With the flag "-c" it is possible to just check the consistency and skip the actual solving. However, the -c flag needs then to be the first commandline argument.

The sudoku needs to be stored in a simple text file, in which each row is on its only line, while in each line the numbers are not seperated.
