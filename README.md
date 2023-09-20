# brzHolidayCalendar
Go Module with help methods for computing working days considering Brazilian financial holidays

# Instalation
Just add the github.com/fabiobarbosa1984/brzHolidayCalendar reference on your go.mod and you are good to go. 

# Holidays Reference
For holidays referente, it was used public data published by [ANBIMA](https://www.anbima.com.br/feriados/feriados.asp).
As defined by the database publisher, only holidays with financial impact was considered. 

# Usage
Import the referente to your .go file and the following methods will be at your service:

## Main method

### NetWorkDays (initialDate time.Time, finalDate time.Time)
Returns the number of working days between (including) initialDate and (excluding) finalDate

## Aux methods
### NextWorkDay (date time.Time)
Returns the first valid work day since informed date (including informed date)

### StringToDate (date string)
Return the informed date as a time.Time object. Expects the string date in Brazilian format (DD/MM/YYYY)

### DateToString (date time.Time)
Convert a time.Time date to string using Brazilian format (DD/MM/YYYY)

# Contact
If you need any help to use this module in your project, feel free to contact me by my e-mail fabio.barbosa@gmail.com or using my linkedin https://www.linkedin.com/in/fabio-dos-santos-barbosa-affc/
