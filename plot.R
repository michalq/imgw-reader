library(DBI)
library(RSQLite)
library(ggplot2)

con <- dbConnect(SQLite(), "identifier.sqlite")
sql <- "
  select y as year, avg(t) as avg_temp, max(tmax) max_temp, min(tmin) min_temp
  from measurements
  where m=1
  group by y, m
"

df <- dbGetQuery(con, sql)

warmest <- df |>
  dplyr::slice_max(avg_temp, n = 1)

p <- ggplot(df, aes(x = year)) +

  # zakres temperatur
  geom_ribbon(
    aes(ymin = min_temp, ymax = max_temp),
    fill = "#2C7FB8",
    alpha = 0.25
  ) +

  # Warmest years
  geom_vline(
    data = warmest,
    aes(xintercept = year),
    linetype = "dashed",
    color = "#2B2B2B",
    linewidth = 0.2
  ) +
  geom_text(
    data = warmest,
    aes(
      x = year,
      y = max(df$max_temp),
      label = paste0(year)
    ),
    angle = 90,
    vjust = -0.4,
    size = 3
  ) +

  # średnia
  geom_line(
    aes(y = avg_temp),
    color = "#08306B",
    linewidth = 0.5
  ) +

  # trend / regresja liniowa
  geom_smooth(
    aes(y = avg_temp),
    method = "lm",
    se = TRUE,
    color = "#D7301F",
    linewidth = 0.2
  ) +

  labs(
    title = "January Temperature Trend",
    subtitle = "Average with min–max range and linear trend",
    x = "Year",
    y = "Temperature (°C)",
    caption = "Source: IMGW"
  ) +
  theme_minimal(base_size = 13) +
  theme(
    plot.title = element_text(face = "bold", size = 18),
    plot.subtitle = element_text(color = "grey30"),
    panel.grid.minor = element_blank()
  )

dbDisconnect(con)