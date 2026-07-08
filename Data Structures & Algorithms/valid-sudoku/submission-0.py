class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for rows in board:
            row_check = {}
            for ele in rows:
                if ele != '.' and ele in row_check:
                    return False
                else:
                    row_check[ele] = 1

        for i in range(9):
            col_check = {}
            for rows in board:
                if rows[i] != '.' and rows[i] in col_check:
                    return False
                else:
                    col_check[rows[i]] = 1

        for i in range(0, 7, 3):
            for j in range(0, 7, 3):
                square_check = {}
                for k in range(i, i+3):
                    for l in range(j, j+3):
                        if board[k][l] != '.' and board[k][l] in square_check:
                            return False
                        else:
                            square_check[board[k][l]] = 1
        return True
                