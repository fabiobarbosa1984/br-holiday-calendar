# brzHolidayCalendar
Go Module with help methods for computing working days considering Brazilian financial holydays

# Instalation
Just add the github.com/fabiobarbosa1984/brzHolidayCalendar reference on your go.mod and you are good to go. 

# Holidays Reference
For holidays referente, it was used public data published by ANBIMA [link](https://www.anbima.com.br/feriados/feriados.asp)
As defined by the database publisher, only holidays with financial impact was considered. 

# Usage
Import the referente to your .go file and the 4 avaiable methods will be at your service:

## Main method

### NeWorkDays (initialDate time.Time, finalDate time.Time)
Returns the number of working days between (including) initialDate and (excluding) finalDate
