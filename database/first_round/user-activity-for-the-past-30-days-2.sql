SELECT 
round(ifnull(count(DISTINCT session_id)/count(DISTINCT user_id), 0),2) as average_sessions_per_user
FROM Activity 
WHERE datediff('2019-07-27', activity_date) <= 29
