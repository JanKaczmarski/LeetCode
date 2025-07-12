from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # hash current field (i, j) into: column, row and box, check if number alredy present
        # if present return False, if False wasn't returned sudoku is valid

        # sudoku board is always 9x9
        BOARD_LEN = 9

        rows = [[0] * BOARD_LEN for _ in range(BOARD_LEN)]
        cols = [[0] * BOARD_LEN for _ in range(BOARD_LEN)]
        # Box could also be represented as ternary number
        boxs = [[[0] * BOARD_LEN for _ in range(BOARD_LEN // 3)] for _ in range(BOARD_LEN // 3)]

        for row in range(BOARD_LEN):
            for col in range(BOARD_LEN):
                if board[row][col] == ".":
                    continue

                value = int(board[row][col])
                box_row, box_col = row // 3, col // 3

                rows[row][value - 1] += 1
                if rows[row][value - 1] > 1:
                    return False

                cols[col][value - 1] += 1
                if cols[col][value - 1] > 1:
                    return False

                boxs[box_row][box_col][value - 1] += 1
                if boxs[box_row][box_col][value - 1] > 1:
                    return False

        return True
