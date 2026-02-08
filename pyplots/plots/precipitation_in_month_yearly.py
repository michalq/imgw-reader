import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from . import register

@register("precipitation-in-month-yearly")
def setup_precipitation_in_month_yearly(parser):
    parser.add_argument("--month", type=int, required=True)

    def run(args, con: sqlite3.Connection):
        sql = """
        SELECT y AS year, AVG(precipitation_mm) AS avg_precipitation_mm
        FROM measurements
        WHERE m = ?
        GROUP BY y, m
        """
        df = pd.read_sql_query(sql, con, params=(args.month,))

        sns.set_theme(style="whitegrid", font_scale=1.1)
        fig, ax = plt.subplots(figsize=(10, 6))

        ax.plot(df["year"], df["avg_precipitation_mm"], color="#08306B", linewidth=1,)

        # Linear trend with confidence interval (like geom_smooth(method="lm", se=TRUE))
        sns.regplot(
            data=df,
            x="year",
            y="avg_precipitation_mm",
            scatter=False,
            ci=95,
            color="#D7301F",
            line_kws={"linewidth": 0.8},
            ax=ax,
        )

        ax.set_title(f"{args.month} Precipitation", fontsize=18, fontweight="bold")
        ax.set_xlabel("Year")
        ax.set_ylabel("Precipitation [mm]")
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
    parser.set_defaults(func=run)