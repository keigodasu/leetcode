class Solution {
    public int removeDuplicates(int[] nums) {
        int uniquCount = 0;
        for(int i = 0; i < nums.length; i++) {
            if(i == 0 || nums[i] != nums[i-1]) {
                uniquCount++;
            }
        }
        
        int[] result = new int[uniquCount];
        int resutlNum = 0;
        for(int i = 0; i < nums.length; i++) {
            if(i == 0 || nums[i] != nums[i-1]) {
                nums[resutlNum] = nums[i];
                resutlNum++;
            }
        }
       
        return resutlNum;
    }
}
