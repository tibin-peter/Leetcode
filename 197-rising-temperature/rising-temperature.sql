-- self join 
SELECT w1.id
FROM Weather w1
JOIN Weather w2 
  ON w1.recordDate = w2.recordDate + INTERVAL '1 day' --check the date arithmatic using intervel only one day
WHERE w1.temperature > w2.temperature;