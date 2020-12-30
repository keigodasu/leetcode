select
'[0-5>' bin, sum(duration/60 < 5) total
from Sessions
union
select
'[5-10>' bin, sum(duration/60 >= 5 and duration/60 < 10) total
from Sessions
union
select
'[10-15>' bin, sum(duration/60 >= 10 and duration/60 < 15) total
from Sessions
union
select
'15 or more', sum(duration/60 >= 15) total
from Sessions;
