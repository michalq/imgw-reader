import argparse
import sqlite3

from plots import PLOTS
from plots.month_year_by_year import generate_month_year_by_year_plot
from plots.average_yearly import generate_average_yearly

parser = argparse.ArgumentParser()
subparsers = parser.add_subparsers(dest="command")

# Command list
subparsers.add_parser("list")
# Command generate
generate_parser = subparsers.add_parser("generate")
generate_parser.add_argument("--name", required=True)

args = parser.parse_args()

con = sqlite3.connect("../identifier.sqlite")

def print_available_plots():
    print("Available plots:")
    for name in PLOTS:
        print("\t- " + name)

if args.command == "list":
    print_available_plots()
elif args.command == "generate":
    if args.name not in PLOTS:
        print("\nInvalid name: " + args.name)
        print_available_plots()
        exit(1)
    PLOTS[args.name](con)
con.close()
