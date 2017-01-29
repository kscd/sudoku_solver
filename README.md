# sudoku_solver

This is a sudoku solver for standard 9x9 sudokus written in Go using the backtracking method.

To solve a sudoku simply type

./sudokusolver /path/to/sudoku

in which "/path/to/sudoku" should be replaced by the actual path to the sudoku. Before the program starts to solve the sudoku, the consistency of the given sudoku is checked. With the flag "-c" it is possible to just check the consistency and skip the actual solving. However, the -c flag needs then to be the first command line argument.

The sudoku needs to be stored in a simple text file in which each row is on its own line while in each line the numbers are not separated. Missing numbers are denoted with spaces.

The text file sudoku.txt contains a simple not solved example sudoku while in sudoku_solved.txt the solved version of this sudoku is held.