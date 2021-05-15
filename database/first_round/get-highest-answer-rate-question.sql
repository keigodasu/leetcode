select
question_id survey_log
from
survey_log
where action = 'answer'
group by question_id
order by count(answer_id is not null)/q_num desc limit 1
