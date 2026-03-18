-- Write your PostgreSQL query statement below
select player_id,min(event_date) as first_login from Activity group by player_id;

--min is the hero here it sort the date and find the minimum value for player
