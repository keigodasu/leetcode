select
ifnull((select round(count(distinct requester_id,accepter_id)/count(distinct sender_id, send_to_id), 2) from RequestAccepted, FriendRequest), 0)
as accept_rate
