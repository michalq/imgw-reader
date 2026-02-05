library(DBI)
library(RSQLite)
library(ggplot2)

con <- dbConnect(SQLite(), "identifier.sqlite")
sql <- "
  select y as year, avg(t) as avg_temp, max(tmax) max_temp
  from measurements
  where m=1
  group by y, m
"

df <- dbGetQuery(con, sql)

ggplot(df, aes(x = year, y = avg_temp)) +
  geom_line() +
  geom_point() +
  labs(
    title = "Average temperature in January",
    x = "Year",
    y = "Monthly avg °C"
  ) +
  theme_minimal()

dbDisconnect(con)