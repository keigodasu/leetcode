select
*
from
users
where mail regexp '^[a-z|A-Z]+[a-z|A-Z|0-9|_|.|-]*@leetcode.com$'
