import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from . import register

@register("month-year-by-year")
def generate_month_year_by_year_plot(con: sqlite3.Connection):
    sql = """
    SELECT y AS year, AVG(t) AS avg_temp, MAX(tmax) AS max_temp, MIN(tmin) AS min_temp
    FROM measurements
    WHERE m = 1
    GROUP BY y, m
    """
    df = pd.read_sql_query(sql, con)

    # Warmest year (max avg_temp)
    warmest = df.loc[df["avg_temp"].idxmax()]

    sns.set_theme(style="whitegrid", font_scale=1.1)
    fig, ax = plt.subplots(figsize=(10, 6))

    # Temperature range (ribbon)
    ax.fill_between(
        df["year"],
        df["min_temp"],
        df["max_temp"],
        color="#2C7FB8",
        alpha=0.25,
    )

    # Warmest year vertical line and label
    ax.axvline(
        x=warmest["year"],
        linestyle="--",
        color="#2B2B2B",
        linewidth=0.5,
    )
    ax.text(
        warmest["year"],
        df["max_temp"].max(),
        str(int(warmest["year"])),
        rotation=90,
        va="bottom",
        ha="center",
        fontsize=10,
    )

    # Average temperature line
    ax.plot(
        df["year"],
        df["avg_temp"],
        color="#08306B",
        linewidth=1,
    )

    # Linear trend with confidence interval (like geom_smooth(method="lm", se=TRUE))
    sns.regplot(
        data=df,
        x="year",
        y="avg_temp",
        scatter=False,
        ci=95,
        color="#D7301F",
        line_kws={"linewidth": 0.8},
        ax=ax,
    )

    ax.set_title("January Temperature Trend", fontsize=18, fontweight="bold")
    ax.set_xlabel("Year")
    ax.set_ylabel("Temperature (°C)")
    fig.suptitle(
        "Average with min–max range and linear trend",
        fontsize=13,
        color="grey",
        y=1.02,
    )
    ax.grid(True, which="major")
    ax.set_axisbelow(True)
    fig.text(0.99, 0.01, "Source: IMGW", ha="right", fontsize=9, color="gray")

    plt.tight_layout()
    plt.show()
