class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_string = []
        for c in s:
            if 'a' <= c <= 'z' or 'A' <= c <= 'Z' or c.isnumeric():
                clean_string.append(c.lower())

        m = len(clean_string) // 2
        for i in range(m):
            front = i
            back = len(clean_string) - front - 1
            if clean_string[front] != clean_string[back]:
                return False
        return True