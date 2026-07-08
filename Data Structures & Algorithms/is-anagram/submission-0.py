class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        s_map, t_map = {}, {}

        for i in range(len(s)):
            s_char, t_char = s[i], t[i]
            s_map[s_char] = 1 + s_map.get(s_char, 0)
            t_map[t_char] = 1 + t_map.get(t_char, 0)

        if s_map == t_map:
            return True
        return False 