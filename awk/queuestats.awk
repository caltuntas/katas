BEGIN {
    FS=":"
}

{
if(length($NF)!=0) {
    queues[$2];
    statuses[$3];
    vals[$2, $3] = $NF;
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

