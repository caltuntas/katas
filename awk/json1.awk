# converts redis hgetall dump to json
BEGIN{
  printf "["
  RS="-----\n";
  FS="\n";
}
{
  if(NF<=0 || NR<2){ 
    next;
  }
  line=""
  for(i=1; i<=NF ; i++){
    if(length($i)!=0 || (length($i)==0 && i!=NF)) {
      if (match($i,/[{\["]/)==0) {
        line=line"\""$i"\""
      } else {
        line=line""$i
      }
      if(i%2) {
        line=line":"
      } else {
        if(i!=NF-1) {
          line=line","
        }
      }
    }
  }
  printf "%s{%s}",(NR==2?"":"," ORS),line
}
END{
  print "]"
}
