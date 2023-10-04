# converts redis hgetall dump to json
BEGIN{
  printf "["
}
{
  if(NF<=0 || NR <=1){
    next;
  }
  line=$0;
  gsub(/\\""/,"\"",line);
  gsub(/: \\"/,"",line);
  gsub(/"{\\"/,"{\"",line);
  gsub(/}"/,"}",line);
  gsub(/\\"/,"\"",line);
  gsub(/"\[/,"[",line);
  gsub(/\]"/,"]",line);
  printf "%s%s",(NR==2?"":"," ORS),line
}
END{
  print "]"
}
