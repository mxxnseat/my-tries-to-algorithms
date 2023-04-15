/*  https://leetcode.com/problems/permutation-in-string

    Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

    In other words, return true if one of s1's permutations is the substring of s2.

    

    Example 1:

    Input: s1 = "ab", s2 = "eidbaooo"
    Output: true
    Explanation: s2 contains one permutation of s1 ("ba").
    Example 2:

    Input: s1 = "ab", s2 = "eidboaoo"
    Output: false
    

    Constraints:

    1 <= s1.length, s2.length <= 10^4
    s1 and s2 consist of lowercase English letters.
*/

function checkInclusion(s1: string, s2: string): boolean {
  const map: Record<string, number> = {},
    window = s1.length;
  for (const c of s1) {
    if (c in map) {
      map[c]++;
    } else {
      map[c] = 1;
    }
  }
  let mapSize = Object.keys(map).length,
    fast = 0,
    slow = 0;
  while (fast < s2.length) {
    const cf = s2[fast];
    if (cf in map) {
      map[cf]--;
      if (map[cf] === 0) {
        mapSize--;
      }
    }
    if (fast - slow + 1 < window) {
      fast++;
    } else if (fast - slow + 1 === window) {
      if (mapSize === 0) {
        return true;
      }
      const cs = s2[slow];
      if (cs in map) {
        map[cs]++;
        if (map[cs] === 1) {
          mapSize++;
        }
      }
      fast++;
      slow++;
    }
  }
  return false;
}

console.log(checkInclusion("abc", "eidbaoo"));
