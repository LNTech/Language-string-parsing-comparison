# Language parsing comparisons
---
Simple tests involving generating test data, writing test data, and reading/parsing delimited test data

All scripts have the same goal:
* Generate 2 24 char long strings, delimited with a , and repeated 5m times
* Write to text file
* Read from text file, parse delimited text (i.e. split it into 2 strings), and store in array/list etc.
* Measure times

Currently only features 2 scripts (GO, Python) with raw GO being the fastest of the two, I attempted compiling the GO script and it actually ran about .3 of a second slower but that could just be a coincidence. You could probably get GO to work even faster if you know it better than I do

Tested on:
* 16gb CL16 3200MHZ DDR4 Ram x2 (32gb total)
* Samsung 870 1tb Sata SSD
* AMD Ryzen 5 3600
* NVIDIA RTX 2070
---

## Graphs for fun
<p>Average read/parse time of 10m strings</p>
<img src="https://i.imgur.com/MNPZzek.png" width="300">
<br>

<a href="https://docs.google.com/spreadsheets/d/1_zwmvHEZuWztUtsCrzLk5l4cBDo2NCoJqucmeOkOoKM/edit?usp=sharing">Spreadsheet for data fun</a>
