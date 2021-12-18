# days
Returns days to/from date or date before/after n days

## Usage

    // n. days from today
    days -d number 

    // days from this date
    days -date YYYY-MM-DD

### Days from today (or specific date)

    // 10 days after today
    days -d 10

    // 10 days before today
    days -d -10

    // 8 days after Christmas
    DDATE=2021-12-24 days -d 8

### Days from date

    // days to Christmas
    days -date 2021-12-25

    // days since last Christmas
    days -date 2020-12-25
