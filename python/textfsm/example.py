import sys

import textfsm

# Run text through the FSM.
# The argument 'template' is a file handle and 'raw_text_data' is a string.
with open(sys.argv[1]) as template:
    re_table = textfsm.TextFSM(template)
with open(sys.argv[2]) as f:
    raw_text_data = f.read()
data = re_table.ParseText(raw_text_data)

# Display result as CSV
# First the column headers
print(", ".join(re_table.header))
# Each row of the table.
for row in data:
    print(", ".join(row))
