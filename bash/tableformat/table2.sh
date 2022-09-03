#/bin/bash
seperator=--------------------
seperator=$seperator$seperator
rows="%-15s| %.7d| %3d| %c\n"
TableWidth=37

printf "%-15s| %-7s| %.3s| %s\n" Name ID Age Grades
printf "%.${TableWidth}s\n" "$seperator"
printf "$rows" "Sherlock Holmes" 122 23 A
printf "$rows" "James Bond" 7 27 F
printf "$rows" "Hercules Poirot" 6811 59 G
printf "$rows" "Jane Marple" 1234567 71 C
