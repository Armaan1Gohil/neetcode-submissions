class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        track = {}

        for s in strs:
            check = [0] * 26
            for c in s:
                check[ord(c) - ord('a')] += 1
            if tuple(check) in track:
                track[tuple(check)].append(s)
            else:
                track[tuple(check)] = [s]
        return track.values()