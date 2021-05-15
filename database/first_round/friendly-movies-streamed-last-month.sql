select
distinct c.title
from
TVProgram t
join
Content c
on t.content_id = c.content_id
where
c.Kids_content = 'Y' and
c.content_type = 'Movies' and
program_date between '2020-06-01' and '2020-06-30'
