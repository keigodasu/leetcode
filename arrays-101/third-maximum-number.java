class Solution {
    public int thirdMax(int[] nums) {
        Set<Integer> setNums = new HashSet<>();
        for(int num: nums) setNums.add(num);
        
        int max = Collections.max(setNums);
        
        if(setNums.size() < 3) {
            return max;
        }
        
        setNums.remove(max);
        int secondMax = Collections.max(setNums);
        setNums.remove(secondMax);
        return Collections.max(setNums);
    }
}

class Solution {
    public int thirdMax(int[] nums) {
        Set<Integer> maximums = new HashSet<Integer>();
        for(int num: nums) {
            maximums.add(num);
            if(maximums.size() > 3) {
                maximums.remove(Collections.min(maximums));
            }
        }
        if(maximums.size() == 3) {
            return Collections.min(maximums);
        }
        return Collections.max(maximums);
    }
}
