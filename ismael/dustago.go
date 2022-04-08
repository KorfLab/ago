/*
write a complexity filter

dust <fasta file> [-w 20 -h 1.3 -lc]

INPUT
>id1
GCATGCTATCGTGCTATACGGCCCCCTATCGGCGCTAGCTATGC
>id2
ATCGTAGCTACGTTAGCGATTCGGATCGTAGCTAGGTAGCATCG

OUTPUT:
replace low entropy sequences to NNNNN
>id1
GCATGCTATCGTGCTATACGGNNNNNTATCGGCGCTAGCTATGC

or change to lowercase: ATATATATATATATA ---> atatatatatata

in python:
need window size, threshold, lc

if entropy less than threshold, output N
if entropy greater than threshold, output normal nuc from seq
need counts of each nucleotide


read in fast file
filter complexity
output new file





*/
