import argparse
import sqlite3

from plots import PLOTS
from plots.month_year_by_year import setup_month_year_by_year_plot
from plots.average_yearly import generate_average_yearly
from plots.snow_depth_in_month_yearly import setup_snow_depth_in_month_yearly
from plots.precipitation_in_month_yearly import setup_precipitation_in_month_yearly

parser = argparse.ArgumentParser()
subparsers = parser.add_subparsers(dest="command")

# Command list
subparsers.add_parser("list")
# Command generate
generate_parser = subparsers.add_parser("generate")
generate_parser.add_argument("--name", required=True)

for name, setup in PLOTS.items():
    setup(subparsers.add_parser(name))

args = parser.parse_args()

con = sqlite3.connect("../identifier.sqlite")

def print_available_plots():
    print("Available plots:")
    for name in PLOTS:
        print("\t- " + name)

if hasattr(args, "func"):
    args.func(args, con)
elif args.command == "list":
    print_available_plots()
elif args.command == "generate":
    if args.name not in PLOTS:
        print("\nInvalid name: " + args.name)
        print_available_plots()
        exit(1)
    PLOTS[args.name](con)
else:
    parser.print_help()
con.close()
