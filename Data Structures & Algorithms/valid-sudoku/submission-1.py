class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row_check = [set() for _ in range(9)]
        col_check = [set() for _ in range(9)]
        square_check = [set() for _ in range(9)]

        for i in range(9):
            for j in range(9):
                ele = board[i][j]

                if ele == '.':
                    continue

                if ele in row_check[i]:
                    return False
                row_check[i].add(ele)

                if ele in col_check[j]:
                    return False
                col_check[j].add(ele)

                square = (i // 3)*3 + j // 3
                if ele in square_check[square]:
                    return False
                square_check[square].add(ele)
        return True
                
                
