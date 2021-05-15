select
distinct a.id,
(select name from accounts where id = a.id) name
from
logins a, logins b
where
a.id = b.id and datediff(a.login_date, b.login_date) between 1 and 4
group by a.id, a.login_date
having count(distinct b.login_date) = 4;
