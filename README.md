# IMGW data parser


Download and insert synop data

Step download=true will download all zips.
If false then only unpack and put into out/out.csv.
We then need to import it to sqlite.

# Build
`make all`

# Running

```
./synop_cli download --raw-dir ./raw/synop
./synop_cli import --raw-dir ./raw/synop --out ./out/out.csv
```

# Generating plots

Go to `pyplots/README.md`.
