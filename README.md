# Introduction
This is a utility to dump the leveldb of harmony protocol.

# Usage
hmydb [option] command

options:
   -d path                 path to the leveldb
   -v                      verbose output

command:
   dump <num>              dump block <num>
   revert <num>            revert to block <num>
   list <number>           list number of blocks and block hash

# vim: expandtab:tabstop=3
