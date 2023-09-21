# converts redis hgetall dump to json
BEGIN{
  printf "["
  RS="[[:space:]]?+[0-9]+)[[:space:]]+[0-9]+)";
  FS="[[:space:]]{3,}[0-9]+)";
}
{
  if(NF<=0){
    next;
  }
  line=""
  for(i=1; i<=NF; i++){
    line=line""$i
    if(i%2) {
      line=line":"
    } else if(i!=NF) {
      line=line","
    }
  }
  gsub(/"{\\"/,"{\"",line);
  gsub(/}"/,"}",line);
  gsub(/\\"/,"\"",line);
  gsub(/"\[/,"[",line);
  gsub(/\]"/,"]",line);
  gsub(/\r\n/,"",line);
  printf "%s{%s}",(NR==2?"":"," ORS),line
}
END{
  print "]"
}
