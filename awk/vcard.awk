function pascalcase(prefix,l) 
{
  len=length(prefix) + 1
  name=substr(l,len)
  line=""
  n=split(name,parts," ");
  for (i=1; i<=n; i++) {
    p=tolower(parts[i])
    where=match(p,/[[:alpha:]]/);
    if (where != 0) {
      u=substr(p,1,where-1) toupper(substr(p,where,1)) substr(p,where+1)
    }
    line=line u " "
  }
  where=match(line,/;[a-z]/)
    if (where != 0) {
      line=substr(line,1,where-1) toupper(substr(line,where,2)) substr(line,where+2)
   }
  return (prefix line)
}

/^FN:/{
  $0=pascalcase("FN:",$0)
} 
/^N:/{
  $0=pascalcase("N:",$0)
} 1
