# The 1 Billion Row Challenge in Go

These are my solutions for the 1brc, in go. Each solution is progressively faster than the next.

| Version | Time   | Changes made                                                                                        |
| ------- | ------ | --------------------------------------------------------------------------------------------------- |
| naive   | 1m 38s | naive version: just a simple loop that goes through the file data and computes min, mean, max.      |
| v2      | 1m 13s | ditched strings.split for strings.Index(";"). in addition, made a small tweak to the map operations |
