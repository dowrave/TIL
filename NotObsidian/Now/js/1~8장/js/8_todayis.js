today = new Date()
start_day = new Date('2023-02-01')

passedTime = today.getTime() - start_day.getTime() // 1970-1-1 이후 지난 시간을 ms로 표시함
passedDay = Math.round(passedTime / (1000 * 60 * 60 * 24))


