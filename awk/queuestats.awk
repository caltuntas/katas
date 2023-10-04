BEGIN {
  FS=":"
  OFS="\t"
  RS="[[:space:]]?+[0-9]+) \"bull:"
  ORS="\r\n"
}

{
  if(length($NF)!=0) {
    gsub(/\r\n/,"",line);
    gsub(/"/, "");
    queues[$1];
    statuses[$2];
    vals[$1, $2] = $NF;
  }
}

END {
    printf "%-40s", "Queue"
    for (status in statuses) {
        printf "%-15s", status
    }
    print ""

    for (queue in queues) {
        printf "%-40s", queue
        for (status in statuses) {
            printf "%-15s", ((queue,status) in vals ? vals[queue,status] : "-")
        }
        print ""
    }
}
