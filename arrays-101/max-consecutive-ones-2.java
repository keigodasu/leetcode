class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int count01 = 0;
        int count02 = 0;
        int max = 0;
        for (int num: nums) {
            if(num ==1) {
                count01++;
            } else {
                count02 = count01 + 1;
                count01 = 0;
            }
            max = Math.max(max, count01+count02);
        }
        
        return max;
    }
}
