#/bin/bash
seperator=-----------------------------------------------
seperator=$seperator$seperator
rows="%-10s| %8s | %8s | %8s | %8s |\n"
width=57
green=$(tput setaf 2)
normal=$(tput sgr0)

baseUrl="https://www.petrolofisi.com.tr/akaryakit-fiyatlari"
cmd="htmlq -t -w '#fuelPricesTableDesktop tbody>tr:first-child' | sed '/^[[:space:]]*$/d;1d; s# TL/LT##g' | tr '\n' ' '"
declare ankara=($(curl -s $baseUrl/ankara-akaryakit-fiyatlari | eval ${cmd}))
declare istanbul=($(curl -s $baseUrl/istanbul-akaryakit-fiyatlari | eval ${cmd}))
declare izmir=($(curl -s $baseUrl/izmir-akaryakit-fiyatlari | eval ${cmd}))
date=$(date "+%d-%m-%Y")

printf "%-10s| %8s | %8s | %8s | %8s\n" "${green}YAKIT TIPI" "ANKARA" "ISTANBUL" "IZMIR" "TARIH${normal}"
printf "%.${width}s\n" "$seperator"
printf "$rows" "Benzin" "${ankara[0]}" "${istanbul[0]}" "${izmir[0]}" "$date"
printf "$rows" "Mazot" "${ankara[1]}" "${istanbul[1]}" "${izmir[1]}" "$date"
printf "$rows" "LPG" "${ankara[3]}" "${istanbul[3]}" "${izmir[3]}" "$date"

